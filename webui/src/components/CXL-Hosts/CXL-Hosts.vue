<template>
  <v-container>
    <!-- Progress linear -->
    <v-dialog v-model="loading">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ loadProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <v-toolbar color="rgba(110, 190, 74, 0.1)">
      <v-toolbar-title>CXL-Host Devices</v-toolbar-title>
    </v-toolbar>
    <div>
      <v-tabs
        v-model="selectedHost"
        color="#6ebe4a"
        bg-color="rgba(110, 190, 74, 0.1)"
      >
        <v-tab
          v-for="host in hostDetailsList"
          :value="host.id"
          :key="host.id"
          style="text-transform: none"
          @click="
            NotDisplayAddNewHostCard();
            selectedHost = host.id;
          "
        >
          <v-row justify="space-between" align="center">
            <v-col>
              {{ host.id }}
            </v-col>
            <v-col>
              <v-btn icon variant="text" id="deleteHostIcon">
                <v-icon color="warning" size="x-small" @click="deleteHost(host)"
                  >mdi-close</v-icon
                >
                <v-tooltip activator="parent" location="end"
                  >Click here to delete this cxl-host</v-tooltip
                >
              </v-btn>
            </v-col>
          </v-row>
        </v-tab>
        <!-- Add a new cxl-host tab-->
        <v-tab>
          <v-btn
            variant="text"
            id="addHostIcon"
            @click="addNewHostWindowButton"
          >
            <v-icon color="#6ebe4a">mdi-plus-thick</v-icon>
            <v-tooltip activator="parent" location="end"
              >Click here to add new cxl-host</v-tooltip
            >
          </v-btn>
        </v-tab>
      </v-tabs>

      <v-window v-model="selectedHost" v-if="selectedHost !== null">
        <v-window-item
          v-for="host in hostDetailsList"
          :key="host.id"
          :value="host.id"
        >
          <v-row class="flex-0" dense>
            <v-col cols="12" sm="12" md="12" lg="4">
              <v-card
                height="350"
                class="card-shadow h-full"
                color="rgba(110, 190, 74, 0.1)"
              >
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Basic Information</v-toolbar-title
                  >
                </v-toolbar>
                <v-card-text>
                  💻A CXL-Host device is a Redfish Service agent providing local
                  memory composition.
                </v-card-text>
                <v-list lines="one">
                  <v-list-item>
                    <v-list-item-title>CXL-Host Id</v-list-item-title>
                    <v-list-item-subtitle>
                      {{ host.id }}
                    </v-list-item-subtitle>
                  </v-list-item>
                  <v-list-item>
                    <v-list-item-title>IpAddress</v-list-item-title>
                    <v-list-item-subtitle>
                      {{ host.ipAddress + ":" + host.port }}
                    </v-list-item-subtitle>
                  </v-list-item>
                  <v-list-item>
                    <v-list-item-title>LocalMemoryGiB</v-list-item-title>
                    <v-list-item-subtitle>
                      {{ (host.localMemoryMiB / 1024).toFixed(0) }}
                    </v-list-item-subtitle>
                  </v-list-item>
                </v-list>
              </v-card>
            </v-col>
            <v-col cols="12" sm="12" md="12" lg="8">
              <v-card height="350" class="card-shadow h-full">
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Ports Information</v-toolbar-title
                  >
                </v-toolbar>
                <Ports :hostId="host.id" />
              </v-card>
            </v-col>
          </v-row>
          <v-row class="card-shadow flex-grow-0" dense>
            <v-col cols="12" sm="12" md="12" lg="6">
              <v-card height="350" class="card-shadow h-full">
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Memory Devices</v-toolbar-title
                  >
                </v-toolbar>
                <MemoryDevices :hostId="host.id" />
              </v-card>
            </v-col>
            <v-col cols="12" sm="12" md="12" lg="6">
              <v-card v-card height="350" class="card-shadow h-full">
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Memory</v-toolbar-title
                  >
                </v-toolbar>
                <Memory :hostId="host.id" />
              </v-card>
            </v-col>
          </v-row>
        </v-window-item>
      </v-window>
      <v-window v-if="addNewHostWindow">
        <v-window-item>
          <v-card>
            <v-card-title class="text-center"> Add New CXL-Host </v-card-title>
            <v-card-text>
              <v-container>
                <v-row class="justify-center">
                  <v-col cols="4" sm="2" md="2">
                    <div style="position: relative; width: 100%">
                      <v-text-field
                        :rules="[rules.required]"
                        v-model="newHost.ipAddress"
                        label="Ip Address"
                        id="inputIpAddressHost"
                      ></v-text-field>
                    </div>
                  </v-col>
                  <v-col cols="4" sm="2" md="2">
                    <div style="position: relative; width: 100%">
                      <v-text-field
                        :rules="[rules.required]"
                        v-model.number="newHost.port"
                        label="Port"
                        id="inputPortHost"
                      ></v-text-field>
                    </div>
                  </v-col>
                </v-row>
                <v-row class="justify-center">
                  <v-col cols="8" sm="4" md="4">
                    <div style="position: relative; width: 100%">
                      <v-text-field
                        :rules="[rules.required]"
                        v-model="newHost.username"
                        label="Username"
                        id="inputUserNameHost"
                      ></v-text-field>
                    </div>
                  </v-col>
                </v-row>
                <v-row class="justify-center">
                  <v-col cols="8" sm="4" md="4">
                    <div style="position: relative; width: 100%">
                      <v-text-field
                        :rules="[rules.required]"
                        v-model="newHost.password"
                        label="Password"
                        id="inputPasswordHost"
                        :type="showHostPassword ? 'text' : 'password'"
                      ></v-text-field>
                      <v-icon
                        style="
                          position: absolute;
                          right: 8px;
                          top: 50%;
                          transform: translateY(-50%);
                        "
                        @click="showHostPassword = !showHostPassword"
                      >
                        {{ showHostPassword ? "mdi-eye" : "mdi-eye-off" }}
                      </v-icon>
                    </div>
                  </v-col>
                </v-row>
                <v-row class="justify-center">
                  <v-col cols="8" sm="4" md="4">
                    <div style="position: relative; width: 100%">
                      <v-text-field
                        v-model="newHost.customId"
                        id="inputCustomedHostId"
                        label="Cxl-host Name (Optional)"
                      ></v-text-field>
                    </div>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>
            <v-card-actions class="justify-center">
              <v-btn
                color="#6ebe4a"
                variant="text"
                id="resetAddHost"
                @click="initContent"
              >
                Reset
              </v-btn>
              <v-btn
                color="#6ebe4a"
                variant="text"
                id="confirmAddHost"
                @click="addNewHost"
              >
                Add
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-window-item>
      </v-window>
    </div>
    <!-- The dialog for adding a host successfully -->
    <v-dialog v-model="dialogAddHostSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">You added a cxl-host successfully!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New Host ID:
          <br />{{ this.addedHostDetails.data.id }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="addHostSuccess"
            @click="
              dialogAddHostSuccess = false;
              this.fetchHosts();
            "
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <!-- The dialog for adding an host failed -->
    <v-dialog v-model="dialogAddNewHostFailed" max-width="600px">
      <v-sheet
        elevation="12"
        max-width="600"
        rounded="lg"
        width="100%"
        class="pa-4 text-center mx-auto"
      >
        <v-icon
          class="mb-5"
          color="error"
          icon="mdi-alert-circle"
          size="112"
        ></v-icon>
        <h2 class="text-h5 mb-6">You added a cxl-host failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          Error Message:
          <br />
          {{ this.errorAddHostDetails }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="addHostFailure"
            @click="dialogAddNewHostFailed = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <!-- The dialog for deleting an cxl-host -->
    <v-dialog v-model="dialogDeleteHost" max-width="600px">
      <v-card>
        <v-alert
          color="warning"
          icon="$warning"
          title="Alert"
          variant="tonal"
          text="Delete this cxl-host? The action cannot be undone."
        ></v-alert>
        <v-card-text>
          <div class="text-h6 pa-12">{{ this.selectedHost }}</div>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="cancelDeleteHost"
            @click="dialogDeleteHost = false"
            >Cancel</v-btn
          >
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="confirmDeleteHost"
            @click="deleteHostConfirm(this.selectedHost)"
            >Delete</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { DefaultApi } from "@/axios/api";
import Ports from "./Ports.vue";
import MemoryDevices from "./MemoryDevices.vue";
import Memory from "./Memory.vue";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export default {
  components: {
    Ports,
    MemoryDevices,
    Memory,
  },

  data: () => ({
    loadProgressText: "Loading the page, please wait...",

    // The rules for the input fields when adding a new cxl-host
    rules: {
      required: (value) => !!value || "Field is required",
    },

    loading: true,

    dialogDeleteHost: false,
    dialogAddHostSuccess: false,
    selectedHost: null,
    dialogAddNewHostFailed: false,

    addNewHostWindow: false,

    hostDetailsList: [],

    newHost: {
      username: "admin",
      password: "admin12345",
      ipAddress: "127.0.0.1",
      port: 8082,
      insecure: true,
      protocol: "http",
      customId: "",
    },
    showHostPassword: false,
    addedHostDetails: "",
    errorAddHostDetails: "",
    hostCount: 0,
  }),

  created() {
    this.fetchHosts();
  },

  methods: {
    async fetchHosts() {
      this.hostDetailsList = [];
      try {
        // Get all hosts
        const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
        const response = await defaultApi.hostsGet();
        this.hostCount = response.data.memberCount;

        for (let i = 0; i < this.hostCount; i++) {
          // Extract the id for each host
          const uri = response.data.members[i];
          const hostId = JSON.stringify(uri).split("/").pop().slice(0, -2);

          // Get host by id
          const detailsResponse = await defaultApi.hostsGetById(hostId);

          // Store host in hosts list
          if (detailsResponse) {
            this.hostDetailsList.push(detailsResponse.data);
          }
        }
      } catch (error) {
        console.error("Error fetching hosts:", error);
      }
      this.loading = false;
    },

    deleteHost(item) {
      this.selectedHost = item;
      this.dialogDeleteHost = true;
    },

    async deleteHostConfirm(selectedHost) {
      const hostId = selectedHost;
      const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
      const response = await defaultApi.hostsDeleteById(hostId);
      // Update the hostDetailsList array, no need to refresh the whole page
      this.hostDetailsList = this.hostDetailsList.filter(
        (host) => host.id !== hostId
      );
      this.selectedHost = null;
      this.dialogDeleteHost = false;
    },

    async addNewHost() {
      try {
        const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
        const response = await defaultApi.hostsPost(this.newHost);
        this.addedHostDetails = response;
        this.errorAddHostDetails = "";
      } catch (error) {
        console.error("Error:", error);
        this.errorAddHostDetails =
          error.message || "An error occurred while deleting hosts";
      }
      // Meet error, open the failed dialog
      if (this.errorAddHostDetails) {
        this.dialogAddNewHostFailed = true;
      }
      // Otherwise, open the successful dialog
      else {
        this.dialogAddHostSuccess = true;
        this.addNewHostWindow = false;
      }
    },

    // Display addNewHostWindow when click the plus icon after the existed cxl-hosts
    addNewHostWindowButton() {
      this.addNewHostWindow = true;
    },

    // Not display addNewHostWindow when click the other tabs
    NotDisplayAddNewHostCard() {
      this.addNewHostWindow = false;
    },

    // Reset the filled newHost fields
    initContent() {
      this.newHost = {
        username: "admin",
        password: "admin12345",
        ipAddress: "127.0.0.1",
        port: 8082,
        insecure: true,
        protocol: "http",
        customId: "",
      };
    },
  },
};
</script>
