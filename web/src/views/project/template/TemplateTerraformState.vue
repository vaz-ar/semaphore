<template>
  <div>
    <v-alert
      type="info"
      text
      color="hsl(348deg, 86%, 61%)"
      style="border-radius: 0;"
      v-if="!premiumFeatures.terraform_backend"
    >
      Terraform/OpenTofu HTTP backend available only in <b>PRO</b> version.
      <v-btn
        class="ml-2"
        color="hsl(348deg, 86%, 61%)"
        href="https://semaphoreui.com/pro"
      >Upgrade
      </v-btn>
    </v-alert>

    <div class="px-4 py-3">
      <div v-for="alias of (aliases || [])" :key="alias.id">
        <code class="mr-2">{{ alias.url }}</code>
        <v-btn
          icon
          @click="copyToClipboard(alias.url, $t('aliasUrlCopied'))"
        >
          <v-icon>mdi-content-copy</v-icon>
        </v-btn>
        <v-btn icon @click="editAlias(alias.id)">
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
        <v-btn icon @click="deleteAlias(alias.id)">
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </div>

      <v-btn color="primary" @click="addAlias()" :disabled="(aliases || []).length === 0">
        {{ aliases == null ? $t('LoadAlias') : $t('AddAlias') }}
      </v-btn>
    </div>

    <v-data-table
      :headers="headers"
      :items="states"
      :footer-props="{ itemsPerPageOptions: [20] }"
      class="mt-0"
    >
      <template v-slot:item.id="{ item }">
        <router-link
          to="`/project/${item.project_id}/templates/${item.id}/terraform/state/${item.id}`"
        >
          {{ item.id }}
        </router-link>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-btn icon class="pl-1 pr-2" @click="deleteState(item)">
          <v-icon class="pr-1">mdi-trash</v-icon>
        </v-btn>
      </template>
    </v-data-table>
  </div>
</template>
<script>

import axios from 'axios';
import EventBus from '@/event-bus';

export default {
  props: {
    template: Object,
    repositories: Array,
    inventory: Array,
    environment: Array,
    premiumFeatures: Object,
    states: Array,
  },

  data() {
    return {
      headers: [{
        text: 'ID',
        value: 'id',
        sortable: false,
      }, {
        text: 'Task',
        value: 'task_id',
        sortable: false,
      }, {
        text: 'Actions',
        value: '',
        sortable: false,
      }],
      aliases: [],
    };
  },

  async created() {
    await this.loadAliases();
  },

  methods: {
    deleteState(id) {
      console.log(id);
    },

    async loadAliases() {
      this.aliases = (await axios({
        method: 'get',
        url: `/api/project/${this.template.project_id}/inventory/${this.template.id}/terraform/aliases`,
        responseType: 'json',
      })).data;
    },

    editAlias(alias) {
      console.log(alias);
    },

    async deleteAlias(alias) {
      await axios({
        method: 'delete',
        url: `/api/project/${this.template.project_id}/inventory/${this.template.inventory_id}/terraform/aliases/${alias}`,
      });
      await this.loadAliases();
    },

    async copyToClipboard(text) {
      try {
        await window.navigator.clipboard.writeText(text);
        EventBus.$emit('i-snackbar', {
          color: 'success',
          text: 'The command has been copied to the clipboard.',
        });
      } catch (e) {
        EventBus.$emit('i-snackbar', {
          color: 'error',
          text: `Can't copy the command: ${e.message}`,
        });
      }
    },

    async addAlias() {
      await axios({
        method: 'post',
        url: `/api/project/${this.template.project_id}/inventory/${this.template.inventory_id}/terraform/aliases`,
        responseType: 'json',
      });
      await this.loadAliases();
    },
  },
};
</script>
