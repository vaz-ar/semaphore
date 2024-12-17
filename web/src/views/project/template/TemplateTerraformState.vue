<template>
  <div v-if="inventories != null">
    <EditDialog
      v-model="inventoryDialog"
      :save-button-text="$t('create')"
      :icon="getAppIcon(template.app)"
      :icon-color="getAppColor(template.app)"
      :title="`${$t('nnew')} ${APP_INVENTORY_TITLE[template.app]}`"
      :max-width="450"
      @save="loadAliases"
    >
      <template v-slot:form="{ onSave, onError, needSave, needReset }">
        <TerraformInventoryForm
          :template-id="template.id"
          :project-id="template.project_id"
          :name-prefix="`${template.name} - `"
          :item-id="'new'"
          :type="`${template.app}-workspace`"
          @save="onSave"
          @error="onError"
          :need-save="needSave"
          :need-reset="needReset"
        />
      </template>
    </EditDialog>

    <div class="px-4 py-3">

      <div class="mb-8">
        <v-btn-toggle v-model="inventoryId" v-if="inventories.length > 0">
          <v-btn v-for="inv in inventories" :key="inv.id" :value="inv.id">
            {{ inv.inventory }}
            <v-icon
              v-if="inv.id === template.inventory_id"
              class="ml-1"
              color="success"
            >mdi-check
            </v-icon>
          </v-btn>
        </v-btn-toggle>

        <span v-else>No workspaces.</span>

        <v-btn large icon class="ml-2" @click="inventoryDialog = true">
          <v-icon>mdi-plus</v-icon>
        </v-btn>
      </div>

      <div class="mb-6" v-if="inventories.length > 0">
        <v-btn
          class="mr-4"
          :disabled="inventoryId === template.inventory_id"
          color="success"
        >Make default
        </v-btn>

        <!--        <v-btn color="primary" class="mr-4">-->
        <!--          Detach-->
        <!--        </v-btn>-->
        <v-btn
          color="error"
          :disabled="inventoryId === template.inventory_id"
        >
          Delete
        </v-btn>
      </div>

      <v-alert
        type="info"
        text
        color="hsl(348deg, 86%, 61%)"
        style="border-radius: 0; margin-left: -16px; margin-right: -16px;"
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

      <v-btn
        color="primary"
        @click="addAlias()"
        :disabled="!premiumFeatures.terraform_backend"
      >
        {{ aliases == null ? $t('LoadAlias') : $t('AddAlias') }}
      </v-btn>
    </div>

    <TerraformStateView
      v-if="premiumFeatures.terraform_backend"
    />

    <v-container v-else>
      <div style="text-align: center; color: grey;">No state available.</div>
    </v-container>

  </div>
</template>
<script>

import axios from 'axios';
import EventBus from '@/event-bus';
import TerraformStateView from '@/components/TerraformStateView.vue';
import TerraformInventoryForm from '@/components/TerraformInventoryForm.vue';
import EditDialog from '@/components/EditDialog.vue';
import { APP_INVENTORY_TITLE } from '@/lib/constants';
import AppsMixin from '@/components/AppsMixin';

export default {
  mixins: [AppsMixin],
  computed: {
    APP_INVENTORY_TITLE() {
      return APP_INVENTORY_TITLE;
    },
  },

  components: { EditDialog, TerraformInventoryForm, TerraformStateView },

  props: {
    template: Object,
    premiumFeatures: Object,
  },

  data() {
    return {
      aliases: [],
      state: null,
      inventories: null,
      inventoryDialog: null,
      inventoryId: null,
    };
  },

  async created() {
    this.inventoryId = this.template.inventory_id;

    this.inventories = (await axios({
      url: `/api/project/${this.template.project_id}/inventory?template_id=${this.template.id}`,
      responseType: 'json',
    })).data;

    if (
      (this.inventoryId == null
      || !this.inventories.some((inv) => inv.id === this.inventoryId))
      && this.inventories.length > 0) {
      this.inventoryId = this.inventories[0].id;
    }

    await this.loadAliases();
  },

  methods: {
    deleteState(id) {
      console.log(id);
    },

    async loadAliases() {
      try {
        this.aliases = (await axios({
          url: `/api/project/${this.template.project_id}/inventory/${this.template.id}/terraform/aliases`,
          responseType: 'json',
        })).data;
      } catch {
        this.aliases = null;
      }
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
