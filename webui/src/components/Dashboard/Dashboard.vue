<template>
  <v-container style="width: 100%; height: 100vh">
    <VueFlow
      ref="vueFlow"
      :nodes="nodes"
      :edges="edges"
      class="pinia-flow"
      fit-view-on-init
    >
    </VueFlow>
  </v-container>
</template>

<script>
import { onMounted, ref } from "vue";
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
    const vueFlow = ref(null);

    // Fetch appliances/blades/hosts when component is mounted
    onMounted(async () => {
      await serviceStore.getServiceVersion();
      await applianceStore.fetchAppliances();
      await hostStore.fetchHosts();
      vueFlow.value.fitView();
    });

    return { nodes, edges, vueFlow };
  },
};
</script>
