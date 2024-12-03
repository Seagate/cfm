<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container
    style="
      width: 100%;
      height: 80vh;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
    "
  >
    <h2 style="margin-bottom: 20px">
      {{ currentTitle }}
    </h2>

    <v-row>
      <v-col>
        <v-btn @click="toggleGraph" style="margin-bottom: 40px" variant="tonal">
          {{ buttonLabel }}
        </v-btn>
      </v-col>
      <v-col>
        <v-btn
          @click="discoverDevices"
          style="margin-bottom: 40px"
          variant="tonal"
        >
          Click to discover new devices</v-btn
        >
      </v-col>
    </v-row>

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

    <!-- The dialog of the warning before the adding the new discovered devices(blades or cxl-hosts) -->
    <v-dialog v-model="dialogNewDiscoveredDevices" max-width="600px">
      <v-card>
        <v-alert
          color="info"
          icon="$info"
          title="New Discovered Devices"
          variant="tonal"
          text="Here are the discovered devices. You can add them to CFM by selecting them and clicking the 'Add' button."
        ></v-alert>
        <v-card-text class="scrollable-content">
          New Discovered Blades:
          <v-checkbox
            v-for="(blade, index) in discoveredBlades"
            :key="index"
            :label="`${blade.name} - ${blade.address}`"
            v-model="selectedBlades"
            :value="blade"
          ></v-checkbox>
          <br>
          New Discovered Hosts:
          <v-checkbox
            v-for="(host, index) in discoveredHosts"
            :key="index"
            :label="`${host.name} - ${host.address}`"
            v-model="selectedHosts"
            :value="host"
          ></v-checkbox>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="info"
            variant="text"
            id="cancelResyncBlade"
            @click="dialogNewDiscoveredDevices = false"
            >Cancel</v-btn
          >
          <v-btn
            color="info"
            variant="text"
            id="confirmResyncBlade"
            @click="addDiscoveredDevices"
            >Add</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogAddNewDiscoveredDevicesWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ addNewDiscoveredDevicesProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <v-dialog v-model="dialogDiscoverDevicesWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ discoverDevicesProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <v-dialog v-model="dialogAddNewDiscoveredDevicesOutput" max-width="600px">
      <v-sheet
        elevation="12"
        max-width="600"
        rounded="lg"
        width="100%"
        class="pa-4 text-center mx-auto"
      >
        <v-icon
          class="mb-5"
          color="success"
          icon="mdi-check-circle"
          size="112"
        ></v-icon>
        <h2 class="text-h5 mb-6">Congrats! New devices were added</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New blades:
          <ul>
          <li v-for="(blade, index) in newBlades" :key="index">
            {{ blade.id }}
          </li>
        </ul><br />New hosts: <br />
        <ul>
          <li v-for="(host, index) in newHosts" :key="index">
            {{ host.id }}
          </li>
        </ul>
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="addNewApplianceSuccess"
            @click="dialogAddNewDiscoveredDevicesOutput = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>
  </v-container>
</template>

<script>
import { ref, onMounted, computed } from "vue";
import { useControlData } from "./initial-control-elements";
import { useData } from "./initial-data-elements";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useHostStore } from "../Stores/HostStore";
import { useBladeStore } from "../Stores/BladeStore";
import { useBladePortStore } from "../Stores/BladePortStore";
import { VueFlow } from "@vue-flow/core";
import { useRouter } from "vue-router";
import { ControlButton, Controls } from "@vue-flow/controls";

