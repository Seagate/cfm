// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { PortInformation, DefaultApi } from "@/axios/api";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useBladePortStore = defineStore('bladePort', {
    state: () => ({
        bladePorts: [] as PortInformation[],
        // Use bladeIds to store the relationship between blade and host, the relationship is used to determine the dataEdges in the dashboard
        bladeIds: [] as { id: string, connectedHostIds: string[] }[],
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

        async fetchBladePorts(applianceId: string, bladeId: string) {
            await this.initializeApi(); // Ensure API is initialized
            this.bladePorts = [];
            try {
                // Get all ports
                const response = await this.defaultApi!.bladesGetPorts(
                    applianceId,
                    bladeId
                );

                const hostIds = [];
                const portsCount = response.data.memberCount;
                for (let i = 0; i < portsCount; i++) {
                    // Extract the id for each port
                    const uri = response.data.members[i];
                    const portId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get port by id
                    const detailsResponse = await this.defaultApi!.bladesGetPortById(
                        applianceId,
                        bladeId,
                        portId
                    );

                    if (detailsResponse) {
                        // Store linked host port in ports list
                        if (detailsResponse.data.linkedPortUri) {
                            // Fetch linked host id from LinkedPortUri
                            const linkedPortUri = detailsResponse.data.linkedPortUri;
                            const hostId = JSON.stringify(linkedPortUri).split("/")[4];
                            const hostPort: string = JSON.stringify(linkedPortUri).split("/").pop()?.slice(0, -1) as string;
                            detailsResponse.data.linkedPortUri = hostId + "/" + hostPort;
                            hostIds.push(hostId);
                        } else {
                            detailsResponse.data.linkedPortUri = "NOT_FOUND";
                        }

                        // Combine LinkStatus, linkWidth and linkSpeed if the port is linked up
                        if (detailsResponse.data.linkStatus && detailsResponse.data.linkStatus == "Link Up") {
                            //Remove the space in LinkStatus
                            const linkeStatus = detailsResponse.data.linkStatus.replace(/\s+/g, '');
                            const linkSpeed = detailsResponse.data.currentSpeedGbps;
                            const linkWidth = detailsResponse.data.width;
                            detailsResponse.data.linkStatus = linkeStatus + "/" + linkWidth + "/" + linkSpeed
                        }

                        // Store port in ports list
                        this.bladePorts.push(detailsResponse.data);
                    }
                }
                this.bladeIds.push({ id: bladeId, connectedHostIds: hostIds });
            } catch (error) {
                console.error("Error fetching ports:", error);
            }
        },
    }
})