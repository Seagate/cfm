<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<script lang="ts">

import axios from "axios";

// Declare a client AxiosInstance for the production model(docker container) to make the webui to get the correct basepath since vite config doesn't work inside docker
const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;
export const apiClient = axios.create({
  baseURL: apiBaseUrl,
});

// Decide if this is the production model using the NODE_ENV from the docker
export function isProduction() {
  return process.env.NODE_ENV === 'production';
}

// Make the API calls in the web UI are prefixed with /api to match the proxy configuration
export const API_BASE_PATH = "/api";

export default {
  methods: {
    getColor(item: string) {
      if (
        item === "Successful" ||
        item === "Success" ||
        item === "Unused" ||
        item === "Enabled"
      )
        return "success";
      if (item === "Reserved") return "warning";
      else return "error";
    },
  },
};
</script>
