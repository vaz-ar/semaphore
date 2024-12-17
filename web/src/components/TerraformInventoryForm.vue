<template>
  <v-form
    ref="form"
    lazy-validation
    v-model="formValid"
    v-if="item != null"
  >
    <v-alert
      :value="formError"
      color="error"
      class="pb-2"
    >{{ formError }}</v-alert>

    <v-text-field
      v-model="item.inventory"
      label="Workspace name"
      :rules="[v => !!v || $t('path_required')]"
      required
      :disabled="formSaving"
    ></v-text-field>
  </v-form>
</template>
<style>
</style>
<script>
/* eslint-disable import/no-extraneous-dependencies,import/extensions */

import ItemFormBase from '@/components/ItemFormBase';
import axios from 'axios';

export default {
  mixins: [ItemFormBase],

  props: {
    namePrefix: String,
    type: String,
    templateId: Number,
  },

  components: {
  },

  data() {
    return {
      keys: null,
    };
  },

  methods: {
    async getNoneKey() {
      return (await axios({
        method: 'get',
        url: `/api/project/${this.projectId}/keys`,
        responseType: 'json',
      })).data.filter((key) => key.type === 'none')[0];
    },

    async beforeSave() {
      let noneKey = await this.getNoneKey();

      if (!noneKey) {
        await axios({
          method: 'post',
          url: `/api/project/${this.projectId}/keys`,
          responseType: 'json',
          data: {
            name: 'None',
            type: 'none',
            project_id: this.projectId,
          },
        });
        noneKey = await this.getNoneKey();
      }

      this.item.name = this.namePrefix + this.item.inventory;
      this.item.type = this.type;
      this.item.ssh_key_id = noneKey.id;
      this.item.template_id = this.templateId;
    },
    getItemsUrl() {
      return `/api/project/${this.projectId}/inventory`;
    },
    getSingleItemUrl() {
      return `/api/project/${this.projectId}/inventory/${this.itemId}`;
    },
  },
};
</script>
