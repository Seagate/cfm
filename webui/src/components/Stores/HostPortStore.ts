// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { PortInformation, DefaultApi } from "@/axios/api";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useHostPortStore = defineStore('hostPort', {
    state: () => ({
        hostPorts: [] as PortInformation[],
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

        async hostPortStore(hostId: string) {
            await this.initializeApi(); // Ensure API is initialized
            this.hostPorts = [];

            try {
                // Get all ports
                const response = await this.defaultApi!.hostsGetPorts(
                    hostId
                );

                const portsCount = response.data.memberCount;
                for (let i = 0; i < portsCount; i++) {
                    // Extract the id for each port
                    const uri = response.data.members[i];
                    const portId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get port by id
                    const detailsResponse = await this.defaultApi!.hostsGetPortById(
                        hostId,
                        portId
                    );
                    // Store port in ports list
                    if (detailsResponse) {
                        if (detailsResponse.data.linkedPortUri != "NOT_FOUND") {
                            // Fetch linked appliance and blade id from LinkedPortUri
                            const uri = JSON.stringify(
                                detailsResponse.data.linkedPortUri
                            ).split("/");
                            const applianceId = uri[4];
                            const bladeId = uri[6];
                            detailsResponse.data.linkedPortUri = applianceId + "/" + bladeId;
                        }
                        this.hostPorts.push(detailsResponse.data);
                    }
                }
            } catch (error) {
                console.error("Error fetching ports:", error);
            }
        },
    }
})