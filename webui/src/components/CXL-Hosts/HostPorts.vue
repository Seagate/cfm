<template>
  <v-container style="padding: 0">
    <v-data-table
      :headers="headers"
      fixed-header
      height="240"
      :items="hostPorts"
    >
    </v-data-table>
  </v-container>
</template>

<script>
import { computed } from "vue";
import { useHostPortStore } from "../Stores/HostPortStore";

export default {
  data() {
    return {
      headers: [
        {
          title: "PortId",
          align: "start",
          key: "id",
        },
        { title: "GCxlId", key: "gCxlId" },
        { title: "LinkedApplianceBlade", key: "linkedPortUri" },
        { title: "LinkStatus", key: "linkStatus" },
      ],
    };
  },

  setup() {
    const hostPortStore = useHostPortStore();

    // Computed property to sort ports by the numerical part of the PortId
    const sortedHostPorts = computed(() => {
      return hostPortStore.hostPorts
        .slice() // Create a copy to avoid mutating the original array
        .sort((a, b) => {
          // Extract the numerical part from the PortId
          const numA = parseInt(a.id.replace(/^\D+/g, ""));
          const numB = parseInt(b.id.replace(/^\D+/g, ""));
          return numA - numB;
        });
    });

    return {
      hostPorts: sortedHostPorts,
    };
  },
};
</script>