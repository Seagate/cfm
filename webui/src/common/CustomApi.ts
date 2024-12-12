// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
import axios, { AxiosInstance } from 'axios';
import { DefaultApi } from "@/axios/api";
import type { Configuration } from '@/axios/configuration';
import https from 'https';
import fs from 'fs';

// Load the certificate
const cert = fs.readFileSync('/usr/local/share/ca-certificates/github_com_seagate_cfm-self-signed.crt');
// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH;

// Create a custom Axios instance with HTTPS configuration
const axiosInstance: AxiosInstance = axios.create({
    baseURL: API_BASE_PATH,
    httpsAgent: new https.Agent({
      ca: cert, 
      rejectUnauthorized: false,
    }),
    headers: {
      'Content-Type': 'application/json',
    }
  });
  
  // Extend the DefaultApi class to use the custom Axios instance
  export class CustomApi extends DefaultApi {
    constructor(configuration?: Configuration, basePath?: string) {
      super(configuration, basePath, axiosInstance);
    }
  }
