// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { PortInformation, DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useHostPortStore = defineStore('hostPort', {
    state: () => ({
        hostPorts: [] as PortInformation[],
    }),

    actions: {
        async hostPortStore(hostId: string) {
            this.hostPorts = [];
            try {
                // Get all ports
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.hostsGetPorts(
                    hostId
                );

                const portsCount = response.data.memberCount;
                for (let i = 0; i < portsCount; i++) {
                    // Extract the id for each port
                    const uri = response.data.members[i];
                    const portId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get port by id
                    const detailsResponse = await defaultApi.hostsGetPortById(
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