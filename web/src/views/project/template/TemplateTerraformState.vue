<template>
  <div v-if="inventories != null">

    <EditDialog
      v-model="editDialog"
      :save-button-text="aliasId === 'new' ? $t('create') : $t('save')"
      :title="`${aliasId === 'new' ? $t('nnew') : $t('edit')} Key`"
      :max-width="450"
      @save="loadAliases()"
    >
      <template v-slot:form="{ onSave, onError, needSave, needReset }">
        <TerraformAliasForm
          :project-id="template.project_id"
          :item-id="aliasId"
          :inventory-id="template.inventory_id"
          @save="onSave"
          @error="onError"
          :need-save="needSave"
          :need-reset="needReset"
        />
      </template>
    </EditDialog>

    <YesNoDialog
      :title="$t('deleteInventory')"
      :text="$t('askDeleteInv')"
      v-model="deleteInventoryDialog"
      @yes="deleteInventory()"
    />

    <EditDialog
      v-model="inventoryDialog"
      :save-button-text="$t('create')"
      :icon="getAppIcon(template.app)"
      :icon-color="getAppColor(template.app)"
      :title="`${$t('nnew')} ${APP_INVENTORY_TITLE[template.app]}`"
      :max-width="450"
      @save="onNewInventory"
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

    <EditDialog
      v-model="attachInventoryDialog"
      :save-button-text="$t('Attach')"
      :icon="getAppIcon(template.app)"
      :icon-color="getAppColor(template.app)"
      :max-width="450"
      title="Choose workspace to attach"
      @save="attachInventory($event.itemId)"
    >
      <template v-slot:form="{ onSave, needSave, needReset }">
        <InventorySelectForm
          :app="template.app"
          :project-id="template.project_id"
          @save="onSave"
          :need-save="needSave"
          :need-reset="needReset"
        />
      </template>
    </EditDialog>

    <div class="px-4 py-3">

      <div class="mb-6">
        <v-btn-toggle
          v-model="inventoryId"
          v-if="inventories.length > 0"
          mandatory
          style="overflow: auto"
          class="mb-2"
        >
          <v-btn v-for="inv in inventories" :key="inv.id" :value="inv.id">
            {{ inv.inventory }}
            <v-icon
              dark
              v-if="inv.id === template.inventory_id"
              class="ml-1"
              color="success"
            >mdi-check
            </v-icon>
          </v-btn>
        </v-btn-toggle>

        <span v-else>No workspaces.</span>

        <v-menu offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              color="primary"
              dark
              fab
              small
              class="ml-2"
              v-bind="attrs"
              v-on="on"
            >
              <v-icon dark>mdi-plus</v-icon>
            </v-btn>
          </template>
          <v-list>
            <v-list-item @click="inventoryDialog = true">
              <v-list-item-icon>
                <v-icon>mdi-pencil</v-icon>
              </v-list-item-icon>
              <v-list-item-title>New workspace</v-list-item-title>
            </v-list-item>
            <v-list-item @click="attachInventoryDialog = true">
              <v-list-item-icon>
                <v-icon>mdi-connection</v-icon>
              </v-list-item-icon>
              <v-list-item-title >Attach existing workspace</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>

      <div class="mb-6" v-if="inventories.length > 0">
        <v-btn
          class="mr-4 mb-2"
          :disabled="inventoryId === template.inventory_id"
          color="success"
          @click="setDefaultInventory()"
        >
          Make default
        </v-btn>

        <v-btn
          color="primary"
          class="mr-4 mb-2"
          :disabled="inventoryId === template.inventory_id"
          @click="detachInventory()"
        >
          Detach
        </v-btn>

        <v-btn
          class="mb-2"
          color="error"
          :disabled="inventoryId === template.inventory_id"
          @click="deleteInventoryDialog = true;"
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
        <span class="mr-2">
          Terraform/OpenTofu HTTP backend available only in <b>PRO</b> version.
        </span>
        <v-btn
          color="hsl(348deg, 86%, 61%)"
          href="https://semaphoreui.com/pro"
        >Upgrade
        </v-btn>
      </v-alert>

      <div class="mb-3 pl-1" v-if="(aliases || []).length === 0">There is no aliases.</div>

      <div v-else v-for="alias of (aliases || [])" :key="alias.id">
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
      class="pl-4 pt-4"
      v-if="premiumFeatures.terraform_backend"
      :project-id="template.project_id"
      :inventory-id="template.inventory_id"
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
import YesNoDialog from '@/components/YesNoDialog.vue';
import InventorySelectForm from '@/components/InventorySelectForm.vue';
import TerraformAliasForm from '@/components/TerraformAliasForm.vue';

export default {
  mixins: [AppsMixin],
  computed: {
    APP_INVENTORY_TITLE() {
      return APP_INVENTORY_TITLE;
    },
  },

  components: {
    TerraformAliasForm,
    InventorySelectForm,
    YesNoDialog,
    EditDialog,
    TerraformInventoryForm,
    TerraformStateView,
  },

  props: {
    template: Object,
    premiumFeatures: Object,
  },

  watch: {
    async inventoryId() {
      await this.loadAliases();
    },
  },

  data() {
    return {
      aliases: [],
      state: null,
      inventories: null,
      inventoryDialog: null,
      inventoryId: null,
      deleteInventoryDialog: null,
      attachInventoryDialog: null,
      aliasId: null,
      editDialog: null,
    };
  },

  async created() {
    this.inventoryId = this.template.inventory_id;

    await this.loadInventories();
    await this.loadAliases();
  },

  methods: {
    deleteState(id) {
      console.log(id);
    },

    async setDefaultInventory() {
      await axios({
        method: 'post',
        url: `/api/project/${this.template.project_id}/templates/${this.template.id}/inventory/${this.inventoryId}/set_default`,
      });
      this.$emit('update-template', {});
    },

    async detachInventory() {
      await axios({
        method: 'post',
        url: `/api/project/${this.template.project_id}/templates/${this.template.id}/inventory/${this.inventoryId}/detach`,
      });
      await this.loadInventories();
    },

    async deleteInventory() {
      await axios({
        method: 'delete',
        url: `/api/project/${this.template.project_id}/inventory/${this.inventoryId}`,
      });
      await this.loadInventories();
    },

    async onNewInventory(e) {
      await this.loadInventories();
      this.inventoryId = e.item.id;
    },

    async attachInventory(inventoryId) {
      await axios({
        method: 'post',
        url: `/api/project/${this.template.project_id}/templates/${this.template.id}/inventory/${inventoryId}/attach`,
      });
      await this.loadInventories();
      this.inventoryId = inventoryId;
    },

    async loadInventories() {
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
    },

    async loadAliases() {
      try {
        this.aliases = (await axios({
          url: `/api/project/${this.template.project_id}/inventory/${this.inventoryId}/terraform/aliases`,
          responseType: 'json',
        })).data;
      } catch {
        this.aliases = null;
      }
    },

    async deleteAlias(alias) {
      await axios({
        method: 'delete',
        url: `/api/project/${this.template.project_id}/inventory/${this.inventoryId}/terraform/aliases/${alias}`,
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

    editAlias(alias) {
      this.aliasId = alias;
      this.editDialog = true;
    },

    async addAlias() {
      this.aliasId = 'new';
      this.editDialog = true;
    },
  },
};
</script>
