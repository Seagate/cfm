// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { Host, Credentials, DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export const useHostStore = defineStore('host', {
    state: () => ({
        hosts: [] as Host[],
        selectedHostId: null as unknown as string,
        selectedHostIp: null as unknown as string,
        selectedHostPortNum: null as unknown as number,
        selectedHostLocalMemory: null as unknown as number | undefined,
        addHostError: null as unknown,
        deleteHostError: null as unknown,
        resyncHostError: null as unknown,
        hostIds: [] as string[],
    }),

    actions: {
        async fetchHosts() {
            this.hosts = [];
            this.hostIds = [];
            try {
                // Get all hosts from OpenBMC
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const responseOfHosts = await defaultApi.hostsGet();
                const hostCount = responseOfHosts.data.memberCount;

                for (let i = 0; i < hostCount; i++) {
                    // Extract the id for each host
                    const uri = responseOfHosts.data.members[i];
                    const hostId: string = JSON.stringify(uri).split("/").pop()?.slice(0, -2) as string;
                    // Get host by id
                    const detailsResponseOfHost = await defaultApi.hostsGetById(
                        hostId
                    );

                    // Store host in hosts
                    if (detailsResponseOfHost) {
                        this.hosts.push(detailsResponseOfHost.data);
                        this.hostIds.push(detailsResponseOfHost.data.id)
                    }
                }
            } catch (error) {
                console.error("Error fetching hosts:", error);
            }
        },

        async addNewHost(newHost: Credentials) {
            this.addHostError = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.hostsPost(newHost);
                const addedHost = response.data;
                // Add the new host to the hosts array
                this.hosts.push(addedHost);
                return addedHost;
            } catch (error) {
                this.addHostError = error;
                console.error("Error:", error);
            }
        },

        async deleteHost(hostId: string) {
            this.deleteHostError = "";
            let deletedHost = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.hostsDeleteById(hostId);
                deletedHost = response.data.id;
                // Remove the deleted host from the hosts array
                if (response) {
                    this.hosts = this.hosts.filter(
                        (host) => host.id !== hostId
                    );
                }
                return deletedHost;
            } catch (error) {
                this.deleteHostError = error;
                console.error("Error:", error);
            }
        },

        async resyncHost(hostId: string) {
            this.resyncHostError = "";
            try {
                const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
                const response = await defaultApi.hostsResyncById(hostId);

                const resyncedHost = response.data;
                return resyncedHost;
            } catch (error) {
                this.resyncHostError = error;
                console.error("Error:", error);
            }
        },

        selectHost(selectedHostId: string, selectedHostIp: string, selectedHostPortNum: number, selectedHostLocalMemory: number | undefined) {
            this.selectedHostId = selectedHostId;
            this.selectedHostIp = selectedHostIp;
            this.selectedHostPortNum = selectedHostPortNum;
            this.selectedHostLocalMemory = selectedHostLocalMemory;
        },
    }
})