<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container style="width: 100%; height: 100vh">
    <VueFlow :nodes="nodes" :edges="edges"> </VueFlow>
  </v-container>
</template>

<script>
import { onMounted } from "vue";
import { useFlowData } from "./initial-elements";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useServiceStore } from "../Stores/ServiceStore";
import { VueFlow } from "@vue-flow/core";

export default {
  components: { VueFlow },

  setup() {
    const applianceStore = useApplianceStore();
    const hostStore = useHostStore();
    const serviceStore = useServiceStore();

    const { nodes, edges } = useFlowData();

    // Fetch appliances/blades/hosts when component is mounted
    onMounted(async () => {
      await serviceStore.getServiceVersion();
      await applianceStore.fetchAppliances();
      await hostStore.fetchHosts();
    });

    return { nodes, edges };
  },
};
</script>
