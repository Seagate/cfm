<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container style="padding: 0">
    <v-data-table
      :headers="headers"
      fixed-header
      height="240"
      :items="selectedBladeMemory"
    >
      <template v-slot:[`item.actions`]="{ item }">
        <v-icon size="small" @click="assignOrUnassign(item)">
          mdi-pencil
        </v-icon>
        <span class="ml-2"></span>
        <v-icon size="small" @click="freeMemory(item)"> mdi-delete </v-icon>
        <v-tooltip activator="parent" location="end"
          >Click free button to free this memory region or click pencil button
          to assign/unassign this memory region.
        </v-tooltip>
      </template>
    </v-data-table>

    <v-dialog v-model="dialogAssignUnassign" max-width="600px">
      <v-card>
        <v-alert color="warning" icon="$warning" title="Alert" variant="tonal">
          Due to limited protections, the CXL-Host <strong> MUST </strong> be
          powered down when being
          <strong> {{ this.operation }}ed </strong> memory.
        </v-alert>
        <v-divider></v-divider>
        <v-card-text>
          <div v-if="assign">
            To assign <strong>{{ this.selectedMemoryRegion.id }}</strong> to a
            port, please <strong> power down </strong> the to be connected
            CXL-Host and select the port from the dropdown, then click the green
            button.
          </div>
          <div v-else>
            To unassign
            <strong>{{ this.selectedMemoryRegion.id }}</strong> from
            <strong>{{ this.selectedMemoryRegion.memoryAppliancePort }}</strong
            >, please click the green button after
            <strong> powering down </strong> the connected CXL-Host.
          </div>
        </v-card-text>
        <v-autocomplete
          v-if="assign"
          v-model="assignPort"
          id="inputSelectedPort"
          label="Assign to Port"
          :items="PortIdArray"
        ></v-autocomplete>
        <v-divider></v-divider>
        <v-card-action>
          <v-spacer></v-spacer>
          <div class="text-end">
            <v-btn
              color="yellow-darken-4"
              variant="text"
              @click="dialogAssignUnassign = false"
              id="cancelAssignOrUnassign"
              >cancel</v-btn
            >
            <v-btn
              color="#6ebe4a"
              variant="text"
              @click="assignUnassignPort(this.operation)"
              id="confirmAssignOrUnassign"
              >{{ assign ? "Assign Memory" : "UnAssign Memroy" }}</v-btn
            >
          </div>
        </v-card-action>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogNoPortForAssignMemory" max-width="600px">
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
        <h2 class="text-h5 mb-6">No available ports</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          All ports are assigned, please unassign one.
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="noPortForAssignMemory"
            @click="dialogNoPortForAssignMemory = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <!-- The dialog of asking the user to wait to assign or unassign memory -->
    <v-dialog v-model="waitAssignUnassignMemory">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div v-if="assign">{{ assignMemoryProgressText }}</div>
              <div v-else>{{ unassignMemoryProgressText }}</div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <v-dialog v-model="assignUnassignSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">{{ this.operation }} memory succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          Memory ID:
          <br />{{ this.selectedMemoryRegion.id }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="assignOrUnassignSuccess"
            @click="assignUnassignSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="assignUnassignFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">{{ this.operation }} memory failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ this.assignUnassignMemoryError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="assignOrUnassignFailure"
            @click="assignUnassignFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogFreeMemory" max-width="600px">
      <v-card>
        <v-alert
          color="warning"
          icon="$warning"
          title="Alert"
          variant="tonal"
          text="Due to limited protections, the CXL-Host MUST be powered down when being unassigned memory from a memory appliance."
        ></v-alert>
        <v-card-text>
          Please <strong>power down </strong>the connected CXL-Host device
          before clicking the <strong>FREE</strong> button to free the memory:
          <strong>{{ this.selectedMemoryRegion.id }}</strong>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="cancelFreeMemory"
            @click="dialogFreeMemory = false"
            >Cancel</v-btn
          >
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="confirmFreeMemory"
            @click="freeMemoryRegionConfirm"
            >Free</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogFreeMemoryWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ freeMemoryProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <v-dialog v-model="dialogFreeMemorySuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">Free memory succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          Memory ID:
          <br />{{ this.selectedMemoryRegion.id }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="freeMemorySuccess"
            @click="dialogFreeMemorySuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogFreeMemoryFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">Free memory failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ this.freeMemoryError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="freeMemoryFailure"
            @click="dialogFreeMemoryFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>
  </v-container>
</template>

<script>
import { computed } from "vue";
import { useBladeMemoryStore } from "../Stores/BladeMemoryStore";
import { useBladePortStore } from "../Stores/BladePortStore";
import { useBladeResourceStore } from "../Stores/BladeResourceStore";
import { useBladeStore } from "../Stores/BladeStore";

export default {
  data() {
    return {
      assignMemoryProgressText: "Assigning Memory, please wait...",
      unassignMemoryProgressText: "Unassigning Memory, please wait...",
      freeMemoryProgressText: "Freeing Memory, please wait...",

      headers: [
        {
          title: "MemoryId",
          align: "start",
          key: "id",
        },
        { title: "SizeGiB", key: "sizeMiB" },
        { title: "AppliancePort", key: "memoryAppliancePort" },
        { title: "Actions", key: "actions", sortable: false },
      ],

      selectedMemoryRegion: null,

      waitAssignUnassignMemory: false,

      dialogAssignUnassign: false,
      // Used to dynamically change the content of the dialogAssignUnassign
      assign: false,

      // Credentials for assign or unassign memory
      assignUnassignMemory: {
        port: "",
        operation: "",
      },
      // operation represents the operation of assign memory or unassign memory
      operation: "",

      assignUnassignSuccess: false,
      assignUnassignFailure: false,
      assignUnassignMemoryError: null, // To be used in failure popup
      assignOrUnassignResponse: null, // To be used in success popup
      dialogNoPortForAssignMemory: false,

      dialogFreeMemory: false,
      dialogFreeMemoryWait: false,
      freeMemoryError: null, // To be used in failure popup
      freeMemoryResponse: null, // To be used in success popup
      dialogFreeMemorySuccess: false,
      dialogFreeMemoryFailure: false,
    };
  },

  methods: {
    //  Open the assign or unassign memory dialog
    assignOrUnassign(item) {
      this.selectedMemoryRegion = item;

      // If the selectedMemoryRegion has a memoryAppliancePort assigned to it, it must be unassigned.
      // Otherwise assign it to a specific port
      if (this.selectedMemoryRegion.memoryAppliancePort) {
        this.operation = "unassign";
        this.assign = false;
        this.dialogAssignUnassign = true;
      } else {
        // If there are no available ports, remind the user and never display the assign memory popup
        if (this.LengthOfPorts == 0) {
          this.dialogNoPortForAssignMemory = true;
        } else {
          this.operation = "assign";
          this.assign = true;
          this.dialogAssignUnassign = true;
        }
      }
    },

    // Assign/unassign memory to/from a specific port
    async assignUnassignPort(operation) {
      const bladeMemoryStore = useBladeMemoryStore();

      this.dialogAssignUnassign = false;

      if (this.selectedMemoryRegion) {
        this.waitAssignUnassignMemory = true;

        // If the memory region is assigned, use the assigned port to unassign
        if (this.selectedMemoryRegion.memoryAppliancePort) {
          this.assignUnassignMemory.port =
            this.selectedMemoryRegion.memoryAppliancePort;
        } else {
          // If the memory region is unassigned, use the input port to assign if it is not empty
          if (this.assignPort) {
            this.assignUnassignMemory.port = this.assignPort;
            // Reset assignPort
            this.assignPort = "";
            // If the input port is empty, stop assigning memory
          } else {
            this.assignUnassignMemoryError = "Assign port need to be selected.";
            this.waitAssignUnassignMemory = false;
            this.assignUnassignFailure = true;
            return;
          }
        }

        this.assignUnassignMemory.operation = operation;

        // Talk with the blade memory store to assign or unassign memory
        this.assignOrUnassignResponse = await bladeMemoryStore.assignOrUnassign(
          this.selectedMemoryRegion.memoryApplianceId,
          this.selectedMemoryRegion.memoryBladeId,
          this.selectedMemoryRegion.id,
          this.assignUnassignMemory
        );
        this.assignUnassignMemoryError =
          bladeMemoryStore.assignOrUnassignMemoryError;

        // Update resources, ports and memory table and show the success popup if assign/unassign memory succeeded
        if (this.assignOrUnassignResponse) {
          const bladeResourceStore = useBladeResourceStore();
          const bladePortStore = useBladePortStore();
          await bladeMemoryStore.fetchBladeMemory(
            this.selectedMemoryRegion.memoryApplianceId,
            this.selectedMemoryRegion.memoryBladeId
          );
          await bladeResourceStore.updateMemoryResourcesStatus(
            this.selectedMemoryRegion.memoryApplianceId,
            this.selectedMemoryRegion.memoryBladeId
          );
          await bladePortStore.fetchBladePorts(
            this.selectedMemoryRegion.memoryApplianceId,
            this.selectedMemoryRegion.memoryBladeId
          );

          this.waitAssignUnassignMemory = false;
          this.assignUnassignSuccess = true;
        } else {
          this.waitAssignUnassignMemory = false;
          this.assignUnassignFailure = true;
        }
      }
    },

    //  Open the free memory dialog
    freeMemory(item) {
      this.selectedMemoryRegion = item;
      this.dialogFreeMemory = true;
    },

    async freeMemoryRegionConfirm() {
      this.dialogFreeMemory = false;
      this.dialogFreeMemoryWait = true;

      const bladeMemoryStore = useBladeMemoryStore();
      if (this.selectedMemoryRegion) {
        this.freeMemoryResponse = await bladeMemoryStore.freeMemory(
          this.selectedMemoryRegion.memoryApplianceId,
          this.selectedMemoryRegion.memoryBladeId,
          this.selectedMemoryRegion.id
        );
        this.freeMemoryError = bladeMemoryStore.freeMemoryError;
      }

      // Update memory allocation chart, resources, memory, ports table and show the success popup if free memory succeeded
      if (!this.freeMemoryError) {
        const bladeResourceStore = useBladeResourceStore();
        const bladePortStore = useBladePortStore();

        await bladeResourceStore.updateMemoryResourcesStatus(
          this.selectedMemoryRegion.memoryApplianceId,
          this.selectedMemoryRegion.memoryBladeId
        );
        await bladePortStore.fetchBladePorts(
          this.selectedMemoryRegion.memoryApplianceId,
          this.selectedMemoryRegion.memoryBladeId
        );

        // Update the blade memory
        const bladeStore = useBladeStore();
        let newAvailableMemory;
        let newAllocatedMmeory;
        // Must handle the scenario where selectedBladeTotalMemoryAvailableMiB or selectedBladeTotalMemoryAllocatedMiB does not exist(value is 0).
        if (
          bladeStore.selectedBladeTotalMemoryAvailableMiB &&
          bladeStore.selectedBladeTotalMemoryAllocatedMiB
        ) {
          newAvailableMemory =
            bladeStore.selectedBladeTotalMemoryAvailableMiB +
            this.selectedMemoryRegion.sizeMiB * 1024;
          newAllocatedMmeory =
            bladeStore.selectedBladeTotalMemoryAllocatedMiB -
            this.selectedMemoryRegion.sizeMiB * 1024;
        } else if (bladeStore.selectedBladeTotalMemoryAllocatedMiB) {
          newAvailableMemory = this.selectedMemoryRegion.sizeMiB * 1024;
          newAllocatedMmeory =
            bladeStore.selectedBladeTotalMemoryAllocatedMiB -
            this.selectedMemoryRegion.sizeMiB * 1024;
        }
        await bladeStore.updateSelectedBladeMemory(
          newAvailableMemory,
          newAllocatedMmeory
        );

        this.dialogFreeMemoryWait = false;
        this.dialogFreeMemorySuccess = true;
      } else {
        this.dialogFreeMemoryWait = false;
        this.dialogFreeMemoryFailure = true;
      }
    },
  },

  setup() {
    const bladeMemoryStore = useBladeMemoryStore();

    // Computed property to sort memory by the numerical part of the MemoryId
    const sortedBladeMemory = computed(() => {
      return bladeMemoryStore.bladeMemory
        .slice() // Create a copy to avoid mutating the original array
        .sort((a, b) => {
          // Extract the numerical part from the MemoryId
          const numA = parseInt(a.id.replace(/^\D+/g, ""));
          const numB = parseInt(b.id.replace(/^\D+/g, ""));
          return numA - numB;
        });
    });

    const bladePortStore = useBladePortStore();

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

    // Returns the length of unassigned ports and keeps it updated for the reminder of no available ports.
    const lengthOfPorts = computed(() => {
      return unassignedPortIds.value.length;
    });

    return {
      selectedBladeMemory: sortedBladeMemory,
      PortIdArray: unassignedPortIds,
      LengthOfPorts: lengthOfPorts,
    };
  },
};
</script>
