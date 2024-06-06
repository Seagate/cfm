import { defineStore } from 'pinia'
import { MemoryResourceBlock, DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useBladeResourceStore = defineStore('bladeResource', {
    state: () => ({
        memoryResources: [] as MemoryResourceBlock[],
    }),

    actions: {
        async fetchMemoryResources(applianceId: string, bladeId: string) {
            this.memoryResources = [];
            try {
                // Get all resources
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.bladesGetResources(
                    applianceId,
                    bladeId
                );
                
                const resourcesCount = response.data.memberCount;
                for (let i = 0; i < resourcesCount; i++) {
                    // Extract the id for each resources
                    const uri = response.data.members[i];
                    const resourceId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get resource by id
                    const detailsResponse = await defaultApi.bladesGetResourceById(
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
    }
})