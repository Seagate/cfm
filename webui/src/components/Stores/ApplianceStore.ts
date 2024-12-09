// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { Appliance, Credentials, DefaultApi, DiscoveredDevice } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";
import axios from 'axios';

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useApplianceStore = defineStore('appliance', {
    state: () => ({
        appliances: [] as Appliance[],
        selectedApplianceId: null as unknown as string,
        addApplianceError: null as unknown,
        deleteApplianceError: null as unknown,
        renameApplianceError: null as unknown,
        applianceIds: [] as { id: string, blades: { id: string, ipAddress: string, status: string | undefined }[] }[],
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
    }),

    actions: {
        async renameAppliance(applianceId: string, newApplianceId: string) {
            this.renameApplianceError = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.appliancesUpdateById(applianceId, newApplianceId);

                // Update the appliances array
                if (response) {
                    this.appliances = this.appliances.filter(
                        (appliance) => appliance.id !== applianceId
                    );
                    this.appliances.push(response.data);
                }

                return response.data
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.renameApplianceError = error.message;

                    if (error.response) {
                        this.renameApplianceError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.renameApplianceError = error;
                }
                console.error("Error:", error);
            }
        },

        async fetchAppliances() {
            this.appliances = [];
            this.applianceIds = [];
            try {
                // Get all appliances from OpenBMC
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const responseOfAppliances = await defaultApi.appliancesGet();
                const applianceCount = responseOfAppliances.data.memberCount;

                for (let i = 0; i < applianceCount; i++) {
                    // Extract the id for each appliance
                    const uri = responseOfAppliances.data.members[i];
                    const applianceId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get appliance by id
                    const detailsResponseOfAppliance = await defaultApi.appliancesGetById(
                        applianceId
                    );

                    // Store appliance in appliances
                    if (detailsResponseOfAppliance) {
                        this.appliances.push(detailsResponseOfAppliance.data);

                        const responseOfBlades = await defaultApi.bladesGet(applianceId);
                        const bladeCount = responseOfBlades.data.memberCount;
                        const associatedBlades = [];

                        for (let i = 0; i < bladeCount; i++) {
                            // Extract the id for each blade
                            const uri = responseOfBlades.data.members[i];
                            const bladeId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                            // Store blade in blades
                            if (bladeId) {
                                const responseOfBlade = await defaultApi.bladesGetById(applianceId, bladeId);
                                const response = { id: responseOfBlade.data.id, ipAddress: responseOfBlade.data.ipAddress, status: responseOfBlade.data.status }
                                associatedBlades.push(response);
                            }
                        }
                        this.applianceIds.push({ id: detailsResponseOfAppliance.data.id, blades: associatedBlades });
                    }
                }
            } catch (error) {
                console.error("Error fetching appliances:", error);
            }
        },

        async discoverBlades() {
            try {
                // Get all the existed blades
                const existedBladeIpAddress: (string | undefined)[] = []
                for (var i = 0; i < this.applianceIds.length; i++) {
                    for (var j = 0; j < this.applianceIds[i].blades.length; j++) {
                        existedBladeIpAddress.push(this.applianceIds[i].blades[j].ipAddress)
                    }
                }

                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                this.discoveredBlades = [];
                const responseOfBlade = await defaultApi.discoverDevices("blade");
                this.discoveredBlades = responseOfBlade.data;

                // Remove the existed blades from the discovered blades
                for (var k = 0; k < this.discoveredBlades.length; k++) {
                    for (var m = 0; m < existedBladeIpAddress.length; m++) {
                        this.discoveredBlades = this.discoveredBlades.filter(
                            (discoveredBlade) => discoveredBlade.address !== existedBladeIpAddress[m]
                        );
                    }
                }

                // Format the device name, remove the .local suffix (e.g. blade device name: granite00.local) from the device name by splitting it with .
                for (var n = 0; n < this.discoveredBlades.length; n++) {
                    this.discoveredBlades[n].name = this.discoveredBlades[n].name!.split('.')[0]
                }

                return this.discoveredBlades
            } catch (error) {
                console.error("Error discovering new devices:", error);
            }
        },

        async addDiscoveredBlades(blade: DiscoveredDevice) {
            const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
            const responseOfApplianceExist = await defaultApi.appliancesGetById(this.defaultApplianceId)

            // If there is no default appliance, add one
            if (!responseOfApplianceExist) {
                this.newApplianceCredentials.customId = this.defaultApplianceId;
                const responseOfAppliance = await defaultApi.appliancesPost(this.newApplianceCredentials);

                // Add the new appliance to the appliances and applianceIds array
                if (responseOfAppliance) {
                    this.appliances.push(responseOfAppliance.data);
                    const newAppliance = { id: responseOfAppliance.data.id, blades: [] }
                    this.applianceIds.push(newAppliance)
                }
            }

            let appliance = this.applianceIds.find(appliance => appliance.id === this.defaultApplianceId);
            // Remove the .local suffix (e.g. blade device name: granite00.local) from the device name by splitting it with . and assign it to the customId
            let deviceName = blade.name!.split('.')[0];
            this.newBladeCredentials.customId = deviceName;
            this.newBladeCredentials.ipAddress = blade.address + "";

            // Add the new discovered blade to the default appliance
            const responseOfBlade = await defaultApi.bladesPost(this.defaultApplianceId, this.newBladeCredentials);

            if (responseOfBlade) {
                const response = { id: responseOfBlade.data.id, ipAddress: responseOfBlade.data.ipAddress, status: responseOfBlade.data.status };
                appliance!.blades.push(response);
            }

            return responseOfBlade.data;
        },

        async addNewAppliance(newAppliance: Credentials) {
            this.addApplianceError = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.appliancesPost(newAppliance);
                const addedAppliance = response.data;
                // Add the new appliance to the appliances array
                this.appliances.push(addedAppliance);
                return addedAppliance;
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.addApplianceError = error.message;
                    if (error.response) {
                        this.addApplianceError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.addApplianceError = error;
                }
                console.error("Error:", error);
            }
        },

        async deleteAppliance(applianceId: string) {
            this.deleteApplianceError = "";
            let deletedAppliance = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.appliancesDeleteById(applianceId);
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
                        this.deleteApplianceError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.deleteApplianceError = error;
                }
                console.error("Error:", error);
            }
        },

        selectAppliance(applianceId: string) {
            this.selectedApplianceId = applianceId;
        },
    }
})