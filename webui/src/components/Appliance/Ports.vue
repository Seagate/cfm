<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container style="padding: 0">
    <v-data-table :headers="headers" fixed-header :items="bladePorts">
      <template v-slot:[`item.memorychunk`]="{ item }">
        {{ linkedMemoryChunks[item.id] }}
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
import { computed } from "vue";
import { useBladePortStore } from "../Stores/BladePortStore";
import { useBladeMemoryStore } from "../Stores/BladeMemoryStore";

export default {
  data() {
    return {
      headers: [
        {
          title: "PortId",
          align: "start",
          key: "id",
        },
        { title: "LinkedHost", key: "linkedPortUri" },
        { title: "LinkedMemory", key: "memorychunk" },
        { title: "GCxlId", key: "gCxlId" },
        { title: "LinkStatus", key: "linkStatus" },
      ],
    };
  },

  setup() {
    const bladePortStore = useBladePortStore();
    const bladeMemoryStore = useBladeMemoryStore();

    // Computed property to sort ports by the numerical part of the PortId
    const sortedBladePorts = computed(() => {
      return bladePortStore.bladePorts
        .slice() // Create a copy to avoid mutating the original array
        .sort((a, b) => {
          // Extract the numerical part from the ResourceId
          const numA = parseInt(a.id.replace(/^\D+/g, ""));
          const numB = parseInt(b.id.replace(/^\D+/g, ""));
          return numA - numB;
        });
    });

    // Match the linked memory chunks and put id(s) to the LinkMemory column
    const linkedMemoryChunks = computed(() => {
      const bladePorts = bladePortStore.bladePorts;
      const bladeMemory = bladeMemoryStore.bladeMemory;

      const linkedMemoryChunks = {};

      bladePorts.forEach((port) => {
        const linkedPortId = port.id;

        if (linkedPortId) {
          const linkedMemory = bladeMemory.find(
            (memory) => memory.memoryAppliancePort === linkedPortId
          );

          if (linkedMemory) {
            linkedMemoryChunks[port.id] = linkedMemory.id;
          }
        }
      });

      return linkedMemoryChunks;
    });

    return {
      bladePorts: sortedBladePorts,
      linkedMemoryChunks,
    };
  },
};
</script>