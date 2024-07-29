// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { Appliance, Credentials, DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useApplianceStore = defineStore('appliance', {
    state: () => ({
        appliances: [] as Appliance[],
        selectedApplianceId: null as unknown as string,
        addApplianceError: null as unknown,
        deleteApplianceError: null as unknown,
        applianceIds: [] as { id: string, bladeIds: string[] }[],
    }),

    actions: {
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
                        const bladeIds = [];

                        for (let i = 0; i < bladeCount; i++) {
                            // Extract the id for each blade
                            const uri = responseOfBlades.data.members[i];
                            const bladeId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                            // Store blade in blades
                            if (bladeId) {
                                bladeIds.push(bladeId);
                            }
                        }
                        this.applianceIds.push({ id: detailsResponseOfAppliance.data.id, bladeIds });
                    }
                }
            } catch (error) {
                console.error("Error fetching appliances:", error);
            }
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
                this.addApplianceError = error;
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
                this.deleteApplianceError = error;
                console.error("Error:", error);
            }
        },

        selectAppliance(applianceId: string) {
            this.selectedApplianceId = applianceId;
        },
    }
})