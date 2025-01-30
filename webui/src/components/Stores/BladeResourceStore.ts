// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { MemoryResourceBlock, DefaultApi } from "@/axios/api";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useBladeResourceStore = defineStore('bladeResource', {
    state: () => ({
        memoryResources: [] as MemoryResourceBlock[],
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

        async fetchMemoryResources(applianceId: string, bladeId: string) {
            await this.initializeApi(); // Ensure API is initialized
            this.memoryResources = [];
            try {
                // Get all resources
                const response = await this.defaultApi!.bladesGetResources(
                    applianceId,
                    bladeId
                );

                const resourcesCount = response.data.memberCount;
                for (let i = 0; i < resourcesCount; i++) {
                    // Extract the id for each resources
                    const uri = response.data.members[i];
                    const resourceId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get resource by id
                    const detailsResponse = await this.defaultApi!.bladesGetResourceById(
                        applianceId,
                        bladeId,
                        resourceId
                    );
                    // Store resource in resources list
                    if (detailsResponse) {
                        // change the unit of memory size from MiB to GiB
                        if (detailsResponse.data.capacityMiB) {
                            detailsResponse.data.capacityMiB =
                                detailsResponse.data.capacityMiB / 1024;
                        }
                        this.memoryResources.push(detailsResponse.data);
                    }
                }
            } catch (error) {
                console.error("Error fetching resources:", error);
            }
        },

        async updateMemoryResourcesStatus(applianceId: string, bladeId: string) {
            await this.initializeApi(); // Ensure API is initialized
            try {
                const updatedResource = await this.defaultApi!.bladesGetResourceStatus(
                    applianceId,
                    bladeId,
                );

                if (updatedResource) {
                    // Create a map to quick look up the updatedResource
                    const resourceMap = new Map<string, string>();
                    updatedResource.data.resourceStatuses.forEach((resource) => {
                        resourceMap.set(resource.id, resource.compositionStatus.compositionState);
                    });

                    // Update the status in memoryResources based on the resource map
                    this.memoryResources.forEach(resource => {
                        if (resourceMap.has(resource.id)) {
                            resource.compositionStatus.compositionState = resourceMap.get(resource.id) + ""
                        }
                    });
                }

            } catch (error) {
                console.error("Error updating resources:", error);
            }

        },
    }
})