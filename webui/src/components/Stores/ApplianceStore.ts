// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from "pinia";
import {
  Appliance,
  Credentials,
  DefaultApi,
  DiscoveredDevice,
} from "@/axios/api";
import axios from "axios";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useApplianceStore = defineStore("appliance", {
  state: () => ({
    appliances: [] as Appliance[],
    selectedApplianceId: null as unknown as string,
    addApplianceError: null as unknown,
    deleteApplianceError: null as unknown,
    renameApplianceError: null as unknown,
    applianceIds: [] as {
      id: string;
      blades: { id: string; ipAddress: string; status: string | undefined }[];
    }[],
    discoveredBlades: [] as DiscoveredDevice[],

    newBladeCredentials: {
      username: "root",
      password: "0penBmc",
      ipAddress: "127.0.0.1",
      port: 443,
      insecure: true,
      protocol: "https",
      customId: "",
    },

    defaultApplianceId: "CMA_Discovered_Blades",
    newApplianceCredentials: {
      username: "root",
      password: "0penBmc",
      ipAddress: "127.0.0.1",
      port: 8443,
      insecure: true,
      protocol: "https",
      customId: "",
    },
    defaultApi: null as DefaultApi | null,
  }),

  actions: {
    async initializeApi() {
      let axiosInstance = undefined;
      if (isProduction()) {
        axiosInstance = apiClient;
      }
      this.defaultApi = new DefaultApi(undefined, API_BASE_PATH, axiosInstance);
    },

    async renameAppliance(applianceId: string, newApplianceId: string) {
      await this.initializeApi(); // Ensure API is initialized
      this.renameApplianceError = "";
      try {
        const response = await this.defaultApi!.appliancesUpdateById(
          applianceId,
          newApplianceId
        );

        // Update the appliances array
        if (response) {
          this.appliances = this.appliances.filter(
            (appliance) => appliance.id !== applianceId
          );
          this.appliances.push(response.data);
        }

        return response.data;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.renameApplianceError = error.message;

          if (error.response) {
            this.renameApplianceError =
              error.response?.data.status.message +
              " (" +
              error.response?.request.status +
              ")";
          }
        } else {
          this.renameApplianceError = error;
        }
        console.error("Error:", error);
      }
    },

    async fetchAppliances() {
      await this.initializeApi(); // Ensure API is initialized
      this.appliances = [];
      this.applianceIds = [];
      try {
        // Get all appliances from OpenBMC
        const responseOfAppliances = await this.defaultApi!.appliancesGet();
        const applianceCount = responseOfAppliances.data.memberCount;

        for (let i = 0; i < applianceCount; i++) {
          // Extract the id for each appliance
          const uri = responseOfAppliances.data.members[i];
          const applianceId: string = JSON.stringify(uri)
            .split("/")
            .pop()
            ?.slice(0, -2) as string;
          // Get appliance by id
          const detailsResponseOfAppliance = await this.defaultApi!.appliancesGetById(
            applianceId
          );

          // Store appliance in appliances
          if (detailsResponseOfAppliance) {
            this.appliances.push(detailsResponseOfAppliance.data);

            const responseOfBlades = await this.defaultApi!.bladesGet(applianceId);
            const bladeCount = responseOfBlades.data.memberCount;
            const associatedBlades = [];

            for (let i = 0; i < bladeCount; i++) {
              // Extract the id for each blade
              const uri = responseOfBlades.data.members[i];
              const bladeId: string = JSON.stringify(uri)
                .split("/")
                .pop()
                ?.slice(0, -2) as string;
              // Store blade in blades
              if (bladeId) {
                try {
                  const responseOfBlade = await this.defaultApi!.bladesGetById(applianceId, bladeId);

                  const response = {
                    id: responseOfBlade.data.id,
                    ipAddress: responseOfBlade.data.ipAddress,
                    status: responseOfBlade.data.status,
                  };

                  associatedBlades.push(response);
                } catch (bladeError) {
                  console.error(`Error fetching blade ${bladeId}:`, bladeError);

                  // Push the failed blade with an empty status
                  // TODO: Get the status for the failed blade from cfm-service
                  associatedBlades.push({
                    id: bladeId,
                    ipAddress: "",
                    status: "",
                  });
                }
              }
            }
            this.applianceIds.push({
              id: detailsResponseOfAppliance.data.id,
              blades: associatedBlades,
            });
          }
        }
      } catch (error) {
        console.error("Error fetching appliances:", error);
      }
    },

    async discoverBlades() {
      await this.initializeApi(); // Ensure API is initialized
      try {
        // Get all the existed blades
        const existedBladeIpAddress: (string | undefined)[] = [];
        for (let i = 0; i < this.applianceIds.length; i++) {
          for (let j = 0; j < this.applianceIds[i].blades.length; j++) {
            existedBladeIpAddress.push(
              this.applianceIds[i].blades[j].ipAddress
            );
          }
        }

        this.discoveredBlades = [];
        const responseOfBlade = await this.defaultApi!.discoverDevices("blade");
        this.discoveredBlades = responseOfBlade.data;

        // Remove the existed blades from the discovered blades
        for (let k = 0; k < this.discoveredBlades.length; k++) {
          for (let m = 0; m < existedBladeIpAddress.length; m++) {
            this.discoveredBlades = this.discoveredBlades.filter(
              (discoveredBlade) =>
                discoveredBlade.address !== existedBladeIpAddress[m]
            );
          }
        }

        // Format the device name, remove the .local suffix (e.g. blade device name: granite00.local) from the device name by splitting it with .
        for (let n = 0; n < this.discoveredBlades.length; n++) {
          this.discoveredBlades[n].name =
            this.discoveredBlades[n].name!.split(".")[0];
        }

        return this.discoveredBlades;
      } catch (error) {
        console.error("Error discovering new devices:", error);
      }
    },

    async addDiscoveredBlades(blade: DiscoveredDevice) {
      await this.initializeApi(); // Ensure API is initialized
      const responseOfApplianceExist = await this.defaultApi!.appliancesGetById(
        this.defaultApplianceId
      );

      // If there is no default appliance, add one
      if (!responseOfApplianceExist) {
        this.newApplianceCredentials.customId = this.defaultApplianceId;
        const responseOfAppliance = await this.defaultApi!.appliancesPost(
          this.newApplianceCredentials
        );

        // Add the new appliance to the appliances and applianceIds array
        if (responseOfAppliance) {
          this.appliances.push(responseOfAppliance.data);
          const newAppliance = { id: responseOfAppliance.data.id, blades: [] };
          this.applianceIds.push(newAppliance);
        }
      }

      const appliance = this.applianceIds.find(
        (appliance) => appliance.id === this.defaultApplianceId
      );
      // Remove the .local suffix (e.g. blade device name: granite00.local) from the device name by splitting it with . and assign it to the customId
      const deviceName = blade.name!.split(".")[0];
      this.newBladeCredentials.customId = deviceName;
      this.newBladeCredentials.ipAddress = blade.address + "";

      // Add the new discovered blade to the default appliance
      const responseOfBlade = await this.defaultApi!.bladesPost(
        this.defaultApplianceId,
        this.newBladeCredentials
      );

      if (responseOfBlade) {
        const response = {
          id: responseOfBlade.data.id,
          ipAddress: responseOfBlade.data.ipAddress,
          status: responseOfBlade.data.status,
        };
        appliance!.blades.push(response);
      }

      return responseOfBlade.data;
    },

    async addNewAppliance(newAppliance: Credentials) {
      await this.initializeApi(); // Ensure API is initialized
      this.addApplianceError = "";
      try {
        const response = await this.defaultApi!.appliancesPost(newAppliance);
        const addedAppliance = response.data;
        // Add the new appliance to the appliances array
        this.appliances.push(addedAppliance);
        return addedAppliance;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.addApplianceError = error.message;
          if (error.response) {
            this.addApplianceError =
              error.response?.data.status.message +
              " (" +
              error.response?.request.status +
              ")";
          }
        } else {
          this.addApplianceError = error;
        }
        console.error("Error:", error);
      }
    },

    async deleteAppliance(applianceId: string) {
      await this.initializeApi(); // Ensure API is initialized
      this.deleteApplianceError = "";
      let deletedAppliance = "";
      try {
        const response = await this.defaultApi!.appliancesDeleteById(applianceId);
        deletedAppliance = response.data.id;
        // Remove the deleted appliance from the appliances array
        if (response) {
          this.appliances = this.appliances.filter(
            (appliance) => appliance.id !== applianceId
          );
        }
        return deletedAppliance;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.deleteApplianceError = error.message;
          if (error.response) {
            this.deleteApplianceError =
              error.response?.data.status.message +
              " (" +
              error.response?.request.status +
              ")";
          }
        } else {
          this.deleteApplianceError = error;
        }
        console.error("Error:", error);
      }
    },

    selectAppliance(applianceId: string) {
      this.selectedApplianceId = applianceId;
    },
  },
});
