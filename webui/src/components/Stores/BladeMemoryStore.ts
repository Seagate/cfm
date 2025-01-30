// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { MemoryRegion, DefaultApi, ComposeMemoryRequest, AssignMemoryRequest } from "@/axios/api";
import axios from 'axios';
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useBladeMemoryStore = defineStore('bladeMemory', {
    state: () => ({
        bladeMemory: [] as MemoryRegion[],
        portIds: null as unknown as string[],
        assignOrUnassignMemoryError: null as unknown,
        freeMemoryError: null as unknown,
        composeMemoryError: null as unknown,
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

        async fetchBladeMemory(applianceId: string, bladeId: string) {
            await this.initializeApi(); // Ensure API is initialized
            this.bladeMemory = [];
            try {
                // Get all memory
                const response = await this.defaultApi!.bladesGetMemory(
                    applianceId,
                    bladeId
                );

                const memoryCount = response.data.memberCount;
                for (let i = 0; i < memoryCount; i++) {
                    // Extract the id for each memory
                    const uri = response.data.members[i];
                    const memoryId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get memory by id
                    const detailsResponse = await this.defaultApi!.bladesGetMemoryById(
                        applianceId,
                        bladeId,
                        memoryId
                    );
                    // Store memory in memory list
                    if (detailsResponse) {
                        // change the unit of memory size from MiB to GiB
                        detailsResponse.data.sizeMiB = detailsResponse.data.sizeMiB / 1024;
                        this.bladeMemory.push(detailsResponse.data);
                    }
                }
            } catch (error) {
                console.error("Error fetching memory:", error);
            }
        },

        async composeMemory(applianceId: string, bladeId: string, newMemoryCredentials: ComposeMemoryRequest) {
            await this.initializeApi(); // Ensure API is initialized
            // Reset the error before each compose memory operation
            this.composeMemoryError = null;
            try {
                const response = await this.defaultApi!.bladesComposeMemory(
                    applianceId, bladeId, newMemoryCredentials
                );
                const newMemory = response.data;
                newMemory.sizeMiB = newMemory.sizeMiB / 1024;
                // Store the new memory chunk to blade Memory array
                this.bladeMemory.push(newMemory);
                return newMemory;
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.composeMemoryError = error.message;
                    if (error.response) {
                        this.composeMemoryError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.composeMemoryError = error;
                }
                console.error("Error:", error);
            }
        },

        async assignOrUnassign(applianceId: string, bladeId: string, memoryId: string, assignMemoryRequest: AssignMemoryRequest) {
            await this.initializeApi(); // Ensure API is initialized
            try {
                const response = await this.defaultApi!.bladesAssignMemoryById(
                    applianceId,
                    bladeId,
                    memoryId,
                    assignMemoryRequest
                );
                return response;
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.assignOrUnassignMemoryError = error.message;
                    if (error.response) {
                        this.assignOrUnassignMemoryError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.assignOrUnassignMemoryError = error;
                }
                console.error("Error assign or unassign memory:", error);
            }
        },

        async freeMemory(applianceId: string, bladeId: string, memoryId: string) {
            await this.initializeApi(); // Ensure API is initialized
            try {
                const response = await this.defaultApi!.bladesFreeMemoryById(
                    applianceId,
                    bladeId,
                    memoryId
                );
                // Remove the memory chunk from the bladeMemory array
                if (response) {
                    this.bladeMemory = this.bladeMemory.filter(
                        (memory) => memory.id !== memoryId
                    );
                }
                return response;
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    this.freeMemoryError = error.message;
                    if (error.response) {
                        this.freeMemoryError = error.response?.data.status.message + " (" + error.response?.request.status + ")";
                    }
                }
                else {
                    this.freeMemoryError = error;
                }
                console.error("Error free memory:", error);
            }
        }
    }
})