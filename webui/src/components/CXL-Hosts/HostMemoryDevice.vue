<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container style="padding: 0">
    <v-data-table
      :headers="headers"
      fixed-header
      height="240"
      :items="hostMemoryDevices"
    >
    </v-data-table>
  </v-container>
</template>

<script>
import { computed } from "vue";
import { useHostMemoryDeviceStore } from "../Stores/HostMemoryDeviceStore";

export default {
  data() {
    return {
      headers: [
        {
          title: "MemoryDeviceId",
          align: "start",
          key: "id",
        },
        { title: "DeviceType", key: "deviceType" },
        { title: "LinkStatus", key: "linkStatus" },
      ],
    };
  },

  setup() {
    const hostMemoryDeviceStore = useHostMemoryDeviceStore();

    // Computed property to sort memory devices by the numerical part of the MemoryDeviceId
    const sortedHostMemoryDevices = computed(() => {
      return hostMemoryDeviceStore.hostMemoryDevices
        .slice() // Create a copy to avoid mutating the original array
        .sort((a, b) => {
          // Extract the numerical part from the MemoryDeviceId
          const numA = parseInt(a.id.replace(/^\D+/g, ""));
          const numB = parseInt(b.id.replace(/^\D+/g, ""));
          return numA - numB;
        });
    });

    return {
      hostMemoryDevices: sortedHostMemoryDevices,
    };
  },
};
</script>