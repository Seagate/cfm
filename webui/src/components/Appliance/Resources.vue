<template>
  <v-container style="padding: 0">
    <v-data-table
      :headers="headers"
      fixed-header
      height="310"
      :items="selectedBladeResources"
    >
      <template v-slot:[`item.compositionStatus.compositionState`]="{ value }">
        <v-chip :color="getColor(value)">
          {{ value }}
        </v-chip>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import Helpers from "../Common/Helpers";
import { computed } from "vue";
import { useBladeResourceStore } from "../Stores/BladeResourceStore";

export default {
  mixins: [Helpers],

  data() {
    return {
      headers: [
        {
          title: "ResourceId",
          align: "start",
          key: "id",
        },
        {
          title: "Status",
          key: "compositionStatus.compositionState",
        },
        { title: "ChannelId", key: "channelId" },
        { title: "ChannelResourceIndex", key: "channelResourceIndex" },
        { title: "CapacityGiB", key: "capacityMiB" },
      ],
    };
  },

  setup() {
    const bladeResourceStore = useBladeResourceStore();

    // Computed property to sort resources by the numerical part of the ResourceId
    const sortedBladeResources = computed(() => {
      return bladeResourceStore.memoryResources
        .slice() // Create a copy to avoid mutating the original array
        .sort((a, b) => {
          // Extract the numerical part from the ResourceId
          const numA = parseInt(a.id.replace(/^\D+/g, ""));
          const numB = parseInt(b.id.replace(/^\D+/g, ""));
          return numA - numB;
        });
    });

    return {
      selectedBladeResources: sortedBladeResources,
    };
  },
};
</script>