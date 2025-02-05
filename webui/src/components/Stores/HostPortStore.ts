// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { PortInformation } from "@/axios/api";
import { getDefaultApi } from "../Common/apiService";

export const useHostPortStore = defineStore('hostPort', {
    state: () => ({
        hostPorts: [] as PortInformation[],
    }),

    actions: {
        async hostPortStore(hostId: string) {
            const defaultApi = getDefaultApi();
            this.hostPorts = [];

            try {
                // Get all ports
                const response = await defaultApi!.hostsGetPorts(
                    hostId
                );

                const portsCount = response.data.memberCount;
                for (let i = 0; i < portsCount; i++) {
                    // Extract the id for each port
                    const uri = response.data.members[i];
                    const portId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get port by id
                    const detailsResponse = await defaultApi!.hostsGetPortById(
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