<template>
  <v-alert
    type="info"
    text
    color="hsl(348deg, 86%, 61%)"
    style="border-radius: 0;"
  >
    Terraform/OpenTofu HTTP backend available only in <b>PRO</b> version.
    <v-btn
      class="ml-2"
      color="hsl(348deg, 86%, 61%)"
      href="https://semaphoreui.com/pro"
    >Upgrade</v-btn>
  </v-alert>
</template>
<script>

export default {
  props: {
    template: Object,
    repositories: Array,
    inventory: Array,
    environment: Array,
  },
  data() {
    return {
    };
  },
};
</script>