export default {
  components: { VueFlow, ControlButton, Controls },

  data() {
    return {
      addNewDiscoveredDevicesProgressText:
        "Adding the selected devices, please wait...",
      discoverDevicesProgressText:"Discovering the selected devices, please wait...",

      dialogNewDiscoveredDevices: false,
      dialogAddNewDiscoveredDevicesWait: false,
      dialogAddNewDiscoveredDevicesOutput: false,
      dialogDiscoverDevicesWait: false,

      discoveredBlades: [],
      discoveredHosts: [],

      selectedBlades: [],
      selectedHosts: [],

      newBlades: [],
      newHosts: [],
    };
  },

  methods: {
    async discoverDevices() {
      this.dialogDiscoverDevicesWait = true;
      const applianceStore = useApplianceStore();
      const hostStore = useHostStore();

      try {
        const responseOfBlades = await applianceStore.discoverBlades();
        const responseOfHosts = await hostStore.discoverHosts();

        const response = (responseOfBlades || []).concat(responseOfHosts || []);
        this.discoveredBlades = responseOfBlades || [];
        this.discoveredHosts = responseOfHosts || [];

        this.dialogNewDiscoveredDevices = true;
        this.dialogDiscoverDevicesWait = false;

        return response.length ? response : [];
      } catch (error) {
        this.dialogDiscoverDevicesWait = false;
        console.error("Error fetching data:", error);
        return [];
      }
    },

    async addDiscoveredDevices() {
      this.dialogNewDiscoveredDevices = false;
      this.dialogAddNewDiscoveredDevicesWait = true;

      const applianceStore = useApplianceStore();
      const hostStore = useHostStore();

      if (this.selectedBlades.length === 0) {
        console.log("No blades selected.");
      } else {
        for (var i = 0; i < this.selectedBlades.length; i++) {
          try {
            const newAddedBlade = await applianceStore.addDiscoveredBlades(
              this.selectedBlades[i]
            );

            if (newAddedBlade) {
              this.newBlades.push(newAddedBlade);
            }
          } catch (error) {
            console.error("Error adding new discovered blade:", error);
          }
        }
      }

      if (this.selectedHosts.length === 0) {
        console.log("No hosts selected.");
      } else {
        for (var i = 0; i < this.selectedHosts.length; i++) {
          try {
            const newAddedHost = await hostStore.addDiscoveredHosts(
              this.selectedHosts[i]
            );
            if (newAddedHost) {
              this.newHosts.push(newAddedHost);
            }
          } catch (error) {
            console.error("Error adding new discovered host:", error);
          }
        }
      }

      this.dialogAddNewDiscoveredDevicesWait = false;
      this.dialogAddNewDiscoveredDevicesOutput = true;
    },
  },

  setup() {
    const applianceStore = useApplianceStore();
    const hostStore = useHostStore();
    const bladeStore = useBladeStore();
    const bladePortStore = useBladePortStore();

    const router = useRouter();
    const { controlNodes, controlEdges } = useControlData();
    const { dataNodes, dataEdges } = useData();

    const searchTerm = ref("");
    const showSearch = ref(false);

    // Set Data Plane as the default one
    const currentGraph = ref("dataPlane");

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

    // Jump to the target page by changing the url
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

    // Toggle graphs by changing the current graph value
    const toggleGraph = () => {
      currentGraph.value =
        currentGraph.value === "controlPlane" ? "dataPlane" : "controlPlane";
    };

    // Change the nodes and edges by the current graph value
    const nodes = computed(() =>
      currentGraph.value === "controlPlane"
        ? controlNodes.value
        : dataNodes.value
    );
    const edges = computed(() =>
      currentGraph.value === "controlPlane"
        ? controlEdges.value
        : dataEdges.value
    );

    // Make the title of this conponent dynamic
    const currentTitle = computed(() =>
      currentGraph.value === "controlPlane"
        ? "CFM Ethernet Connections"
        : "CFM CXL Connections"
    );

    // Make the switch button label of this conponent dynamic
    const buttonLabel = computed(() =>
      currentGraph.value === "controlPlane"
        ? "Switch to Data Plane"
        : "Switch to Control Plane"
    );

    // Fetch appliances/blades/hosts when component is mounted
    onMounted(async () => {
      await applianceStore.fetchAppliances();
      await hostStore.fetchHosts();
      // Ensure blade ports are fetched after appliances, this action will create the edges for dataPlane
      for (const appliance of applianceStore.applianceIds) {
        for (const blade of appliance.blades) {
          await bladePortStore.fetchBladePorts(appliance.id, blade.id);
        }
      }
    });

    return {
      nodes,
      edges,
      currentGraph,
      handleNodeClick,
      searchTerm,
      handleSearch,
      showSearch,
      toggleSearch,
      toggleGraph,
      currentTitle,
      buttonLabel,
    };
  },
};
</script>

<style scoped>
.scrollable-content {
  max-height: 300px;
  overflow-y: auto;
}
</style>
