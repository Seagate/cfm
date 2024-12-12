// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import { defineStore } from 'pinia'
//import { DefaultApi } from "@/axios/api";
//import { BASE_PATH } from "@/axios/base";
import { CustomApi } from "@/common/CustomApi"
import { Configuration } from '@/axios';

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
//const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH as ConfigurationParameters | undefined;

export const useServiceStore = defineStore('cfm-service', {
    state: () => ({
        serviceVersion: null as unknown as string,
    }),
    actions: {
        async getServiceVersion() {
            try {
              const config = new Configuration;
              config.basePath = "https://localhost:8080"
              const defaultApi = new CustomApi(config);
              const response = await defaultApi.cfmV1Get();
              this.serviceVersion = response.data.version;
            } catch (error) {
              console.error("Error fetching CFM Service version:", error);
            }
          },
    },
})