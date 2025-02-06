<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container
    style="
      width: 100%;
      height: 100vh;
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
          Discover new devices</v-btn
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
    <v-card
      class="parent-card"
      style="
        width: 100%;
        height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
      "
    >
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
      <v-card
        class="child-card"
        style="
          width: 20%;
          height: 50%;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
        "
      >
        <v-card-text>
          Devices
          <v-list-item>
            <v-list-item-title>
              <v-icon color="#f2ae72" class="mr-2">mdi-rectangle</v-icon>
              CMA
            </v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>
              <v-icon color="#f2e394" class="mr-2">mdi-rectangle</v-icon>
              Blade
            </v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>
              <v-icon color="#d9ecd0" class="mr-2">mdi-rectangle</v-icon>
              Host
            </v-list-item-title>
          </v-list-item>
          <br />Status
          <v-list-item>
            <v-list-item-title>
              <v-icon color="#6ebe4a" class="mr-2"
                >mdi-rectangle-outline</v-icon
              >
              online
            </v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>
              <v-icon color="#b00020" class="mr-2"
                >mdi-rectangle-outline</v-icon
              >
              offline
            </v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>
              <v-icon color="#ff9f40" class="mr-2"
                >mdi-rectangle-outline</v-icon
              >
              unavailable
            </v-list-item-title>
          </v-list-item>
          <v-list-item>
            <v-list-item-title>
              <v-icon color="#B0B0B0" class="mr-2"
                >mdi-rectangle-outline</v-icon
              >
              unknown
            </v-list-item-title>
          </v-list-item>
        </v-card-text>
      </v-card>
    </v-card>

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
          <br />
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
            id="cancelAddSelecteddDevices"
            @click="dialogNewDiscoveredDevices = false"
            >Cancel</v-btn
          >
          <v-btn
            color="info"
            variant="text"
            id="confirmAddSelecteddDevices"
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
      <v-card
        elevation="12"
        max-width="600"
        rounded="lg"
        width="100%"
        class="pa-4 text-center mx-auto"
      >
        <v-alert
          color="info"
          icon="$info"
          title="Results of adding devices"
          variant="tonal"
        ></v-alert>
        <br />
        <div
          v-if="
            (newBlades && newBlades.length) ||
            (failedBlades && failedBlades.length)
          "
        >
          New blades:
          <v-list>
            <v-list-item v-for="(blade, index) in newBlades" :key="index">
              <v-list-item-subtitle>
                <v-icon left style="color: green">mdi-check-circle</v-icon
                >{{ blade.id }}</v-list-item-subtitle
              >
            </v-list-item>
            <v-list-item v-for="(blade, index) in failedBlades" :key="index">
              <v-list-item-subtitle>
                <v-icon left style="color: red">mdi-close-circle</v-icon>
                {{ blade.name }}</v-list-item-subtitle
              >
            </v-list-item>
          </v-list>
        </div>
        <div
          v-if="
            (newHosts && newHosts.length) || (failedHosts && failedHosts.length)
          "
        >
          New hosts:
          <v-list>
            <v-list-item v-for="(host, index) in newHosts" :key="index">
              <v-list-item-subtitle
                ><v-icon left style="color: green">mdi-check-circle</v-icon
                >{{ host.id }}</v-list-item-subtitle
              >
            </v-list-item>
            <v-list-item v-for="(host, index) in failedHosts" :key="index">
              <v-list-item-subtitle>
                <v-icon left style="color: red">mdi-close-circle</v-icon>
                {{ host.name }}</v-list-item-subtitle
              >
            </v-list-item>
          </v-list>
        </div>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="info"
            rounded
            variant="flat"
            width="90"
            id="addNewDiscoveredDevicesOutput"
            @click="dialogAddNewDiscoveredDevicesOutput = false"
          >
            Done
          </v-btn>
        </div>
      </v-card>
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
      discoverDevicesProgressText: "Discovering devices, please wait...",

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
      failedBlades: [],
      failedHosts: [],
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
      // Initialize the new and failed devices
      this.newBlades = [];
      this.newHosts = [];
      this.failedBlades = [];
      this.failedHosts = [];

      this.dialogNewDiscoveredDevices = false;
      this.dialogAddNewDiscoveredDevicesWait = true;

      const applianceStore = useApplianceStore();
      const hostStore = useHostStore();
      const bladePortStore = useBladePortStore();

      if (this.selectedBlades.length === 0) {
        console.log("No blades selected.");
      } else {
        for (let i = 0; i < this.selectedBlades.length; i++) {
          try {
            const newAddedBlade = await applianceStore.addDiscoveredBlades(
              this.selectedBlades[i]
            );

            if (newAddedBlade) {
              this.newBlades.push(newAddedBlade);
            }

          } catch (error) {
            this.failedBlades.push(this.selectedBlades[i]);
            console.error("Error adding new discovered blade:", error);
          }
        }
      }

      if (this.selectedHosts.length === 0) {
        console.log("No hosts selected.");
      } else {
        for (let i = 0; i < this.selectedHosts.length; i++) {
          try {
            const newAddedHost = await hostStore.addDiscoveredHosts(
              this.selectedHosts[i]
            );

            if (newAddedHost) {
              this.newHosts.push(newAddedHost);
            }
            
          } catch (error) {
            this.failedHosts.push(this.selectedHosts[i]);
            console.error("Error adding new discovered host:", error);
          }
        }
      }

      // Update the graph content
      await applianceStore.fetchAppliances();
      await hostStore.fetchHosts();
      for (const appliance of applianceStore.applianceIds) {
        for (const blade of appliance.blades) {
          await bladePortStore.fetchBladePorts(appliance.id, blade.id);
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

.parent-card {
  position: relative;
}

.child-card {
  position: absolute;
  top: 0;
  right: 0;
  background-color: #f5f5f5;
  border: 1px solid #ccc;
}
</style>
