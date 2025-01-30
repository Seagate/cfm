// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { MemoryRegion, DefaultApi } from "@/axios/api";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useHostMemoryStore = defineStore('hostMemory', {
    state: () => ({
        hostMemory: [] as MemoryRegion[],
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

        async hostMemoryStore(hostId: string) {
            await this.initializeApi(); // Ensure API is initialized
            this.hostMemory = [];

            try {
                // Get all memory
                const response = await this.defaultApi!.hostGetMemory(
                    hostId
                );

                const memoryCount = response.data.memberCount;
                for (let i = 0; i < memoryCount; i++) {
                    // Extract the id for each memory
                    const uri = response.data.members[i];
                    const memoryId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get memory by id
                    const detailsResponse = await this.defaultApi!.hostsGetMemoryById(
                        hostId,
                        memoryId
                    );
                    // Store memory in memory list
                    if (detailsResponse) {
                        // change the unit of memory size from MiB to GiB
                        detailsResponse.data.sizeMiB = +(detailsResponse.data.sizeMiB / 1024).toFixed(0);
                        this.hostMemory.push(detailsResponse.data);
                    }
                }
            } catch (error) {
                console.error("Error fetching memory:", error);
            }
        },
    }
})