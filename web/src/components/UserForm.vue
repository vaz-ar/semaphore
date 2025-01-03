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
      v-model="item.name"
      :label="$t('name')"
      :rules="[v => !!v || $t('name_required')]"
      required
      :disabled="formSaving"
    ></v-text-field>

    <v-text-field
      v-model="item.username"
      :label="$t('username')"
      :rules="[v => !!v || $t('user_name_required')]"
      required
      :disabled="formSaving"
    ></v-text-field>

    <v-text-field
      v-model="item.email"
      :label="$t('email')"
      :rules="[v => !!v || $t('email_required')]"
      required
      :disabled="item.external || formSaving"
    >

      <template v-slot:append>
        <v-chip outlined color="green" disabled small style="opacity: 1">private</v-chip>
      </template>
    </v-text-field>

    <v-text-field
      v-model="item.password"
      :label="$t('password')"
      type="password"
      :required="isNew"
      :rules="isNew ? [v => !!v || $t('password_required')] : []"
      :disabled="item.external || formSaving"
    ></v-text-field>

    <v-row>
      <v-col>
        <v-checkbox
          v-model="item.admin"
          :label="$t('adminUser')"
          v-if="isAdmin"
        ></v-checkbox>
      </v-col>
      <v-col>
        <v-checkbox
          v-model="item.alert"
          :label="$t('sendAlerts')"
        ></v-checkbox>
      </v-col>
    </v-row>

    <v-divider class="mb-4" />

    <h4>2FA</h4>

    <v-checkbox
      v-model="totpEnabled"
      label="Time-based one-time password"
    ></v-checkbox>

    <img
      v-if="totpQrUrl"
      :src="totpQrUrl"
      style="
        aspect-ratio: 1;
        border-radius: 6px;
        display: block;
        margin: 0 auto 10px auto;
      "
      alt="QR code"
    />

  </v-form>
</template>
<script>
import ItemFormBase from '@/components/ItemFormBase';
import axios from 'axios';

export default {
  props: {
    isAdmin: Boolean,
  },

  mixins: [ItemFormBase],

  data() {
    return {
      totpEnabled: false,
      totpQrUrl: null,
    };
  },

  watch: {
    async totpEnabled(val) {
      if (val) {
        if (this.item.totp == null) {
          this.item.totp = (await axios({
            method: 'post',
            url: `/api/users/${this.itemId}/2fas/totp`,
            responseType: 'json',
          })).data;
          this.totpQrUrl = `/api/users/${this.item.id}/2fas/totp/${this.item.totp.id}/qr`;
        }
      } else if (this.item.totp != null) {
        await axios({
          method: 'delete',
          url: `/api/users/${this.itemId}/2fas/totp/${this.item.totp.id}`,
          responseType: 'json',
        });
        this.item.totp = null;
        this.totpQrUrl = null;
      }
    },
  },

  methods: {
    afterLoadData() {
      this.totpEnabled = this.item.totp != null;
      this.totpQrUrl = `/api/users/${this.item.id}/2fas/totp/${this.item.totp.id}/qr`;
    },

    getItemsUrl() {
      return '/api/users';
    },

    getSingleItemUrl() {
      return `/api/users/${this.itemId}`;
    },
  },
};
</script>
