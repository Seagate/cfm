<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
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
              <v-menu>
                <template v-slot:activator="{ props }">
                  <v-btn
                    color="#6ebe4a"
                    icon="mdi-dots-vertical"
                    variant="text"
                    v-bind="props"
                    @click.stop="selectAppliance(appliance.id)"
                  ></v-btn>
                </template>
                <v-list>
                  <v-list-item
                    v-for="(item, i) in applianceDropItems"
                    :key="i"
                    :value="item"
                    :id="item.id"
                    @click="item.function"
                  >
                    <template v-slot:prepend>
                      <v-icon
                        :icon="item.icon"
                        :color="item.iconColor"
                      ></v-icon>
                    </template>
                    <v-list-item-title>{{ item.text }}</v-list-item-title>
                  </v-list-item>
                </v-list>
              </v-menu>
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
                :disabled="isAddBladeButtonDisabled"
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
                  Number(blade.totalMemoryAllocatedMiB),
                  blade.status
                )
              "
            >
              <v-row justify="space-between" align="center">
                <v-col> {{ blade.id }} </v-col>
                <v-col>
                  <v-menu>
                    <template v-slot:activator="{ props }">
                      <v-btn
                        color="#6ebe4a"
                        icon="mdi-dots-vertical"
                        variant="text"
                        v-bind="props"
                        @click.stop="
                          selectBlade(
                            blade.id,
                            blade.ipAddress,
                            blade.port,
                            Number(blade.totalMemoryAvailableMiB),
                            Number(blade.totalMemoryAllocatedMiB),
                            blade.status
                          )
                        "
                      ></v-btn>
                    </template>
                    <v-list>
                      <v-list-item
                        v-for="(item, i) in bladeDropItems"
                        :key="i"
                        :value="item"
                        :id="item.id"
                        @click="item.function"
                      >
                        <template v-slot:prepend>
                          <v-icon
                            :icon="item.icon"
                            :color="item.iconColor"
                          ></v-icon>
                        </template>
                        <v-list-item-title>{{ item.text }}</v-list-item-title>
                      </v-list-item>
                    </v-list>
                  </v-menu>
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
                        <v-list-item-title>Status</v-list-item-title>
                        <v-list-item-subtitle
                          :style="{ fontWeight: 'bold', color: statusColor }"
                          >{{ selectedBladeStatus }}</v-list-item-subtitle
                        >
                        <template v-slot:prepend>
                          <v-avatar>
                            <v-icon :color="statusColor">{{
                              statusIcon
                            }}</v-icon>
                          </v-avatar>
                        </template>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title prepend-icon="mdi-account-circle"
                          >Appliance Id</v-list-item-title
                        >
                        <v-list-item-subtitle>
                          {{ selectedApplianceId }}
                        </v-list-item-subtitle>
                        <template v-slot:prepend>
                          <v-avatar>
                            <v-icon color="#6ebe4a">mdi-account-circle</v-icon>
                          </v-avatar>
                        </template>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Blade Id</v-list-item-title>
                        <v-list-item-subtitle>
                          {{ selectedBladeId }}
                        </v-list-item-subtitle>
                        <template v-slot:prepend>
                          <v-avatar>
                            <v-icon color="#6ebe4a">mdi-shield-account</v-icon>
                          </v-avatar>
                        </template>
                      </v-list-item>
                      <v-list-item>
                        <v-list-item-title>Ip Address</v-list-item-title>
                        <v-list-item-subtitle>
                          {{ selectedBladeIp + ":" + selectedBladePort }}
                        </v-list-item-subtitle>
                        <template v-slot:prepend>
                          <v-avatar>
                            <v-icon color="#6ebe4a">mdi-ip</v-icon>
                          </v-avatar>
                        </template>
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
                    <v-card-text
                      style="max-height: 420px; overflow: hidden; padding: 0"
                    >
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
                    <v-card-text
                      style="max-height: 420px; overflow: hidden; padding: 0"
                    >
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
                    <v-card-text
                      style="max-height: 420px; overflow: hidden; padding: 0"
                    >
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
        <h2 class="text-h5 mb-6">Add an appliance succeeded!</h2>
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
        <h2 class="text-h5 mb-6">Add an appliance failed!</h2>
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
        <h2 class="text-h5 mb-6">Delete an appliance succeeded!</h2>
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
        <h2 class="text-h5 mb-6">Delete an appliance failed!</h2>
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

    <!-- The dialog for renaming an appliance -->
    <v-dialog v-model="dialogRenameAppliance" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="text-h5">Rename Appliance</span>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-container>
            <v-row class="justify-center">
              <v-col>
                <div style="position: relative; width: 100%">
                  <v-text-field
                    v-model="selectedApplianceId"
                    label="Current Appliance Name"
                    :style="{ width: '100%' }"
                    id="currentApplianceId"
                    readonly
                  ></v-text-field>
                  <v-text-field
                    v-model="renameApplianceCredentials.customId"
                    label="New Appliance Name"
                    :style="{ width: '100%' }"
                    id="renameApplianceId"
                    :rules="[rules.required]"
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
            id="cancelRenameAppliance"
            @click="dialogRenameAppliance = false"
          >
            Cancel
          </v-btn>
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="confirmRenameAppliance"
            @click="
              renameApplianceConfirm(
                selectedApplianceId,
                renameApplianceCredentials.customId
              )
            "
          >
            Rename
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogRenameApplianceSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">Rename an appliance succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New Appliance Id:
          <br />{{ renamedApplianceId }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="renameApplianceSuccess"
            @click="dialogRenameApplianceSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogRenameApplianceFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">Rename an appliance failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ renameApplianceError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="renameApplianceFailure"
            @click="dialogRenameApplianceFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogRenameApplianceWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ renameApplianceProgressText }}
              </div>
            </template>
          </v-progress-linear>
        </v-col>
      </v-row>
    </v-dialog>

    <!-- The dialog for renaming a blade -->
    <v-dialog v-model="dialogRenameBlade" max-width="600px">
      <v-card>
        <v-card-title>
          <span class="text-h5">Rename Blade</span>
        </v-card-title>
        <v-divider></v-divider>
        <v-card-text>
          <v-container>
            <v-row class="justify-center">
              <v-col>
                <div style="position: relative; width: 100%">
                  <v-text-field
                    v-model="selectedBladeId"
                    label="Current Blade Name"
                    :style="{ width: '100%' }"
                    id="currentBladeId"
                    readonly
                  ></v-text-field>
                  <v-text-field
                    v-model="renameBladeCredentials.customId"
                    label="New Blade Name"
                    :style="{ width: '100%' }"
                    id="renameBladeId"
                    :rules="[rules.required]"
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
            id="cancelRenameBlade"
            @click="dialogRenameBlade = false"
          >
            Cancel
          </v-btn>
          <v-btn
            color="#6ebe4a"
            variant="text"
            id="confirmRenameBlade"
            @click="
              renameBladeConfirm(
                selectedApplianceId,
                selectedBladeId,
                renameBladeCredentials.customId
              )
            "
          >
            Rename
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogRenameBladeSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">Rename a Blade succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          New Blade Id:
          <br />{{ renamedBladeId }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="renameBladeSuccess"
            @click="dialogRenameBladeSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogRenameBladeFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">Rename a Blade failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ renameBladeError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="renameBladeFailure"
            @click="dialogRenameBladeFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogRenameBladeWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ renameBladeProgressText }}
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
        <h2 class="text-h5 mb-6">Add a blade succeeded!</h2>
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
        <h2 class="text-h5 mb-6">Add a blade failed!</h2>
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
        <h2 class="text-h5 mb-6">Delete a blade succeeded!</h2>
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
        <h2 class="text-h5 mb-6">Delete a blade failed!</h2>
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

    <!-- The dialog of the warning before the resynchronizing a blade by the user -->
    <v-dialog v-model="dialogResyncBlade" max-width="600px">
      <v-card>
        <v-alert
          color="warning"
          icon="$warning"
          title="Alert"
          variant="tonal"
          text="Resynchronizing a blade deletes the blade and adds it back."
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
            id="cancelResyncBlade"
            @click="dialogResyncBlade = false"
            >Cancel</v-btn
          >
          <v-btn
            color="yellow-darken-4"
            variant="text"
            id="confirmResyncBlade"
            @click="resyncBladeConfirm(selectedApplianceId, selectedBladeId)"
            >Resync</v-btn
          >
        </v-card-actions>
      </v-card>
    </v-dialog>

    <v-dialog v-model="dialogResyncBladeSuccess" max-width="600px">
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
        <h2 class="text-h5 mb-6">Resync the blade succeeded!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          The contents of blade {{ resyncBladeId }} are updated.
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="success"
            rounded
            variant="flat"
            width="90"
            id="resyncBladeSuccess"
            @click="dialogResyncBladeSuccess = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogResyncBladeFailure" max-width="600px">
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
        <h2 class="text-h5 mb-6">Resync the blade failed!</h2>
        <p class="mb-4 text-medium-emphasis text-body-2">
          {{ resyncBladeError }}
        </p>
        <v-divider class="mb-4"></v-divider>
        <div class="text-end">
          <v-btn
            class="text-none"
            color="error"
            rounded
            variant="flat"
            width="90"
            id="resyncBladeFailure"
            @click="dialogResyncBladeFailure = false"
          >
            Done
          </v-btn>
        </div>
      </v-sheet>
    </v-dialog>

    <v-dialog v-model="dialogResyncBladeWait">
      <v-row align-content="center" class="fill-height" justify="center">
        <v-col cols="6">
          <v-progress-linear color="#6ebe4a" height="50" indeterminate rounded>
            <template v-slot:default>
              <div class="text-center">
                {{ resyncBladeProgressText }}
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
import { useRouter } from "vue-router";

export default {
  data() {
    return {
      loadProgressText: "Loading the page, please wait...",
      addApplianceProgressText: "Adding appliance, please wait...",
      deleteApplianceProgressText: "Deleting appliance, please wait...",
      addBladeProgressText: "Adding blade, please wait...",
      deleteBladeProgressText: "Deleting blade, please wait...",
      resyncBladeProgressText: "Resynchronizing the blade, please wait...",
      renameApplianceProgressText: "Renaming appliance, please wait...",
      renameBladeProgressText: "Renaming blade, please wait...",

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

      renameApplianceCredentials: {
        customId: "",
      },
      renamedApplianceId: null as unknown as string | undefined, // Be used on success popup
      dialogRenameAppliance: false,
      renameApplianceError: null as unknown,
      dialogRenameApplianceSuccess: false,
      dialogRenameApplianceFailure: false,
      dialogRenameApplianceWait: false,

      applianceDropItems: [
        {
          text: "Delete",
          icon: "mdi-delete",
          function: this.deleteApplianceWindowButton,
          id: "deleteApplianceWindow",
          iconColor: "warning",
        },
        {
          text: "Rename",
          icon: "mdi-rename-box",
          function: this.renameAppliance,
          id: "renameApplianceWindow",
          iconColor: "primary",
        },
      ],

      renameBladeCredentials: {
        customId: "",
      },
      renamedBladeId: null as unknown as string | undefined, // Be used on success popup
      dialogRenameBlade: false,
      renameBladeError: null as unknown,
      dialogRenameBladeSuccess: false,
      dialogRenameBladeFailure: false,
      dialogRenameBladeWait: false,

      bladeDropItems: [
        {
          text: "Delete",
          icon: "mdi-delete",
          function: this.deleteBladeWindowButton,
          id: "deleteBladeWindow",
          iconColor: "warning",
        },
        {
          text: "Rename",
          icon: "mdi-rename-box",
          function: this.renameBlade,
          id: "renameBladeWindow",
          iconColor: "primary",
        },
        {
          text: "Resync",
          icon: "mdi-sync-circle",
          function: this.resyncBladeWindowButton,
          id: "resyncBladeWindow",
          iconColor: "#6ebe4a",
        },
      ],

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

      resyncBladeId: null as unknown as string | undefined, // Be used on success popup
      dialogResyncBlade: false,
      dialogResyncBladeWait: false,
      dialogResyncBladeSuccess: false,
      dialogResyncBladeFailure: false,
      resyncBladeError: null as unknown,
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

  computed: {
    isAddBladeButtonDisabled() {
      return this.selectedApplianceId === "CMA_Discovered_Blades";
    },
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

      // Display success  popup once adding new appliance succeeded
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

    renameAppliance() {
      this.dialogRenameAppliance = true;
    },

    /* Triggle the API appliancesUpdateById in appliance store to rename an appliance */
    async renameApplianceConfirm(applianceId: string, newApplianceId: string) {
      // Make the rename appliance popup disappear and waiting popup appear
      this.dialogRenameAppliance = false;
      this.dialogRenameApplianceWait = true;

      const applianceStore = useApplianceStore();
      const newAppliance = await applianceStore.renameAppliance(
        applianceId,
        newApplianceId
      );

      this.renameApplianceError = applianceStore.renameApplianceError as string;

      if (!this.renameApplianceError) {
        this.renamedApplianceId = newAppliance?.id;

        // Set the renamed appliance as the selected appliance
        const appliances = computed(() => applianceStore.appliances);
        if (appliances.value.length > 0) {
          applianceStore.selectAppliance(newApplianceId);
        }

        this.dialogRenameApplianceWait = false;
        this.dialogRenameApplianceSuccess = true;
      } else {
        this.dialogRenameApplianceWait = false;
        this.dialogRenameApplianceFailure = true;
      }

      // Reset the credentials
      this.renameApplianceCredentials = {
        customId: "",
      };
    },

    renameBlade() {
      this.dialogRenameBlade = true;
    },

    /* Triggle the API bladesUpdateById in blade store to rename a blade */
    async renameBladeConfirm(
      applianceId: string,
      bladeId: string,
      newBladeId: string
    ) {
      // Make the rename blade popup disappear and waiting popup appear
      this.dialogRenameBlade = false;
      this.dialogRenameBladeWait = true;

      const bladeStore = useBladeStore();
      const newBlade = await bladeStore.renameBlade(
        applianceId,
        bladeId,
        newBladeId
      );

      this.renameBladeError = bladeStore.renameBladeError as string;

      if (!this.renameBladeError) {
        this.renamedBladeId = newBlade?.id;

        // Set the renamed Blade as the selected Blade
        const Blades = computed(() => bladeStore.blades);
        if (Blades.value.length > 0) {
          const defaultBlade = newBlade;
          bladeStore.selectBlade(
            defaultBlade!.id,
            defaultBlade!.ipAddress,
            defaultBlade!.port,
            Number(defaultBlade!.totalMemoryAvailableMiB),
            Number(defaultBlade!.totalMemoryAllocatedMiB),
            defaultBlade!.status
          );
        }

        this.dialogRenameBladeWait = false;
        this.dialogRenameBladeSuccess = true;
      } else {
        this.dialogRenameBladeWait = false;
        this.dialogRenameBladeFailure = true;
      }

      // Reset the credentials
      this.renameBladeCredentials = {
        customId: "",
      };
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

      // Display success  popup once adding new blade succeeded
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
            Number(defaultBlade!.totalMemoryAllocatedMiB),
            defaultBlade!.status
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
            Number(defaultBlade.totalMemoryAllocatedMiB),
            defaultBlade.status
          );
        }
        this.dialogDeleteBladeWait = false;
        this.dialogDeleteBladeSuccess = true;
      } else {
        this.dialogDeleteBladeWait = false;
        this.dialogDeleteBladeFailure = true;
      }
    },

    /* Open the resync blade popup */
    resyncBladeWindowButton() {
      this.dialogResyncBlade = true;
    },

    async resyncBladeConfirm(applianceId: string, bladeId: string) {
      this.dialogResyncBlade = false;
      this.dialogResyncBladeWait = true;

      const bladeStore = useBladeStore();
      await bladeStore.resyncBlade(applianceId, bladeId);

      this.resyncBladeId = bladeId;

      this.resyncBladeError = bladeStore.resyncBladeError;

      // Display the blade once resync blade succeeded
      if (!this.resyncBladeError) {
        // Manually trigger the update actions
        await this.updateBladeContent(bladeId);
        this.dialogResyncBladeWait = false;
        this.dialogResyncBladeSuccess = true;
      } else {
        this.dialogResyncBladeWait = false;
        this.dialogResyncBladeFailure = true;
      }
    },

    // Method to manually update the content for the resync blade
    async updateBladeContent(bladeId: string) {
      const bladeStore = useBladeStore();
      const bladeResourceStore = useBladeResourceStore();
      const bladePortStore = useBladePortStore();
      const bladeMemoryStore = useBladeMemoryStore();
      const applianceStore = useApplianceStore();

      await Promise.all([
        bladeStore.fetchBladeById(applianceStore.selectedApplianceId, bladeId),
        bladeResourceStore.fetchMemoryResources(
          applianceStore.selectedApplianceId,
          bladeId
        ),
        bladePortStore.fetchBladePorts(
          applianceStore.selectedApplianceId,
          bladeId
        ),
        bladeMemoryStore.fetchBladeMemory(
          applianceStore.selectedApplianceId,
          bladeId
        ),
      ]);
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

    const router = useRouter();

    // Method to update the URL
    const updateUrlWithBladeId = (applianceId: string, bladeId: string) => {
      // Construct the URL based on whether a blade ID is provided
      const newPath = `/appliances/${applianceId}/blades/${bladeId}`;
      router.push(newPath);
    };

    const updateUrlWithoutBladeId = (applianceId: string) => {
      // Construct the URL based on whether a blade ID is provided
      const newPath = `/appliances/${applianceId}`;
      router.push(newPath);
    };

    // Fetch appliances when component is mounted
    onMounted(async () => {
      loading.value = true;
      await applianceStore.fetchAppliances();
      if (applianceStore.appliances.length > 0) {
        let selectedAppliance:
          | {
              id: string;
              ipAddress?: string;
              port?: number;
              status?: string;
              blades?: { uri: string };
              totalMemoryAvailableMiB?: number;
              totalMemoryAllocatedMiB?: number;
            }
          | undefined;

        // Check if applianceId exists in the URL
        const applianceIdInUrl = router.currentRoute.value.params.appliance_id;
        if (applianceIdInUrl) {
          // Find the appliance with the applianceId from the URL
          selectedAppliance = applianceStore.appliances.find(
            (appliance) => appliance.id === applianceIdInUrl
          );
        }
        // If no applianceId in the URL or no appliance found with the applianceId, default to the first appliance
        if (!selectedAppliance) {
          selectedAppliance = applianceStore.appliances[0];
        }
        applianceStore.selectAppliance(selectedAppliance.id);
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
            let selectedBlade:
              | {
                  id: string;
                  ipAddress: string;
                  port: number;
                  status?: string;
                  ports?: { uri: string };
                  resources?: { uri: string };
                  memory?: { uri: string };
                  totalMemoryAvailableMiB?: number;
                  totalMemoryAllocatedMiB?: number;
                }
              | undefined;

            // Check if bladeId exists in the URL
            const bladeIdInUrl = router.currentRoute.value.params.blade_id;
            if (bladeIdInUrl) {
              // Find the blade with the bladeId from the URL
              selectedBlade = bladeStore.blades.find(
                (blade) => blade.id === bladeIdInUrl
              );
            }
            // If no bladeId in the URL or no blade found with the bladeId, default to the first blade
            if (!selectedBlade) {
              selectedBlade = bladeStore.blades[0];
            }
            bladeStore.selectBlade(
              selectedBlade.id,
              selectedBlade.ipAddress,
              selectedBlade.port,
              Number(selectedBlade.totalMemoryAvailableMiB),
              Number(selectedBlade.totalMemoryAllocatedMiB),
              selectedBlade.status
            );
            // Update the URL with the first blade's ID
            updateUrlWithBladeId(newVal, selectedBlade.id);
          } else {
            // If there are no blades, update the URL with only the appliance ID
            updateUrlWithoutBladeId(newVal);
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

          await bladeStore.fetchBladeById(
            applianceStore.selectedApplianceId,
            newBladeId
          );

          const selectedBlade = bladeStore.blades.find(
            (blade) => blade.id === newBladeId
          );

          if (selectedBlade) {
            bladeStore.selectBlade(
              selectedBlade.id,
              selectedBlade.ipAddress,
              selectedBlade.port,
              Number(selectedBlade.totalMemoryAvailableMiB),
              Number(selectedBlade.totalMemoryAllocatedMiB),
              selectedBlade.status
            );
          }

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
          // Update the URL with the new blade ID
          updateUrlWithBladeId(applianceStore.selectedApplianceId, newBladeId);
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
    const selectedBladeStatus = computed(() => bladeStore.selectedBladeStatus);

    const statusColor = computed(() => {
      if (selectedBladeStatus.value === "online") {
        return "#6ebe4a";
      } else if (selectedBladeStatus.value === "unavailable") {
        return "#ff9f40";
      } else if (selectedBladeStatus.value === "offline") {
        return "#b00020";
      } else {
        return "#B0B0B0"; // Default unknown status
      }
    });

    const statusIcon = computed(() => {
      if (selectedBladeStatus.value === "online") {
        return "mdi-check-circle";
      } else if (selectedBladeStatus.value === "unavailable") {
        return "mdi-alert-circle";
      } else if (selectedBladeStatus.value === "offline") {
        return "mdi-close-circle";
      } else {
        return "mdi-help-circle"; // Default unknown status
      }
    });

    // Methods to update state
    const selectAppliance = (applianceId: string) => {
      applianceStore.selectAppliance(applianceId);
    };
    const selectBlade = (
      bladeId: string,
      bladeIp: string,
      bladePort: number,
      bladeMemoryAvailable: number,
      bladeMemoryAllocated: number,
      bladeStatus: string | undefined
    ) => {
      bladeStore.selectBlade(
        bladeId,
        bladeIp,
        bladePort,
        bladeMemoryAvailable,
        bladeMemoryAllocated,
        bladeStatus
      );
    };

    return {
      appliances,
      blades,
      selectedApplianceId,
      selectedBladeId,
      selectedBladePort,
      selectedBladeIp,
      selectedBladeStatus,
      statusColor,
      statusIcon,
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
