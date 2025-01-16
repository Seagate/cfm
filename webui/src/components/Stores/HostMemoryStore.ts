// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { MemoryRegion, DefaultApi } from "@/axios/api";
// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
import { API_BASE_PATH } from "../Common/Helpers.vue";

export const useHostMemoryStore = defineStore('hostMemory', {
    state: () => ({
        hostMemory: [] as MemoryRegion[],
    }),

    actions: {
        async hostMemoryStore(hostId: string) {
            this.hostMemory = [];
            try {
                // Get all memory
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.hostGetMemory(
                    hostId
                );

                const memoryCount = response.data.memberCount;
                for (let i = 0; i < memoryCount; i++) {
                    // Extract the id for each memory
                    const uri = response.data.members[i];
                    const memoryId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get memory by id
                    const detailsResponse = await defaultApi.hostsGetMemoryById(
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