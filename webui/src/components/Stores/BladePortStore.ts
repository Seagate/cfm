// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { PortInformation, DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useBladePortStore = defineStore('bladePort', {
    state: () => ({
        bladePorts: [] as PortInformation[],
    }),

    actions: {
        async fetchBladePorts(applianceId: string, bladeId: string) {
            this.bladePorts = [];
            try {
                // Get all ports
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.bladesGetPorts(
                    applianceId,
                    bladeId
                );

                const portsCount = response.data.memberCount;
                for (let i = 0; i < portsCount; i++) {
                    // Extract the id for each port
                    const uri = response.data.members[i];
                    const portId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get port by id
                    const detailsResponse = await defaultApi.bladesGetPortById(
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
                        } else {
                            detailsResponse.data.linkedPortUri = "NOT_FOUND";
                        }

                        // Combine LinkStatus, linkWidth and linkSpeed if the port is linked up
                        if (detailsResponse.data.linkStatus && detailsResponse.data.linkStatus == "LinkUp") {
                            const linkeStatus = detailsResponse.data.linkStatus;
                            const linkSpeed = detailsResponse.data.currentSpeedGbps;
                            const linkWidth = detailsResponse.data.width;
                            detailsResponse.data.linkStatus = linkeStatus + "/" + linkWidth + "/" + linkSpeed
                        }

                        // Store port in ports list
                        this.bladePorts.push(detailsResponse.data);
                    }
                }
            } catch (error) {
                console.error("Error fetching ports:", error);
            }
        },
    }
})