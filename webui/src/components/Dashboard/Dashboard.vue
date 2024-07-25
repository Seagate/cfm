<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container style="width: 100%; height: 80vh">
    <h2 style="text-align: center; margin-bottom: 20px">
      CFM Ethernet Connections
    </h2>

    <div style="position: relative; width: 25%">
      <v-text-field
        v-if="showSearch"
        v-model="searchTerm"
        label="Search"
        @input="handleSearch"
        style="margin-bottom: 5px"
      ></v-text-field>
      <v-icon
        v-if="showSearch"
        style="position: absolute; right: 8px; top: 8px; cursor: pointer"
        @click="showSearch = false"
      >
        {{ "mdi-close" }}
      </v-icon>
    </div>

    <VueFlow
      :nodes="nodes"
      :edges="edges"
      class="basic-flow"
      :default-viewport="{ zoom: 1 }"
      :min-zoom="0.2"
      :max-zoom="4"
      @node-click="handleNodeClick"
    >
      <Controls position="top-left">
        <ControlButton title="Search" @click="toggleSearch">
          <v-icon>mdi-magnify</v-icon>
        </ControlButton>
      </Controls>
    </VueFlow>
  </v-container>
</template>

<script>
import { ref, onMounted } from "vue";
import { useFlowData } from "./initial-elements";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useBladeStore } from "../Stores/BladeStore";
import { useServiceStore } from "../Stores/ServiceStore";
import { VueFlow } from "@vue-flow/core";
import { useRouter } from "vue-router";
import { ControlButton, Controls } from "@vue-flow/controls";

export default {
  components: { VueFlow, ControlButton, Controls },

  setup() {
    const applianceStore = useApplianceStore();
    const hostStore = useHostStore();
    const bladeStore = useBladeStore();
    const serviceStore = useServiceStore();

    const router = useRouter();
    const { nodes, edges } = useFlowData();

    const searchTerm = ref("");
    const showSearch = ref(false);

    const handleSearch = () => {
      const term = searchTerm.value.toLowerCase();
      const elements = document.querySelectorAll(".basic-flow *");

      // Reset the color of all elements
      elements.forEach((el) => {
        el.style.color = "#000";
      });

      if (term) {
        let firstMatch = null;
        elements.forEach((el) => {
          // Change color of the matched element
          if (el.textContent.toLowerCase().includes(term)) {
            el.style.color = "red";
            if (!firstMatch) {
              firstMatch = el;
            }
          }
        });

        if (firstMatch) {
          firstMatch.scrollIntoView({ behavior: "smooth", block: "center" });
        }
      }
    };

    const toggleSearch = () => {
      showSearch.value = !showSearch.value;
    };

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

    return {
      nodes,
      edges,
      handleNodeClick,
      searchTerm,
      handleSearch,
      showSearch,
      toggleSearch,
    };
  },
};
</script>
