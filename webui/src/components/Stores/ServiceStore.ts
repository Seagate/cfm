// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia';
import { DefaultApi } from "@/axios/api";
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
import { isProduction, apiClient, API_BASE_PATH } from "../Common/Helpers.vue";

export const useServiceStore = defineStore('cfm-service', {
  state: () => ({
    serviceVersion: null as unknown as string,
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

    async getServiceVersion() {
      await this.initializeApi(); // Ensure API is initialized
      try {
        const response =  await this.defaultApi!.cfmV1Get();
        this.serviceVersion = response.data.version;
      } catch (error) {
        console.error("Error fetching CFM Service version:", error);
      }
    },
  },
})