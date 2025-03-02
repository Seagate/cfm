// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { MemoryDeviceInformation } from "@/axios/api";
import { getDefaultApi } from "../Common/apiService";

export const useHostMemoryDeviceStore = defineStore('hostMemoryDevices', {
    state: () => ({
        hostMemoryDevices: [] as MemoryDeviceInformation[],
    }),

    actions: {
        async hostMemoryDeviceStore(hostId: string) {
            const defaultApi = getDefaultApi();
            this.hostMemoryDevices = [];
            try {
                // Get all memory devices
                const response = await defaultApi!.hostsGetMemoryDevices(
                    hostId
                );

                const memoryDeviceCount = response.data.memberCount;
                for (let i = 0; i < memoryDeviceCount; i++) {
                    // Extract the id for each memory device
                    const uri = response.data.members[i];
                    const memoryDeviceId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get memory device by id
                    const detailsResponse = await defaultApi!.hostsGetMemoryDeviceById(
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