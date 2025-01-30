// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from "pinia";
import { Host, Credentials, DefaultApi, DiscoveredDevice } from "@/axios/api";
import axios from "axios";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useHostStore = defineStore("host", {
  state: () => ({
    hosts: [] as Host[],
    selectedHostId: null as unknown as string,
    selectedHostIp: null as unknown as string,
    selectedHostPortNum: null as unknown as number,
    selectedHostLocalMemory: null as unknown as number | undefined,
    selectedHostStatus: null as unknown as string | undefined,
    addHostError: null as unknown,
    deleteHostError: null as unknown,
    resyncHostError: null as unknown,
    renameHostError: null as unknown,
    hostIds: [] as {
      id: string;
      ipAddress: string;
      status: string | undefined;
    }[],
    discoveredHosts: [] as DiscoveredDevice[],

    newHostCredentials: {
      username: "admin",
      password: "admin12345",
      ipAddress: "127.0.0.1",
      port: 8082,
      insecure: true,
      protocol: "http",
      customId: "",
    },
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

    async fetchHosts() {
      this.hosts = [];
      this.hostIds = [];
      
      try {
        await this.initializeApi(); // Ensure API is initialized
        // Get all hosts from OpenBMC
        const responseOfHosts = await this.defaultApi!.hostsGet();
        const hostCount = responseOfHosts.data.memberCount;

        for (let i = 0; i < hostCount; i++) {
          // Extract the id for each host
          const uri = responseOfHosts.data.members[i];
          const hostId: string = JSON.stringify(uri)
            .split("/")
            .pop()
            ?.slice(0, -2) as string;
          // Get host by id
          if (hostId) {
            try {
              const detailsResponseOfHost = await this.defaultApi!.hostsGetById(hostId);

              // Store host in hosts
              if (detailsResponseOfHost) {
                this.hosts.push(detailsResponseOfHost.data);
                const host = {
                  id: detailsResponseOfHost.data.id,
                  ipAddress: detailsResponseOfHost.data.ipAddress,
                  status: detailsResponseOfHost.data.status,
                };
                this.hostIds.push(host);
              }
            }
            catch (hostError) {
              console.error(`Error fetching host ${hostId}:`, hostError);
              // Push the failed host with an empty status
              // TODO: Get the status for the failed host from cfm-service
              this.hostIds.push({
                id: hostId,
                ipAddress: "",
                status: "",
              });
            }
          }
        }
      } catch (error) {
        console.error("Error fetching hosts:", error);
      }
    },

    async fetchHostById(hostId: string) {
      await this.initializeApi(); // Ensure API is initialized
      try {
        const detailsResponseOfHost = await this.defaultApi!.hostsGetById(hostId);

        const host = detailsResponseOfHost.data;
        this.updateSelectHostStatus(host.status);

        // Update hosts in case this host changes
        if (host) {
          this.hosts = this.hosts.map((h) =>
            h.id === hostId ? detailsResponseOfHost.data : h
          );
        }

        return host;
      } catch (error) {
        console.error("Error fetching host by id:", error);
      }
    },

    async discoverHosts() {
      await this.initializeApi(); // Ensure API is initialized
      try {
        // Get all the existed hosts
        const existedHostIpAddress: (string | undefined)[] = [];
        for (let i = 0; i < this.hostIds.length; i++) {
          existedHostIpAddress.push(this.hostIds[i].ipAddress);
        }

        this.discoveredHosts = [];
        const responseOfHost = await this.defaultApi!.discoverDevices("cxl-host");
        this.discoveredHosts = responseOfHost.data;

        // Remove the existed hosts from the discovered hosts
        for (let k = 0; k < this.discoveredHosts.length; k++) {
          for (let m = 0; m < existedHostIpAddress.length; m++) {
            this.discoveredHosts = this.discoveredHosts.filter(
              (discoveredHost) =>
                discoveredHost.address !== existedHostIpAddress[m]
            );
          }
        }

        // Format the device name, remove the .local suffix (e.g. host device name: host00.local) from the device name by splitting it with .
        for (let n = 0; n < this.discoveredHosts.length; n++) {
          this.discoveredHosts[n].name =
            this.discoveredHosts[n].name!.split(".")[0];
        }

        return this.discoveredHosts;
      } catch (error) {
        console.error("Error discovering new devices:", error);
      }
    },

    async addDiscoveredHosts(host: DiscoveredDevice) {
      await this.initializeApi(); // Ensure API is initialized

      // Remove the .local suffix (e.g. host device name: host00.local) from the device name by splitting it with . and assign it to the customId
      const deviceName = host.name!.split(".")[0];
      this.newHostCredentials.customId = deviceName;
      this.newHostCredentials.ipAddress = host.address + "";

      // Add the new didcovered host
      const responseOfHost = await this.defaultApi!.hostsPost(
        this.newHostCredentials
      );

      // Update the hostIds and hosts
      if (responseOfHost) {
        const response = {
          id: responseOfHost.data.id,
          ipAddress: responseOfHost.data.ipAddress,
          status: responseOfHost.data.status,
        };
        this.hosts.push(responseOfHost.data);
        this.hostIds.push(response);
      }
      return responseOfHost.data;
    },

    async addNewHost(newHost: Credentials) {
      await this.initializeApi(); // Ensure API is initialized
      this.addHostError = "";
      try {
        const response = await this.defaultApi!.hostsPost(newHost);
        const addedHost = response.data;
        // Add the new host to the hosts array
        this.hosts.push(addedHost);
        return addedHost;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.addHostError = error.message;
          if (error.response) {
            this.addHostError =
              error.response?.data.status.message +
              " (" +
              error.response?.request.status +
              ")";
          }
        } else {
          this.addHostError = error;
        }
        console.error("Error:", error);
      }
    },

    async deleteHost(hostId: string) {
      await this.initializeApi(); // Ensure API is initialized
      this.deleteHostError = "";
      let deletedHost = "";

      try {
        const response = await this.defaultApi!.hostsDeleteById(hostId);
        deletedHost = response.data.id;
        // Remove the deleted host from the hosts array
        if (response) {
          this.hosts = this.hosts.filter((host) => host.id !== hostId);
        }
        return deletedHost;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.deleteHostError = error.message;
          if (error.response) {
            this.deleteHostError =
              error.response?.data.status.message +
              " (" +
              error.response?.request.status +
              ")";
          }
        } else {
          this.deleteHostError = error;
        }
        console.error("Error:", error);
      }
    },

    async renameHost(hostId: string, newHostId: string) {
      await this.initializeApi(); // Ensure API is initialized
      this.renameHostError = "";
      try {
        const response = await this.defaultApi!.hostsUpdateById(hostId, newHostId);

        // Update the hosts array
        if (response) {
          this.hosts = this.hosts.filter((host) => host.id !== hostId);
          this.hosts.push(response.data);
        }
        return response.data;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.renameHostError = error.message;
          if (error.response) {
            this.renameHostError =
              error.response?.data.status.message +
              " (" +
              error.response?.request.status +
              ")";
          }
        } else {
          this.renameHostError = error;
        }
        console.error("Error:", error);
      }
    },

    async resyncHost(hostId: string) {
      await this.initializeApi(); // Ensure API is initialized
      this.resyncHostError = "";

      try {
        const response = await this.defaultApi!.hostsResyncById(hostId);

        const resyncedHost = response.data;
        return resyncedHost;
      } catch (error) {
        if (axios.isAxiosError(error)) {
          this.resyncHostError = error.message;
          if (error.response) {
            this.resyncHostError =
              error.response?.data.status.message +
              " (" +
              error.response?.request.status +
              ")";
          }
        } else {
          this.resyncHostError = error;
        }
        console.error("Error:", error);
      }
    },

    selectHost(
      selectedHostId: string,
      selectedHostIp: string,
      selectedHostPortNum: number,
      selectedHostLocalMemory: number | undefined,
      status: string | undefined
    ) {
      this.selectedHostId = selectedHostId;
      this.selectedHostIp = selectedHostIp;
      this.selectedHostPortNum = selectedHostPortNum;
      this.selectedHostLocalMemory = selectedHostLocalMemory;
      this.selectedHostStatus = status;
    },
    updateSelectHostStatus(status: string | undefined) {
      this.selectedHostStatus = status;
    },
  },
});
