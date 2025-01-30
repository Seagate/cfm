// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { MemoryDeviceInformation, DefaultApi } from "@/axios/api";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useHostMemoryDeviceStore = defineStore('hostMemoryDevices', {
    state: () => ({
        hostMemoryDevices: [] as MemoryDeviceInformation[],
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

        async hostMemoryDeviceStore(hostId: string) {
            await this.initializeApi(); // Ensure API is initialized
            this.hostMemoryDevices = [];
            try {
                // Get all memory devices
                const response = await this.defaultApi!.hostsGetMemoryDevices(
                    hostId
                );

                const memoryDeviceCount = response.data.memberCount;
                for (let i = 0; i < memoryDeviceCount; i++) {
                    // Extract the id for each memory device
                    const uri = response.data.members[i];
                    const memoryDeviceId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get memory device by id
                    const detailsResponse = await this.defaultApi!.hostsGetMemoryDeviceById(
                        hostId,
                        memoryDeviceId
                    );
                    // Store memory in memory list
                    if (detailsResponse) {
                        this.hostMemoryDevices.push(detailsResponse.data);
                    }
                }
            } catch (error) {
                console.error("Error fetching memory device:", error);
            }
        },
    }
})