<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container>
    <v-dialog v-model="dialogComposeMemory" max-width="600px">
      <template v-slot:activator="{ props }">
        <v-btn
          v-bind="props"
          class="text-none text-subtitle-1"
          color="#6ebe4a"
          size="x-small"
          variant="plain"
          prepend-icon="mdi-plus-circle"
          id="composeMemoryButton"
        >
          Compose Memory
        </v-btn>
      </template>

      <v-card>
        <v-card-title>
          <span class="text-h5">Compose Memory</span>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12" sm="6" md="6">
                <v-text-field
                  :rules="[rules.required]"
                  v-model="associatedApplianceId"
                  label="Associated Appliance Id"
                  id="inputApplianceId"
                  readonly
                  style="color: #9e9e9e; pointer-events: none"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="6" md="6">
                <v-text-field
                  :rules="[rules.required]"
                  v-model="bladeId"
                  label="Blade Id"
                  id="inputBladeId"
                  readonly
                  style="color: #9e9e9e; pointer-events: none"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row>
              <v-col cols="12" sm="6" md="4">
                <v-text-field
                  :rules="[rules.required]"
                  v-model.number="memorySizeGiB"
                  label="Memory Size (GiB)"
                  id="inputMemorySize"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="6" md="4">
                <v-autocomplete
                  :rules="[rules.required]"
                  v-model.number="showQoS"
                  label="QoS"
                  id="inputQos"
                  :items="qoSs"
                ></v-autocomplete>
              </v-col>
              <v-col cols="12" sm="6" md="4">
                <v-autocomplete
                  v-model="showPort"
                  label="Port (Optional)"
                  id="inputPort"
                  :items="PortIdArray"
                ></v-autocomplete>
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
            id="cancelComposeMemory"
            @click="dialogComposeMemory = false"
          >
            Cancel
          </v-btn>
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="warningComposeMemory"
            @click="warningBeforeComposeMemory"
          >
            Compose Memory
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogComposeMemorySuccess" max-width="600px">
      <v-sheet
        elevation="12"
        max-width="600"
        rounded="lg"
        width="100%"
        class="pa-4 text-center mx-auto"
      >
        <v-icon
          class="mb-5"
          :color="partialSuccess ? 'warning' : 'success'"
          icon="mdi-check-circle"
          size="112"
        ></v-icon>
        <h2 class="text-h5 mb-6">
          Compose memory {{ partialSuccess ? "partially" : "" }} succeeded!
        </h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New Memory ID:
          <br />{{ newMemoryId }} <br />{{ memorySizeNotEqual }} <br />{{
            partialSuccess
          }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            :color="partialSuccess ? 'warning' : 'success'"
            rounded
            variant="flat"
            width="90"
            id="composeMemorySuccess"
            @click="dialogComposeMemorySuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogComposeMemoryFailed" max-width="600px">
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
        <h2 class="text-h5 mb-6">Compose memory failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ composeMemoryError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="composeMemoryFailure"
            @click="dialogComposeMemoryFailed = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogComposeMemoryAlert" max-width="600px">
      <v-card>
        <v-alert
          color="warning"
          icon="$warning"
          title="Alert"
          variant="tonal"
          text="Due to limited protections, the CXL-Host MUST be powered down when being assigned memory from a memory appliance."
        ></v-alert>
        <v-card-text>
          <div>
            Please <strong>power down</strong> the connected CXL-Host device
            before clicking the <strong>Continue</strong> button to compose
            memory.
          </div>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="cancelComposeMemoryAlert"
            @click="dialogComposeMemoryAlert = false"
            >Cancel</v-btn
          >
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="confirmComposeMemory"
            @click="composeMemoryConfirm"
          >
            Continue
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogComposeMemoryWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ composeMemoryProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>
  </v-container>
</template>

<script lang="ts">
import { useBladePortStore } from "../Stores/BladePortStore";
import { useBladeMemoryStore } from "../Stores/BladeMemoryStore";
import { useBladeResourceStore } from "../Stores/BladeResourceStore";
import { useBladeStore } from "../Stores/BladeStore";
import { computed } from "vue";
import { Qos } from "../../axios/api";

export default {
  props: {
    passAssociatedApplianceIdToComposeMemoryButton: String,
    passBladeIdToComposeMemoryButton: String,
  },

  data() {
    return {
      composeMemoryProgressText: "Composing Memory, please wait...",

      // The rules for the input fields when adding a new appliance/blade
      rules: {
        required: (value: any) => !!value || "Field is required",
      },

      dialogComposeMemorySuccess: false,
      dialogComposeMemoryFailed: false,
      dialogComposeMemory: false,
      dialogComposeMemoryAlert: false,
      dialogComposeMemoryWait: false,

      associatedApplianceId: this
        .passAssociatedApplianceIdToComposeMemoryButton as string,
      bladeId: this.passBladeIdToComposeMemoryButton as string,

      memorySizeGiB: 8,
      newMemoryCredentials: {
        port: "", // port is optional
        memorySizeMiB: 0,
        QoS: Qos.NUMBER_1,
      },
      qoSs: [Qos.NUMBER_1, Qos.NUMBER_2, Qos.NUMBER_4, Qos.NUMBER_8],
      showQoS: Qos.NUMBER_1,
      showPort: "",
      partialSuccess: "",

      newMemoryId: "", // Be used on success popup
      memorySizeNotEqual: "",
      newMemorySize: null as unknown as number | undefined,
      composeMemoryError: null as unknown, // Be used on failure popup
    };
  },

  setup() {
    const bladePortStore = useBladePortStore();
    const bladeMemoryStore = useBladeMemoryStore();

    // Computed property to get all port IDs
    const allPortIds = computed(() =>
      bladePortStore.bladePorts.map((port) => port.id)
    );

    // Computed property to get the IDs of assigned ports
    const assignedPortIds = computed(() =>
      bladeMemoryStore.bladeMemory
        .filter((memoryRegion) => memoryRegion.memoryAppliancePort)
        .map((memoryRegion) => memoryRegion.memoryAppliancePort)
    );

    // Computed property to get the unassigned port IDs
    const unassignedPortIds = computed(() => {
      const assignedIdsSet = new Set(assignedPortIds.value);
      return allPortIds.value.filter((id) => !assignedIdsSet.has(id));
    });

    // Add an empty string to allow the user to select 'none' when composing memory
    const PortIds = computed(() => ["", ...unassignedPortIds.value]);

    return {
      PortIdArray: PortIds,
    };
  },

  methods: {
    // Warning to power down the CXL-Host device if being assigned memory from a memory appliance.
    warningBeforeComposeMemory() {
      this.dialogComposeMemory = false;
      if (this.showPort) {
        this.dialogComposeMemoryAlert = true;
      } else {
        this.composeMemoryConfirm();
      }
    },

    // Trigger the API bladesComposeMemory to compose memory
    async composeMemoryConfirm() {
      this.newMemorySize = 0;
      this.memorySizeNotEqual = "";

      this.dialogComposeMemoryAlert = false;
      this.dialogComposeMemoryWait = true;

      const bladeMemoryStore = useBladeMemoryStore();

      // Handle the new memory credentials
      this.newMemoryCredentials.memorySizeMiB = this.memorySizeGiB * 1024;
      this.newMemoryCredentials.QoS = this.showQoS;
      this.newMemoryCredentials.port = this.showPort;

      // Reset showPort
      this.showPort = "";

      const newMemoryRegion = await bladeMemoryStore.composeMemory(
        this.associatedApplianceId,
        this.bladeId,
        this.newMemoryCredentials
      );
      this.newMemoryId = newMemoryRegion?.id + "";

      this.composeMemoryError = bladeMemoryStore.composeMemoryError;

      // Display the success popup and update resources, ports and memory information if compose memory succeeded
      if (!this.composeMemoryError) {
        this.partialSuccess = "";

        if (
          !newMemoryRegion?.memoryAppliancePort &&
          this.newMemoryCredentials.port
        ) {
          this.partialSuccess =
            "Note: Memory allocation succeeded but memory port assignment failed.";
        }
        // Get the size of the new memory chunk to update the blade memory,
        // The final allocated memory size may be changed by the compose memory algorithm, so it may not be the same with the input one
        this.newMemorySize = newMemoryRegion?.sizeMiB;

        const bladeResourceStore = useBladeResourceStore();
        const bladePortStore = useBladePortStore();

        await bladeResourceStore.fetchMemoryResources(
          this.associatedApplianceId,
          this.bladeId
        );
        await bladePortStore.fetchBladePorts(
          this.associatedApplianceId,
          this.bladeId
        );

        // Update the blade memory if compose memory succeeded
        const bladeStore = useBladeStore();
        let newAvailableMemory: number | undefined = undefined;
        let newAllocatedMmeory: number | undefined = undefined;
        // Must handle the scenario where selectedBladeTotalMemoryAvailableMiB or selectedBladeTotalMemoryAllocatedMiB does not exist(value is 0).
        if (this.newMemorySize !== undefined) {
          if (
            bladeStore.selectedBladeTotalMemoryAvailableMiB &&
            bladeStore.selectedBladeTotalMemoryAllocatedMiB
          ) {
            newAvailableMemory =
              bladeStore.selectedBladeTotalMemoryAvailableMiB -
              this.newMemorySize * 1024;
            newAllocatedMmeory =
              bladeStore.selectedBladeTotalMemoryAllocatedMiB +
              this.newMemorySize * 1024;
          } else if (bladeStore.selectedBladeTotalMemoryAvailableMiB) {
            newAvailableMemory =
              bladeStore.selectedBladeTotalMemoryAvailableMiB -
              this.newMemorySize * 1024;
            newAllocatedMmeory = 0 + this.newMemorySize * 1024;
          }
        }

        await bladeStore.updateSelectedBladeMemory(
          newAvailableMemory,
          newAllocatedMmeory
        );
        this.dialogComposeMemoryWait = false;
        if (this.memorySizeGiB != this.newMemorySize) {
          this.memorySizeNotEqual =
            "Note: Memory size adjustment has occurred.  The requested QoS of " +
            this.showQoS +
            " requires a minimum of " +
            this.newMemorySize +
            " GiB.";
        }
        this.dialogComposeMemorySuccess = true;
      } else {
        this.dialogComposeMemoryWait = false;
        this.dialogComposeMemoryFailed = true;
      }
    },
  },
};
</script>