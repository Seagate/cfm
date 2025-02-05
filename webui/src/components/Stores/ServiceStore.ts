// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia';
import { getDefaultApi } from "../Common/apiService";

export const useServiceStore = defineStore('cfm-service', {
  state: () => ({
    serviceVersion: null as unknown as string,
  }),
  actions: {
    async getServiceVersion() {
      const defaultApi = getDefaultApi();
      try {
        const response =  await defaultApi!.cfmV1Get();
        this.serviceVersion = response.data.version;
      } catch (error) {
        console.error("Error fetching CFM Service version:", error);
      }
    },
  },
})