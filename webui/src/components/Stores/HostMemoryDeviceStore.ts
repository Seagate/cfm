import { defineStore } from 'pinia'
import { MemoryDeviceInformation, DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useHostMemoryDeviceStore = defineStore('hostMemoryDevices', {
    state: () => ({
        hostMemoryDevices: [] as MemoryDeviceInformation[],
    }),

    actions: {
        async hostMemoryDeviceStore(hostId: string) {
            this.hostMemoryDevices = [];
            try {
                // Get all memory devices
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.hostsGetMemoryDevices(
                    hostId
                );

                const memoryDeviceCount = response.data.memberCount;
                for (let i = 0; i < memoryDeviceCount; i++) {
                    // Extract the id for each memory device
                    const uri = response.data.members[i];
                    const memoryDeviceId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get memory device by id
                    const detailsResponse = await defaultApi.hostsGetMemoryDeviceById(
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