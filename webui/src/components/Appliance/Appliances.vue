<template>
  <v-container>
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

    <div>
      <v-tabs
        v-model="selectedApplianceId"
        color="#6ebe4a"
        bg-color="rgba(110, 190, 74, 0.1)"
      >
        <v-tab>
          <v-btn
            variant="text"
            id="addAppliance"
            @click="addNewApplianceWindowButton"
          >
            <v-icon start color="#6ebe4a">mdi-plus-thick</v-icon>
            APPLIANCE
            <v-tooltip activator="parent" location="end"
              >Click here to add new memory appliance</v-tooltip
            >
          </v-btn>
        </v-tab>

        <v-tab
          v-for="appliance in appliances"
          :value="appliance.id"
          :key="appliance.id"
          :id="appliance.id"
          @click="selectAppliance(appliance.id)"
        >
          <v-row justify="space-between" align="center">
            <v-col> {{ appliance.id }} </v-col>
            <v-col>
              <v-btn icon variant="text">
                <v-icon
                  size="x-small"
                  color="warning"
                  @click="deleteApplianceWindowButton"
                  id="deleteApplianceWindow"
                  >mdi-close</v-icon
                >
                <v-tooltip activator="parent" location="end"
                  >Click here to delete this memory appliance</v-tooltip
                >
              </v-btn>
            </v-col>
          </v-row>
        </v-tab>
      </v-tabs>
      <v-window v-model="selectedApplianceId">
        <v-window-item
          v-for="appliance in appliances"
          :key="appliance.id"
          :value="appliance.id"
        >
          <v-tabs
            v-model="selectedBladeId"
            color="#6ebe4a"
            bg-color="rgba(110, 190, 74, 0.1)"
          >
            <v-tab>
              <v-btn
                variant="text"
                id="addBlade"
                @click="addNewBladeWindowButton"
              >
                <v-icon start color="#6ebe4a">mdi-plus-thick</v-icon>
                BLADE
                <v-tooltip activator="parent" location="end"
                  >Click here to add new blade</v-tooltip
                >
              </v-btn>
            </v-tab>

            <v-tab
              v-for="blade in blades"
              :value="blade.id"
              :key="blade.id"
              :id="blade.id"
              @click="
                selectBlade(
                  blade.id,
                  blade.ipAddress,
                  blade.port,
                  Number(blade.totalMemoryAvailableMiB),
                  Number(blade.totalMemoryAllocatedMiB)
                )
              "
            >
              <v-row justify="space-between" align="center">
                <v-col> {{ blade.id }} </v-col>
                <v-col>
                  <v-btn icon variant="text">
                    <v-icon
                      color="warning"
                      size="x-small"
                      @click="deleteBladeWindowButton"
                      id="deleteBladeWindowButton"
                      >mdi-close</v-icon
                    >
                    <v-tooltip activator="parent" location="end"
                      >Click here to delete this blade</v-tooltip
                    >
                  </v-btn>
                </v-col>
              </v-row>
            </v-tab>
          </v-tabs>

          <v-window v-model="selectedBladeId">
            <v-window-item
              v-for="blade in blades"
              :value="blade.id"
              :key="blade.id"
            >
              <!-- Content for the selected blade -->
              <!-- ---------------------------------------------- -->
              <!---First Row -->
              <!-- ---------------------------------------------- -->
              <v-row class="flex-0" dense>
                <v-col cols="12" md="4" xl="2">
                  <!-- Basic Information -->
                  <v-card
                    class="card-shadow"
                    height="420"
                    color="rgba(110, 190, 74, 0.1)"
                  >
                    <v-toolbar height="45">
                      <v-toolbar-title style="cursor: pointer"
                        >Basic Information</v-toolbar-title
                      >
                      <v-btn
                        icon="mdi-help-circle"
                        size="x-small"
                        id="basicInfoExplanation"
                        @click="dialogBasicInfoExplanation = true"
                      ></v-btn>
                    </v-toolbar>
                    <v-card-text>
                      <h2 class="text-h6 text-green-lighten-2">Blade</h2>
                      A blade is associated with one appliance and it is Redfish
                      service running on OpenBMC.
                    </v-card-text>
                    <v-list lines="one">
                      <v-list-item>
                        <v-list-item-title prepend-icon="mdi-account-circle"
                          >Appliance Id</v-list-item-title
                        >
                        <v-list-item-subtitle>
                          {{ selectedApplianceId }}
                        </v-list-item-subtitle>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Blade Id</v-list-item-title>
                        <v-list-item-subtitle>
                          {{ selectedBladeId }}
                        </v-list-item-subtitle>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Ip Address</v-list-item-title>
                        <v-list-item-subtitle>
                          {{ selectedBladeIp + ":" + selectedBladePort }}
                        </v-list-item-subtitle>
                      </v-list-item>
                    </v-list>
                  </v-card>
                </v-col>
                <v-col cols="12" md="8" xl="4">
                  <!-- Memory Management -->
                  <v-card class="card-shadow" height="420">
                    <v-toolbar height="45">
                      <v-toolbar-title style="cursor: pointer"
                        >Memory</v-toolbar-title
                      >
                      <v-btn
                        icon="mdi-help-circle"
                        size="x-small"
                        id="memoryInfoExplanation"
                        @click="dialogMemoryInfoExplanation = true"
                      ></v-btn>
                    </v-toolbar>
                    <v-card-text
                      class="d-flex align-center"
                      style="max-height: 420px; overflow-y: auto"
                    >
                      <MemoryChart />
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" xl="6">
                  <!-- Ports-->
                  <v-card class="card-shadow h-full" height="420">
                    <v-toolbar height="45">
                      <v-toolbar-title style="cursor: pointer"
                        >Ports</v-toolbar-title
                      >
                      <v-btn
                        icon="mdi-help-circle"
                        size="x-small"
                        id="portExplanation"
                        @click="dialogPortExplanation = true"
                      ></v-btn>
                    </v-toolbar>
                    <v-card-text style="max-height: 420px; overflow-y: auto">
                      <BladePorts />
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
              <!-- ---------------------------------------------- -->
              <!---Second Row -->
              <!-- ---------------------------------------------- -->
              <v-row class="card-shadow flex-grow-0" dense>
                <v-col cols="12" xl="6">
                  <!-- Memory Region-->
                  <v-card class="card-shadow h-full" height="420">
                    <v-toolbar height="45">
                      <v-toolbar-title style="cursor: pointer"
                        >Memory Region and Management</v-toolbar-title
                      >
                      <v-btn
                        icon="mdi-help-circle"
                        size="x-small"
                        id="memoryRegionExplanation"
                        @click="dialogMemoryRegionExplanation = true"
                      ></v-btn>
                    </v-toolbar>
                    <v-card-text style="max-height: 420px; overflow-y: auto">
                      <v-row>
                        <v-col cols="12" sm="3" xl="3">
                          <!-- Pass the selected appliance id and selected blade id to ComposeMemoryButton -->
                          <ComposeMemoryButton
                            :passAssociatedApplianceIdToComposeMemoryButton="
                              selectedApplianceId
                            "
                            :passBladeIdToComposeMemoryButton="selectedBladeId"
                          />
                        </v-col>
                        <v-col>
                          <v-card-text class="text-h9 text-grey-darken-1"
                            >👈 Click the left
                            <strong>Compose Memory</strong> button to
                            allocate/assign memory for your cxl-host
                            device.</v-card-text
                          >
                        </v-col>
                      </v-row>
                      <v-divider></v-divider>
                      <BladeMemory />
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col cols="12" xl="6">
                  <!-- Resources-->
                  <v-card class="card-shadow h-full" height="420">
                    <v-toolbar height="45">
                      <v-toolbar-title style="cursor: pointer"
                        >Resource Blocks</v-toolbar-title
                      >
                      <v-btn
                        icon="mdi-help-circle"
                        size="x-small"
                        id="resourceExplanation"
                        @click="dialogResourceExplanation = true"
                      ></v-btn>
                    </v-toolbar>
                    <v-card-text style="max-height: 420px; overflow-y: auto">
                      <MemoryResources />
                    </v-card-text>
                  </v-card>
                </v-col>
              </v-row>
            </v-window-item>
          </v-window>
        </v-window-item>
      </v-window>
    </div>

    <!-- The dialog for adding a new appliance -->
    <v-dialog
      v-model="dialogNewAppliance"
      @keyup.enter="addNewApplianceConfirm"
      max-width="600px"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">New Appliance</span>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-container>
            <v-row class="justify-center">
              <v-col>
                <div style="position: relative; width: 100%">
                  <v-text-field
                    v-model="newApplianceCredentials.customId"
                    label="Appliance Name (Optional)"
                    :style="{ width: '100%' }"
                    id="inputCustomedApplianceId"
                  ></v-text-field>
                </div>
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
            id="cancelAddNewAppliance"
            @click="dialogNewAppliance = false"
          >
            Cancel
          </v-btn>
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="confirmAddNewAppliance"
            @click="addNewApplianceConfirm"
          >
            Add
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogAddApplianceSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">You added an appliance successfully!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New Appliance Id:
          <br />{{ newApplianceId }}
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
            @click="dialogAddApplianceSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogAddApplianceFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">You added new appliance failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ addApplianceError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="addNewApplianceFailure"
            @click="dialogAddApplianceFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogAddApplianceWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ addApplianceProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <!-- The dialog for deleting an appliance -->
    <v-dialog v-model="dialogDeleteAppliance" max-width="600px">
      <v-card>
        <v-alert
          color="warning"
          icon="$warning"
          title="Alert"
          variant="tonal"
          text="Delete appliance? The action cannot be undone."
        ></v-alert>
        <v-card-text>
          <div class="text-h6 pa-12">{{ selectedApplianceId }}</div>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="cancelDeleteAppliance"
            @click="dialogDeleteAppliance = false"
            >Cancel</v-btn
          >
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="confrimDeleteAppliance"
            @click="deleteApplianceConfirm(selectedApplianceId)"
            >Delete</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogDeleteApplianceSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">You deleted appliance successfully!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          Deleted Appliance Id:
          <br />{{ deletedApplianceId }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="deleteApplianceSuccess"
            @click="dialogDeleteApplianceSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogDeleteApplianceFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">You deleted appliance failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ deleteApplianceError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="deleteApplianceFailure"
            @click="dialogDeleteApplianceFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogDeleteApplianceWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ deleteApplianceProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <!-- The dialog for adding a new blade -->
    <v-dialog
      v-model="dialogNewBlade"
      @keyup.enter="addNewBladeConfirm()"
      max-width="600px"
    >
      <v-card>
        <v-card-title>
          <span class="text-h5">Add Blade</span>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-container>
            <v-row class="justify-center">
              <v-col>
                <v-text-field
                  :rules="[rules.required]"
                  v-model="newBladeCredentials.ipAddress"
                  label="Ip Address"
                  id="inputIpAddress"
                ></v-text-field>
              </v-col>
              <v-col>
                <v-text-field
                  :rules="[rules.required]"
                  v-model.number="newBladeCredentials.port"
                  label="Port"
                  id="inputPortId"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="justify-center">
              <v-col>
                <v-text-field
                  :rules="[rules.required]"
                  v-model="newBladeCredentials.username"
                  label="Username"
                  id="inputUserName"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="justify-center">
              <v-col>
                <div style="position: relative; width: 100%">
                  <v-text-field
                    :rules="[rules.required]"
                    v-model="newBladeCredentials.password"
                    label="Password"
                    id="inputPassword"
                    :type="showBladePassword ? 'text' : 'password'"
                  ></v-text-field>
                  <v-icon
                    style="
                      position: absolute;
                      right: 8px;
                      top: 50%;
                      transform: translateY(-50%);
                    "
                    @click="showBladePassword = !showBladePassword"
                  >
                    {{ showBladePassword ? "mdi-eye" : "mdi-eye-off" }}
                  </v-icon>
                </div>
              </v-col>
            </v-row>
            <v-row class="justify-center">
              <v-col>
                <v-text-field
                  v-model="newBladeCredentials.customId"
                  label="Blade Name (Optional)"
                  id="inputCustomedBladeId"
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
            id="cancelAddBlade"
            @click="dialogNewBlade = false"
          >
            Cancel
          </v-btn>
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="confirmAddBlade"
            @click="addNewBladeConfirm()"
          >
            Add
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogAddBladeSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">You added a blade successfully!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New Blade Id:
          <br />{{ newBladeId }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="addBladeSuccess"
            @click="dialogAddBladeSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogAddBladeFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">You added new blade failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ addBladeError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="addBladeFailure"
            @click="dialogAddBladeFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogAddBladeWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ addBladeProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <!-- The dialog of the warning before the deletion of a blade by the user -->
    <v-dialog v-model="dialogDeleteBlade" max-width="600px">
      <v-card>
        <v-alert
          color="warning"
          icon="$warning"
          title="Alert"
          variant="tonal"
          text="Delete blade? The action cannot be undone."
        ></v-alert>
        <v-card-text>
          <div class="text-h6 pa-12">{{ selectedBladeId }}</div>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="cancelDeleteBlade"
            @click="dialogDeleteBlade = false"
            >Cancel</v-btn
          >
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="confirmDeleteBlade"
            @click="deleteBladeConfirm"
            >Delete</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogDeleteBladeSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">You deleted a blade successfully!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          Deleted Blade Id:
          <br />{{ deletedBladeId }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="deleteBladeSuccess"
            @click="dialogDeleteBladeSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogDeleteBladeFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">You deleted a blade failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ deleteBladeError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="deleteBladeFailure"
            @click="dialogDeleteBladeFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogDeleteBladeWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ deleteBladeProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <!-- These dialogs appear when the user clicks on the corresponding question mark. -->
    <v-dialog v-model="dialogBasicInfoExplanation" max-width="600px">
      <v-card
        text="This is the basic information of the current blade."
      ></v-card>
    </v-dialog>
    <v-dialog v-model="dialogMemoryInfoExplanation" max-width="600px">
      <v-card
        text="This chart displays the available memory and allocated memory information for the current blade, and the information is updated with the compose or free memory action(s)."
      ></v-card>
    </v-dialog>
    <v-dialog v-model="dialogPortExplanation" max-width="600px">
      <v-card
        text="This table displays the ports information for the current blade, and the information is updated with the compose or free memory action(s)."
      ></v-card>
    </v-dialog>
    <v-dialog v-model="dialogMemoryRegionExplanation" max-width="600px">
      <v-card
        text="This table shows the memory chunks information of the current blade, the user is allowed to compose memory and free memory here, the corresponding memory management pie chart, ports information and resources information are constantly updated."
      ></v-card>
    </v-dialog>
    <v-dialog v-model="dialogResourceExplanation" max-width="600px">
      <v-card
        text="This table shows the resource blocks information of the current blade, the user can check the detail information of each resource block and the information is updated with the compose or free memory action(s)."
      ></v-card>
    </v-dialog>
  </v-container>
</template>

<script lang="ts">
import { watch, computed, onMounted, ref } from "vue";
import { useApplianceStore } from "../Stores/ApplianceStore";
import { useBladeStore } from "../Stores/BladeStore";
import { useBladeResourceStore } from "../Stores/BladeResourceStore";
import { useBladePortStore } from "../Stores/BladePortStore";
import MemoryChart from "./MemoryChart.vue";
import MemoryResources from "./Resources.vue";
import BladePorts from "./Ports.vue";
import BladeMemory from "./Memory.vue";
import ComposeMemoryButton from "./ComposeMemoryButton.vue";
import { useBladeMemoryStore } from "../Stores/BladeMemoryStore";

export default {
  data() {
    return {
      loadProgressText: "Loading the page, please wait...",
      addApplianceProgressText: "Adding appliance, please wait...",
      deleteApplianceProgressText: "Deleting appliance, please wait...",
      addBladeProgressText: "Adding blade, please wait...",
      deleteBladeProgressText: "Deleting blade, please wait...",

      // The rules for the input fields when adding a new appliance/blade
      rules: {
        required: (value: any) => !!value || "Field is required",
      },

      showBladePassword: false,

      dialogBasicInfoExplanation: false,
      dialogMemoryInfoExplanation: false,
      dialogPortExplanation: false,
      dialogMemoryRegionExplanation: false,
      dialogResourceExplanation: false,

      newApplianceCredentials: {
        username: "root",
        password: "0penBmc",
        ipAddress: "127.0.0.1",
        port: 8443,
        insecure: true,
        protocol: "https",
        customId: "",
      },
      newApplianceId: "", // Be used on success popup
      dialogNewAppliance: false,
      dialogAddApplianceWait: false,
      dialogAddApplianceSuccess: false,
      dialogAddApplianceFailure: false,
      addApplianceError: null as unknown,

      deletedApplianceId: null as unknown as string | undefined, // Be used on success popup
      dialogDeleteAppliance: false,
      deleteApplianceError: null as unknown,
      dialogDeleteApplianceSuccess: false,
      dialogDeleteApplianceFailure: false,
      dialogDeleteApplianceWait: false,

      newBladeCredentials: {
        username: "root",
        password: "0penBmc",
        ipAddress: "127.0.0.1",
        port: 443,
        insecure: true,
        protocol: "https",
        customId: "",
      },
      newBladeId: "", // Be used on success popup
      dialogNewBlade: false,
      dialogAddBladeWait: false,
      dialogAddBladeSuccess: false,
      dialogAddBladeFailure: false,
      addBladeError: null as unknown,

      deletedBladeId: null as unknown as string | undefined, // Be used on success popup
      dialogDeleteBlade: false,
      dialogDeleteBladeWait: false,
      dialogDeleteBladeSuccess: false,
      dialogDeleteBladeFailure: false,
      deleteBladeError: null as unknown,
    };
  },

  // The child components
  components: {
    MemoryChart,
    MemoryResources,
    BladePorts,
    BladeMemory,
    ComposeMemoryButton,
  },

  methods: {
    /* Open the add appliance popup */
    addNewApplianceWindowButton() {
      this.dialogNewAppliance = true;
    },

    /* Triggle the API appliancesPost in appliance store to add a new appliance */
    async addNewApplianceConfirm() {
      // Make the add appliance popup disappear
      this.dialogNewAppliance = false;

      this.dialogAddApplianceWait = true;
      const applianceStore = useApplianceStore();

      const newAppliance = await applianceStore.addNewAppliance(
        this.newApplianceCredentials
      );
      this.addApplianceError = applianceStore.addApplianceError as string;

      // Display success  popup once adding new appliance successfully
      if (!this.addApplianceError) {
        this.newApplianceId = newAppliance?.id + "";

        // Set the new added appliance as the selected appliance
        const appliances = computed(() => applianceStore.appliances);
        if (appliances.value.length > 0) {
          applianceStore.selectAppliance(this.newApplianceId);
        }
        this.dialogAddApplianceWait = false;
        this.dialogAddApplianceSuccess = true;
      } else {
        this.dialogAddApplianceWait = false;
        this.dialogAddApplianceFailure = true;
      }
      // Reset the customId to empty
      this.newApplianceCredentials.customId = "";
    },

    /* Open the delete appliance popup */
    deleteApplianceWindowButton() {
      this.dialogDeleteAppliance = true;
    },

    /* Trigger the API to delete the appliance */
    async deleteApplianceConfirm(selectedAppliance: string) {
      // Make the delete appliance popup disappear
      this.dialogDeleteAppliance = false;

      this.dialogDeleteApplianceWait = true;

      const applianceStore = useApplianceStore();

      const deletedAppliance = await applianceStore.deleteAppliance(
        selectedAppliance
      );

      this.deleteApplianceError = applianceStore.deleteApplianceError;

      // Update the appliances and set the default selected appliance
      if (!this.deleteApplianceError) {
        this.deletedApplianceId = deletedAppliance;
        const appliances = computed(() => applianceStore.appliances);

        // Check if there are any appliances left after deletion
        if (appliances.value.length > 0) {
          // Set the first appliance as selected
          const selectedApplianceId = appliances.value[0].id;
          applianceStore.selectAppliance(selectedApplianceId);
        }

        this.dialogDeleteApplianceWait = false;
        this.dialogDeleteApplianceSuccess = true;
      } else {
        this.dialogDeleteApplianceWait = false;
        this.dialogDeleteApplianceFailure = true;
      }
    },

    /* Open the add blade popup */
    addNewBladeWindowButton() {
      this.dialogNewBlade = true;
    },

    /* Triggle the API bladesPost in blade store to add a new blade */
    async addNewBladeConfirm() {
      // Make the add blade popup disappear
      this.dialogNewBlade = false;

      this.dialogAddBladeWait = true;

      const bladeStore = useBladeStore();
      const newBlade = await bladeStore.addNewBlade(
        this.selectedApplianceId,
        this.newBladeCredentials
      );

      this.addBladeError = bladeStore.addBladeError;

      // Display success  popup once adding new blade successfully
      if (!this.addBladeError) {
        this.newBladeId = newBlade!.id;

        // Set the new added blade as the selected one
        const blades = computed(() => bladeStore.blades);
        if (blades.value.length > 0) {
          const defaultBlade = newBlade;
          bladeStore.selectBlade(
            defaultBlade!.id,
            defaultBlade!.ipAddress,
            defaultBlade!.port,
            Number(defaultBlade!.totalMemoryAvailableMiB),
            Number(defaultBlade!.totalMemoryAllocatedMiB)
          );
        }
        this.dialogAddBladeWait = false;
        this.dialogAddBladeSuccess = true;
      } else {
        this.dialogAddBladeWait = false;
        this.dialogAddBladeFailure = true;
      }

      // Reset the Credentials
      this.newBladeCredentials = {
        username: "root",
        password: "0penBmc",
        ipAddress: "127.0.0.1",
        port: 443,
        insecure: true,
        protocol: "https",
        customId: "",
      };
    },

    /* Open the delete blade popup */
    deleteBladeWindowButton() {
      this.dialogDeleteBlade = true;
    },

    /* Trigger the API bladesDeleteById in blade store to delete the blade */
    async deleteBladeConfirm() {
      // Make the delete blade popup disappear
      this.dialogDeleteBlade = false;

      this.dialogDeleteBladeWait = true;

      const bladeStore = useBladeStore();

      const deletedBlade = await bladeStore.deleteBlade(
        this.selectedApplianceId,
        this.selectedBladeId
      );

      this.deleteBladeError = bladeStore.deleteBladeError;

      // Update the blades and set the default selected blade
      if (!this.deleteBladeError) {
        this.deletedBladeId = deletedBlade;

        const blades = computed(() => bladeStore.blades);
        // Check if there are any blades associated the selected appliance left after deletion
        if (blades.value.length > 0) {
          // Set the first blade as selected
          const defaultBlade = blades.value[0];
          bladeStore.selectBlade(
            defaultBlade.id,
            defaultBlade.ipAddress,
            defaultBlade.port,
            Number(defaultBlade.totalMemoryAvailableMiB),
            Number(defaultBlade.totalMemoryAllocatedMiB)
          );
        }
        this.dialogDeleteBladeWait = false;
        this.dialogDeleteBladeSuccess = true;
      } else {
        this.dialogDeleteBladeWait = false;
        this.dialogDeleteBladeFailure = true;
      }
    },
  },

  setup() {
    // Set up loading for progress linear
    const loading = ref(false);

    const applianceStore = useApplianceStore();
    const bladeStore = useBladeStore();
    const bladeResourceStore = useBladeResourceStore();
    const bladePortStore = useBladePortStore();
    const bladeMemoryStore = useBladeMemoryStore();

    // Fetch appliances when component is mounted
    onMounted(async () => {
      loading.value = true;
      await applianceStore.fetchAppliances();
      if (applianceStore.appliances.length > 0) {
        const firstApplianceId = applianceStore.appliances[0].id;
        applianceStore.selectAppliance(firstApplianceId);
      }
      loading.value = false;
    });

    // Watch for changes in selected appliance and fetch the associated blades
    watch(
      () => applianceStore.selectedApplianceId,
      async (newVal, oldVal) => {
        if (newVal !== null && newVal !== oldVal) {
          loading.value = true;
          await bladeStore.fetchBlades(newVal);
          if (bladeStore.blades.length > 0) {
            // Set the first blade for the associated appliance as the default selected blade
            const defaultBlade = bladeStore.blades[0];
            bladeStore.selectBlade(
              defaultBlade.id,
              defaultBlade.ipAddress,
              defaultBlade.port,
              Number(defaultBlade.totalMemoryAvailableMiB),
              Number(defaultBlade.totalMemoryAllocatedMiB)
            );
          }
          loading.value = false;
        }
      },
      { immediate: true }
    );

    // Watch for changes in selected blade and fetch the associated resources and ports
    watch(
      () => bladeStore.selectedBladeId,
      async (newBladeId, oldBladeId) => {
        if (newBladeId !== null && newBladeId !== oldBladeId) {
          loading.value = true;
          // Fetch resources and ports for the newly selected blade
          await Promise.all([
            bladeResourceStore.fetchMemoryResources(
              applianceStore.selectedApplianceId,
              newBladeId
            ),
            bladePortStore.fetchBladePorts(
              applianceStore.selectedApplianceId,
              newBladeId
            ),
            bladeMemoryStore.fetchBladeMemory(
              applianceStore.selectedApplianceId,
              newBladeId
            ),
          ]);
          loading.value = false;
        }
      },
      { immediate: true }
    );

    // Computed properties to access state
    const appliances = computed(() => applianceStore.appliances);
    const blades = computed(() => bladeStore.blades);
    const selectedApplianceId = computed(
      () => applianceStore.selectedApplianceId
    );
    const selectedBladeId = computed(() => bladeStore.selectedBladeId);
    const selectedBladeIp = computed(() => bladeStore.selectedBladeIp);
    const selectedBladePort = computed(() => bladeStore.selectedBladePortNum);

    // Methods to update state
    const selectAppliance = (applianceId: string) => {
      applianceStore.selectAppliance(applianceId);
    };
    const selectBlade = (
      bladeId: string,
      bladeIp: string,
      bladePort: number,
      bladeMemoryAvailable: number,
      bladeMemoryAllocated: number
    ) => {
      bladeStore.selectBlade(
        bladeId,
        bladeIp,
        bladePort,
        bladeMemoryAvailable,
        bladeMemoryAllocated
      );
    };

    return {
      appliances,
      blades,
      selectedApplianceId,
      selectedBladeId,
      selectedBladePort,
      selectedBladeIp,
      selectAppliance,
      selectBlade,
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