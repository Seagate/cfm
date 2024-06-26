<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container style="padding: 0">
    <v-data-table
      :headers="headers"
      fixed-header
      height="240"
      :items="hostMemory"
    >
    </v-data-table>
  </v-container>
</template>

<script>
import { computed } from "vue";
import { useHostMemoryStore } from "../Stores/HostMemoryStore";

export default {
  data() {
    return {
      headers: [
        {
          title: "MemoryId",
          align: "start",
          key: "id",
        },
        { title: "Type", key: "type" },
        { title: "SizeGiB", key: "sizeMiB" },
      ],
    };
  },

  setup() {
    const hostMemoryStore = useHostMemoryStore();

    // Computed property to sort memory by the numerical part of the MemoryId
    const sortedHostMemory = computed(() => {
      return hostMemoryStore.hostMemory
        .slice() // Create a copy to avoid mutating the original array
        .sort((a, b) => {
          // Extract the numerical part from the MemoryId
          const numA = parseInt(a.id.replace(/^\D+/g, ""));
          const numB = parseInt(b.id.replace(/^\D+/g, ""));
          return numA - numB;
        });
    });

    return {
      hostMemory: sortedHostMemory,
    };
  },
};
</script>