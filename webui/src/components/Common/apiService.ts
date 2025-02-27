// Copyright (c) 2025 Seagate Technology LLC and/or its Affiliates
import axios from "axios";
import { DefaultApi } from "@/axios/api";

let defaultApi: DefaultApi | null = null;

// Declare a client AxiosInstance for the docker model(docker container) to make the webui to get the correct basepath since vite config doesn't work inside docker
// Determine the current hostname or IP address
const ipAddress = typeof window !== 'undefined' ? window.location.hostname : 'localhost';
// Check if ipAddress is valid
const isValidIpAddress = ipAddress && ipAddress !== '';
// Construct the backend URL with HTTPS protocol and port 8080, or use the environment variable as a fallback
const apiBaseUrl = isValidIpAddress ? `https://${ipAddress}:8080` : import.meta.env.VITE_API_BASE_URL;

const apiClient = axios.create({
  baseURL: apiBaseUrl,
});

// Decide if this is the docker model using the NODE_ENV from the docker
// Use the isDocker flag to force the Web UI to find the correct basepath in apiClient for the docker model
function isDocker() {
  return process.env.NODE_ENV === "docker";
}

// Make the API calls in the web UI are prefixed with /api to match the proxy configuration
// Use API_BASE_PATH to override the BASE_PATH in the generated client code for the development model
const API_BASE_PATH = "/api";

function initializeApi() {
  let axiosInstance = undefined;
  if (isDocker()) {
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