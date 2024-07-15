// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
import { DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

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