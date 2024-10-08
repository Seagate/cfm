// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { Blade, Credentials, DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";
import axios from 'axios';
// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useBladeStore = defineStore('blade', {
    state: () => ({
        blades: [] as Blade[],
        selectedBladeId: null as unknown as string,
        selectedBladeIp: null as unknown as string,
        selectedBladePortNum: null as unknown as number,
        selectedBladeTotalMemoryAvailableMiB: null as unknown as number | undefined,
        selectedBladeTotalMemoryAllocatedMiB: null as unknown as number | undefined,
        selectedBladeStatus: null as unknown as string | undefined,
        addBladeError: null as unknown,
        deleteBladeError: null as unknown,
        resyncBladeError: null as unknown,
    }),
    actions: {
        async fetchBlades(applianceId: string) {
            this.blades = [];
            try {
                // Get all blades from OpenBMC
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const responseOfBlades = await defaultApi.bladesGet(applianceId);
                const bladeCount = responseOfBlades.data.memberCount;

                for (let i = 0; i < bladeCount; i++) {
                    // Extract the id for each blade
                    const uri = responseOfBlades.data.members[i];
                    const bladeId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;

                    // Get blade by id
                    const detailsResponseOfBlade = await defaultApi.bladesGetById(
                        applianceId,
                        bladeId
                    );

                    // Store blade in blades
                    if (detailsResponseOfBlade) {
                        this.blades.push(detailsResponseOfBlade.data);
                    }
                }
            } catch (error) {
                console.error("Error fetching blades:", error);
            }
        },

        async fetchBladeById(applianceId: string, bladeId: string) {
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const detailsResponseOfBlade = await defaultApi.bladesGetById(
                    applianceId,
                    bladeId
                );

                const blade = detailsResponseOfBlade.data;

                // Update the memory for the memory chart because the chart is decided by the blade store not the blade memory store
                this.updateSelectedBladeMemory(blade.totalMemoryAvailableMiB, blade.totalMemoryAllocatedMiB)

                return blade;
            } catch (error) {
                console.error("Error fetching blade by id:", error);
            }
        },

        async resyncBlade(applianceId: string, bladeId: string) {
            this.resyncBladeError = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.bladesResyncById(applianceId, bladeId);

                const resyncedBlade = response.data;
                return resyncedBlade;
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.resyncBladeError = error.message;
                    if (error.response) {
                        this.resyncBladeError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.resyncBladeError = error;
                }
                console.error("Error:", error);
            }
        },


        async addNewBlade(applianceId: string, newBladeCredentials: Credentials) {
            this.addBladeError = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.bladesPost(
                    applianceId,
                    newBladeCredentials
                );
                const newBlade = response.data;
                // Add the new blade to the blades array
                this.blades.push(newBlade);
                return newBlade;
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.addBladeError = error.message;
                    if (error.response) {
                        this.addBladeError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.addBladeError = error;
                }
                console.error("Error:", error);
            }
        },

        async deleteBlade(applianceId: string, bladeId: string) {
            let deletedBlade = "";
            this.deleteBladeError = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.bladesDeleteById(applianceId, bladeId);
                deletedBlade = response.data.id;
                // Remove the deleted blade from the blades array
                if (response) {
                    this.blades = this.blades.filter(
                        (blade) => blade.id !== bladeId
                    );
                }
                return deletedBlade;
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.deleteBladeError = error.message;
                    if (error.response) {
                        this.deleteBladeError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.deleteBladeError = error;
                }
                console.error("Error:", error);
            }
        },


        selectBlade(bladeId: string, selectedBladeIp: string, selectBladePortNum: number, selectedBladeTotalMemoryAvailableMiB: number, selectedBladeTotalMemoryAllocatedMiB: number, status: string | undefined) {
            this.selectedBladeId = bladeId;
            this.selectedBladeIp = selectedBladeIp;
            this.selectedBladePortNum = selectBladePortNum;
            this.selectedBladeTotalMemoryAvailableMiB = selectedBladeTotalMemoryAvailableMiB;
            this.selectedBladeTotalMemoryAllocatedMiB = selectedBladeTotalMemoryAllocatedMiB;
            this.selectedBladeStatus = status;
        },

        updateSelectedBladeMemory(availableMemory: number | undefined, allocatedMemory: number | undefined) {
            this.selectedBladeTotalMemoryAvailableMiB = availableMemory;
            this.selectedBladeTotalMemoryAllocatedMiB = allocatedMemory;
        }
    }
})