<template>
  <div v-if="state">
    <pre>
      <code>{{ state.state }}</code>
    </pre>
  </div>
  <div v-else-if="error">{{ error.message }}</div>
</template>

<script>
import axios from 'axios';

export default {
  props: {
    projectId: Number,
    inventoryId: Number,
  },

  data() {
    return {
      state: null,
      error: null,
    };
  },

  async created() {
    try {
      this.state = (await axios.get(`/api/project/${this.projectId}/inventory/${this.inventoryId}/terraform/states/latest`)).data;
    } catch (e) {
      if (e.response.status === 404) {
        this.error = {
          message: 'No state available.',
        };
      } else {
        this.error = e;
      }
    }
  },
};
</script>
