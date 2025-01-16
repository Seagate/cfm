// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia';
import { DefaultApi } from "@/axios/api";
// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
import { API_BASE_PATH } from "../Common/Helpers.vue";

export const useServiceStore = defineStore('cfm-service', {
    state: () => ({
        serviceVersion: null as unknown as string,
    }),
    actions: {
        async getServiceVersion() {
            try {
              const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
              const response = await defaultApi.cfmV1Get();
              this.serviceVersion = response.data.version;
            } catch (error) {
              console.error("Error fetching CFM Service version:", error);
            }
          },
    },
})