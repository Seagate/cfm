<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-card
    tile
    flat
    width="100%"
    color="#6ebe4a"
    dark
    class="d-flex align-center justify-center"
  >
    &copy; {{ new Date().getFullYear() }} Seagate | CFM Service Version:
    {{ serviceVersion }} | CFM Web UI Version:
    {{ uiVersion }}
  </v-card>
</template>

<script>
import { DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";
import packageJson from "/package.json";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export default {
  data() {
    return {
      uiVersion: null, // Fetch from the package.json file
      serviceVersion: null, // Fetch from cfm-service through API cfmV1Get
    };
  },

  mounted() {
    this.getServiceVersion();
    this.uiVersion = packageJson.version;
  },

  methods: {
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
};
</script>
