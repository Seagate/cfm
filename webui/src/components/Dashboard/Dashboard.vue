<template>
  <v-container style="width: 100%; height: 100vh">
    <VueFlow :nodes="nodes" :edges="edges" @node-click="handleNodeClick">
    </VueFlow>
  </v-container>
</template>

<script>
import { onMounted } from "vue";
import { useFlowData } from "./initial-elements";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useBladeStore } from "../Stores/BladeStore";
import { useServiceStore } from "../Stores/ServiceStore";
import { VueFlow } from "@vue-flow/core";
import { useRouter } from "vue-router";

export default {
  components: { VueFlow },

  setup() {
    const applianceStore = useApplianceStore();
    const hostStore = useHostStore();
    const bladeStore = useBladeStore();
    const serviceStore = useServiceStore();

    const router = useRouter();
    const { nodes, edges } = useFlowData();

    // Define the node click handler to skip to the target node device's detail page
    const handleNodeClick = async (event) => {
      const node = event.node || event;

      if (node && node.data && node.data.url) {
        const url = node.data.url;

        // Before pushing the url, the selected appliance/blade/host should be updated, otherwise the url will change back to the selected device(s)
        if (node.type == "appliance") {
          applianceStore.selectAppliance(node.data.label);
        }
        if (node.type == "blade") {
          applianceStore.selectAppliance(node.data.associatedAppliance);
          bladeStore.selectBlade(node.data.label);
        }
        if (node.type == "host") {
          hostStore.selectedHostId = node.data.label;
        }

        router.push(url);
      } else {
        console.error("Device not found");
      }
    };

    // Fetch appliances/blades/hosts when component is mounted
    onMounted(async () => {
      await serviceStore.getServiceVersion();
      await applianceStore.fetchAppliances();
      await hostStore.fetchHosts();
    });

    return { nodes, edges, handleNodeClick };
  },
};
</script>
