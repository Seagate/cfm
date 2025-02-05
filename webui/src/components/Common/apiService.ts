// Copyright (c) 2025 Seagate Technology LLC and/or its Affiliates
import axios from "axios";
import { DefaultApi } from "@/axios/api";

let defaultApi: DefaultApi | null = null;
// Declare a client AxiosInstance for the production model(docker container) to make the webui to get the correct basepath since vite config doesn't work inside docker
const apiBaseUrl = import.meta.env.VITE_API_BASE_URL;
const apiClient = axios.create({
  baseURL: apiBaseUrl,
});

// Decide if this is the production model using the NODE_ENV from the docker
// Use the isProduction flag to force the Web UI to find the correct basepath in apiClient for the production model
function isProduction() {
  return process.env.NODE_ENV === "production";
}

// Make the API calls in the web UI are prefixed with /api to match the proxy configuration
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
const API_BASE_PATH = "/api";

function initializeApi() {
  let axiosInstance = undefined;
  if (isProduction()) {
    axiosInstance = apiClient;
  }
  defaultApi = new DefaultApi(undefined, API_BASE_PATH, axiosInstance);
}

export function getDefaultApi() {
  if (!defaultApi) {
    initializeApi();
  }
  return defaultApi;
}