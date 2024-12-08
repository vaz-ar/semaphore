<template>
  <v-container fluid class="pb-0">
    <v-alert
      text
      type="info"
      class="mb-0 ml-4 mr-4 mb-2"
      v-if="template.description"
    >{{ template.description }}
    </v-alert>

    <v-row>
      <v-col>
        <v-list subheader>
          <v-list-item>
            <v-list-item-icon>
              <v-icon>mdi-book-play</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>{{ $t('playbook') }}</v-list-item-title>
              <v-list-item-subtitle>{{ template.playbook }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
      <v-col>
        <v-list subheader>
          <v-list-item>
            <v-list-item-icon>
              <v-icon>{{ TEMPLATE_TYPE_ICONS[template.type] }}</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>{{ $t('type') }}</v-list-item-title>
              <v-list-item-subtitle>{{ $t(TEMPLATE_TYPE_TITLES[template.type]) }}
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
      <v-col>
        <v-list subheader>
          <v-list-item>
            <v-list-item-icon>
              <v-icon>mdi-monitor</v-icon>
            </v-list-item-icon>

            <v-list-item-content>
              <v-list-item-title>{{ $t('inventory') }}</v-list-item-title>
              <v-list-item-subtitle>
                {{ (inventory.find((x) => x.id === template.inventory_id) || {name: '—'}).name }}
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
      <v-col>
        <v-list subheader>
          <v-list-item>
            <v-list-item-icon>
              <v-icon>mdi-code-braces</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ $t('environment') }}</v-list-item-title>
              <v-list-item-subtitle>
                {{ environment.find((x) => x.id === template.environment_id).name }}
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
      <v-col>
        <v-list subheader>
          <v-list-item>
            <v-list-item-icon>
              <v-icon>mdi-git</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>{{ $t('repository2') }}</v-list-item-title>
              <v-list-item-subtitle>
                {{ repositories.find((x) => x.id === template.repository_id).name }}
              </v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
    </v-row>

    <LineChart :source-data="stats"/>

  </v-container>

</template>
<script>
import {
  TEMPLATE_TYPE_ACTION_TITLES,
  TEMPLATE_TYPE_ICONS,
  TEMPLATE_TYPE_TITLES,
} from '@/lib/constants';
import axios from 'axios';
import LineChart from '@/components/LineChart.vue';

export default {
  components: { LineChart },
  props: {
    template: Object,
    repositories: Array,
    inventory: Array,
    environment: Array,
  },
  data() {
    return {
      TEMPLATE_TYPE_ICONS,
      TEMPLATE_TYPE_TITLES,
      TEMPLATE_TYPE_ACTION_TITLES,
      stats: null,
    };
  },
  async created() {
    this.stats = (await axios({
      method: 'get',
      url: `/api/project/${this.template.project_id}/templates/${this.template.id}/stats`,
      responseType: 'json',
    })).data;
  },
};
</script>
