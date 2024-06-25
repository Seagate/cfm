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
        v-model="selectedHostId"
        color="#6ebe4a"
        bg-color="rgba(110, 190, 74, 0.1)"
      >
        <v-tab
          v-for="host in hosts"
          :value="host.id"
          :key="host.id"
          :id="host.id"
          @click="
            selectHost(host.id, host.ipAddress, host.port, host.localMemoryMiB)
          "
        >
          <v-row justify="space-between" align="center">
            <v-col> {{ host.id }} </v-col>
            <v-col>
              <v-btn icon variant="text">
                <v-icon
                  size="x-small"
                  color="warning"
                  @click="deleteHostWindowButton"
                  id="deleteHostWindow"
                  >mdi-close</v-icon
                >
                <v-tooltip activator="parent" location="end"
                  >Click here to delete this cxl-host</v-tooltip
                >
              </v-btn>
            </v-col>
          </v-row>
        </v-tab>
        <v-tab>
          <v-btn variant="text" id="addHost" @click="addNewHostWindowButton">
            <v-icon start color="#6ebe4a">mdi-plus-thick</v-icon>
            <v-tooltip activator="parent" location="end"
              >Click here to add new cxl-host</v-tooltip
            >
          </v-btn>
        </v-tab>
      </v-tabs>

      <v-window v-model="selectedHostId">
        <v-window-item v-for="host in hosts" :value="host.id" :key="host.id">
          <!-- Content for the selected host -->
          <!-- ---------------------------------------------- -->
          <!---First Row -->
          <!-- ---------------------------------------------- -->
          <v-row class="flex-0" dense>
            <v-col cols="12" sm="12" md="12" lg="4">
              <!-- Basic Information -->
              <v-card
                class="card-shadow"
                height="350"
                color="rgba(110, 190, 74, 0.1)"
              >
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Basic Information</v-toolbar-title
                  >
                </v-toolbar>
                <v-card-text>
                  <h2 class="text-h6 text-green-lighten-2">
                    CXL-Host
                    <v-btn icon variant="text">
                      <v-icon
                        @click="resyncHostWindowButton"
                        id="resyncHostButton"
                        >mdi-sync-circle</v-icon
                      >
                      <v-tooltip activator="parent" location="end"
                        >Click here to resynchronize this CXL-Host
                        device</v-tooltip
                      >
                    </v-btn>
                  </h2>
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
                      {{
                        host.localMemoryMiB
                          ? (host.localMemoryMiB / 1024).toFixed(0)
                          : "N/A"
                      }}
                    </v-list-item-subtitle>
                  </v-list-item>
                </v-list>
              </v-card>
            </v-col>
            <v-col cols="12" sm="12" md="12" lg="8">
              <!-- Ports-->
              <v-card class="card-shadow h-full" height="350">
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Ports Information</v-toolbar-title
                  >
                </v-toolbar>
                <v-card-text
                  style="max-height: 350px; overflow: hidden; padding: 0"
                >
                  <HostPorts />
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
          <!-- ---------------------------------------------- -->
          <!---Second Row -->
          <!-- ---------------------------------------------- -->
          <v-row class="card-shadow flex-grow-0" dense>
            <v-col cols="12" sm="12" md="12" lg="6">
              <!-- Memory Devices -->
              <v-card class="card-shadow h-full" height="350">
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Memory Devices</v-toolbar-title
                  >
                </v-toolbar>
                <v-card-text
                  style="max-height: 420px; overflow: hidden; padding: 0"
                >
                  <HostMemoryDevices />
                </v-card-text>
              </v-card>
            </v-col>
            <v-col cols="12" sm="12" md="12" lg="6">
              <!-- Memory -->
              <v-card class="card-shadow h-full" height="350">
                <v-toolbar height="45">
                  <v-toolbar-title style="cursor: pointer"
                    >Memory</v-toolbar-title
                  >
                </v-toolbar>
                <v-card-text
                  style="max-height: 420px; overflow: hidden; padding: 0"
                >
                  <HostMemory />
                </v-card-text>
              </v-card>
            </v-col>
          </v-row>
        </v-window-item>
      </v-window>
    </div>
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
          <div class="text-h6 pa-12">{{ selectedHostId }}</div>
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
            @click="deleteHostConfirm(selectedHostId)"
            >Delete</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogDeleteHostWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ deleteHostProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <v-dialog v-model="dialogDeleteHostSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">Delete a cxl-host succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          Deleted Host Id:
          <br />{{ deletedHostId }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="deleteHostSuccess"
            @click="dialogDeleteHostSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogDeleteHostFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">Delete a cxl-host failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ deleteHostError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="deleteHostFailure"
            @click="dialogDeleteHostFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogAddHost" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="text-h5">Add New CXL-Host</span>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-container>
            <v-row class="justify-center">
              <v-col>
                <v-text-field
                  :rules="[rules.required]"
                  v-model="newHostCredentials.ipAddress"
                  label="Ip Address"
                  id="inputIpAddressHost"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  :rules="[rules.required]"
                  v-model.number="newHostCredentials.port"
                  label="Port"
                  id="inputPortHost"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="justify-center">
              <v-col>
                <v-text-field
                  :rules="[rules.required]"
                  v-model="newHostCredentials.username"
                  label="Username"
                  id="inputUserNameHost"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="justify-center">
              <v-col>
                <div style="position: relative; width: 100%">
                  <v-text-field
                    :rules="[rules.required]"
                    v-model="newHostCredentials.password"
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
              <v-col>
                <v-text-field
                  v-model="newHostCredentials.customId"
                  id="inputCustomedHostId"
                  label="Cxl-host Name (Optional)"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="cancelAddHost"
            @click="dialogAddHost = false"
          >
            Cancel
          </v-btn>
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="confirmAddHost"
            @click="addNewHostConfirm"
          >
            Add
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogAddHostWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ addHostProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

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
        <h2 class="text-h5 mb-6">Add a cxl-host succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New CXL-Host Id:
          <br />{{ newHostId }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="addNewHostSuccess"
            @click="dialogAddHostSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogAddHostFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">Add a cxl-host failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ addHostError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="addNewHostFailure"
            @click="dialogAddHostFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <!-- The dialog of warning before resynchronizing a cxl-host by the user -->
    <v-dialog v-model="dialogResyncHost" max-width="600px">
      <v-card>
        <v-alert
          color="warning"
          icon="$warning"
          title="Alert"
          variant="tonal"
          text="Resynchronizing a CXL-Host deletes the host and adds it back."
        ></v-alert>
        <v-card-text>
          <div class="text-h6 pa-12">{{ selectedHostId }}</div>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="cancelResyncHost"
            @click="dialogResyncHost = false"
            >Cancel</v-btn
          >
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="confirmResyncHost"
            @click="resyncHostConfirm(selectedHostId)"
            >Resync</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogResyncHostSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">Resync the CXL-Host succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          The contents of CXL-Host {{ resyncHostId }} are updated.
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="resyncHostSuccess"
            @click="dialogResyncHostSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogResyncHostFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">Resync the Host failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ resyncHostError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="resyncHostFailure"
            @click="dialogResyncHostFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogResyncHostWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ resyncHostProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>
  </v-container>
</template>

<script lang="ts">
import { watch, computed, onMounted, ref } from "vue";
import { useHostStore } from "../Stores/HostStore";
import { useHostPortStore } from "../Stores/HostPortStore";
import { useHostMemoryStore } from "../Stores/HostMemoryStore";
import { useHostMemoryDeviceStore } from "../Stores/HostMemoryDeviceStore";
import { useRouter } from "vue-router";
import HostPorts from "./HostPorts.vue";
import HostMemory from "./HostMemory.vue";
import HostMemoryDevices from "./HostMemoryDevice.vue";

export default {
  data() {
    return {
      loadProgressText: "Loading the page, please wait...",
      resyncHostProgressText: "Resynchronizing the CXL-Host, please wait...",
      // The rules for the input fields when adding a new cxl-host
      rules: {
        required: (value: any) => !!value || "Field is required",
      },

      newHostCredentials: {
        username: "admin",
        password: "admin12345",
        ipAddress: "127.0.0.1",
        port: 8082,
        insecure: true,
        protocol: "http",
        customId: "",
      },
      addHostProgressText: "Adding host, please wait...",
      showHostPassword: false,
      newHostId: "",
      dialogAddHost: false,
      dialogAddHostWait: false,
      dialogAddHostSuccess: false,
      dialogAddHostFailure: false,
      addHostError: null as unknown,

      deleteHostProgressText: "Deleting host, please wait...",
      dialogDeleteHost: false,
      dialogDeleteHostWait: false,
      deletedHostId: null as unknown as string | undefined,
      dialogDeleteHostSuccess: false,
      dialogDeleteHostFailure: false,
      deleteHostError: null as unknown,

      resyncHostId: null as unknown as string | undefined, // Be used on success popup
      dialogResyncHost: false,
      dialogResyncHostWait: false,
      dialogResyncHostSuccess: false,
      dialogResyncHostFailure: false,
      resyncHostError: null as unknown,
    };
  },

  // The child components
  components: {
    HostPorts,
    HostMemory,
    HostMemoryDevices,
  },

  methods: {
    addNewHostWindowButton() {
      this.dialogAddHost = true;
    },

    /* Triggle the API hostsPost in host store to add a new host */
    async addNewHostConfirm() {
      // Make the add host popup disappear
      this.dialogAddHost = false;

      this.dialogAddHostWait = true;
      const hostStore = useHostStore();

      const newHost = await hostStore.addNewHost(this.newHostCredentials);
      this.addHostError = hostStore.addHostError as string;

      // Display success  popup once adding new host succeeded
      if (!this.addHostError) {
        this.newHostId = newHost?.id + "";

        // Set the new added host as the selected host
        const hosts = computed(() => hostStore.hosts);
        if (hosts.value.length > 0) {
          hostStore.selectHost(
            newHost?.id + "",
            newHost?.ipAddress + "",
            Number(newHost?.port),
            newHost?.localMemoryMiB
          );
        }
        this.dialogAddHostWait = false;
        this.dialogAddHostSuccess = true;
      } else {
        this.dialogAddHostWait = false;
        this.dialogAddHostFailure = true;
      }
      // Reset the credentials
      this.initContent();
    },

    deleteHostWindowButton() {
      this.dialogDeleteHost = true;
    },

    /* Trigger the API to delete the host */
    async deleteHostConfirm(selectedHost: string) {
      // Make the delete host popup disappear
      this.dialogDeleteHost = false;

      this.dialogDeleteHostWait = true;

      const hostStore = useHostStore();

      const deletedHost = await hostStore.deleteHost(selectedHost);

      this.deleteHostError = hostStore.deleteHostError;

      // Update the hosts and set the default selected host
      if (!this.deleteHostError) {
        this.deletedHostId = deletedHost;
        const hosts = computed(() => hostStore.hosts);

        // Check if there are any hosts left after deletion
        if (hosts.value.length > 0) {
          // Set the first host as selected
          const selectedHost = hosts.value[0];
          hostStore.selectHost(
            selectedHost.id,
            selectedHost.ipAddress,
            selectedHost.port,
            selectedHost.localMemoryMiB
          );
        }

        this.dialogDeleteHostWait = false;
        this.dialogDeleteHostSuccess = true;
      } else {
        this.dialogDeleteHostWait = false;
        this.dialogDeleteHostFailure = true;
      }
    },

    resyncHostWindowButton() {
      this.dialogResyncHost = true;
    },

    async resyncHostConfirm(hostId: string) {
      this.dialogResyncHost = false;
      this.dialogResyncHostWait = true;

      const hostStore = useHostStore();
      await hostStore.resyncHost(hostId);

      this.resyncHostId = hostId;

      this.resyncHostError = hostStore.resyncHostError;

      // Display the cxl-host once resync cxl-host succeeded
      if (!this.resyncHostError) {
        // Manually trigger the update actions
        await this.updateHostContent(hostId);
        this.dialogResyncHostWait = false;
        this.dialogResyncHostSuccess = true;
      } else {
        this.dialogResyncHostWait = false;
        this.dialogResyncHostFailure = true;
      }
    },

    // Method to manually update the content for the resync cxl-host
    async updateHostContent(hostId: string) {
      const hostPortStore = useHostPortStore();
      const hostMemoryStore = useHostMemoryStore();
      const hostMemoryDeviceStore = useHostMemoryDeviceStore();

      await Promise.all([
        hostPortStore.hostPortStore(hostId),
        hostMemoryStore.hostMemoryStore(hostId),
        hostMemoryDeviceStore.hostMemoryDeviceStore(hostId),
      ]);
    },

    // Reset the filled newHost fields
    initContent() {
      this.newHostCredentials = {
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

  setup() {
    // Set up loading for progress linear
    const loading = ref(false);

    const hostStore = useHostStore();
    const hostPortStore = useHostPortStore();
    const hostMemoryStore = useHostMemoryStore();
    const hostMemoryDeviceStore = useHostMemoryDeviceStore();

    const router = useRouter();
    // Method to update the URL
    const updateUrl = (hostId: string) => {
      const newPath = `/hosts/${hostId}`;
      router.push(newPath);
    };

    // Fetch hosts when component is mounted
    onMounted(async () => {
      loading.value = true;
      await hostStore.fetchHosts();

      if (hostStore.hosts.length > 0) {
        let selectedHost:
          | {
              id: string;
              ipAddress: string;
              port: number;
              status?: string;
              ports?: { uri: string };
              "memory-devices"?: { uri: string };
              memory?: { uri: string };
              localMemoryMiB?: number;
              remoteMemoryMiB?: number;
            }
          | undefined;
        // Check if hostId exists in the URL
        const hostIdInUrl = router.currentRoute.value.params.host_id;
        if (hostIdInUrl) {
          // Find the host with the hostId from the URL
          selectedHost = hostStore.hosts.find(
            (host) => host.id === hostIdInUrl
          );
        }
        // If no hostId in the URL or no host found with the hostId, default to the first host
        if (!selectedHost) {
          selectedHost = hostStore.hosts[0];
        }
        hostStore.selectHost(
          selectedHost?.id + "",
          selectedHost?.ipAddress + "",
          Number(selectedHost?.port),
          selectedHost?.localMemoryMiB
        );
      }

      loading.value = false;
    });

    // Watch for changes in selected host and fetch the associated resources, ports, memory and memory devices
    watch(
      () => hostStore.selectedHostId,
      async (newHostId, oldHostId) => {
        if (newHostId !== null && newHostId !== oldHostId) {
          loading.value = true;
          // Fetch resources and ports for the newly selected host
          await Promise.all([
            hostPortStore.hostPortStore(newHostId),
            hostMemoryStore.hostMemoryStore(newHostId),
            hostMemoryDeviceStore.hostMemoryDeviceStore(newHostId),
          ]);

          // Update the URL with the new host ID
          updateUrl(newHostId);
          loading.value = false;
        }
      },
      { immediate: true }
    );

    // Computed properties to access state
    const hosts = computed(() => hostStore.hosts);

    const selectedHostId = computed(() => hostStore.selectedHostId);
    const selectedHostIp = computed(() => hostStore.selectedHostIp);
    const selectedHostPort = computed(() => hostStore.selectedHostPortNum);

    // Methods to update state
    const selectHost = (
      hostId: string,
      hostIp: string,
      hostPort: number,
      hostLocalMemory: number | undefined
    ) => {
      hostStore.selectHost(hostId, hostIp, hostPort, hostLocalMemory);
    };

    return {
      hosts,
      selectedHostId,
      selectedHostPort,
      selectedHostIp,
      selectHost,
      loading,
    };
  },
};
</script>

<style>
.v-tab {
  text-transform: none !important;
}
.highlighted-tab {
  font-weight: bold;
}
</style>