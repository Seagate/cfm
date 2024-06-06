<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="memoryDeviceDetailsList"
      :sort-by="[{ key: 'memoryDeviceId', order: 'asc' }]"
      :search="search"
      :height="180"
    >
    </v-data-table>
  </v-container>
</template>

<script>
import { DefaultApi } from "@/axios/api";
import { BASE_PATH } from "@/axios/base";

// Use API_BASE_PATH to overwrite the BASE_PATH in the generated client code
const API_BASE_PATH = process.env.BASE_PATH || BASE_PATH;

export default {
  props: {
    hostId: String,
  },

  data() {
    return {
      hostIdForMemory: this.hostId,

      headers: [
        {
          title: "MemoryDeviceId",
          align: "start",
          key: "id",
        },
        { title: "DeviceType", key: "deviceType" },
        { title: "LinkStatus", key: "linkStatus" },
      ],
      memoryDeviceDetailsList: [],
      search: "",
      memoryDevicesCount: 0,
    };
  },
  mounted() {
    this.loadData();
  },
  methods: {
    async loadData() {
      try {
        await this.fetchMemoryDevices();
      } catch (error) {
        console.error("Error loading data:", error);
      }
    },
    async fetchMemoryDevices() {
      try {
        // Get all memory devices
        const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
        const response = await defaultApi.hostsGetMemoryDevices(
          this.hostIdForMemory
        );
        this.memoryDevicesCount = response.data.memberCount;
        for (let i = 0; i < this.memoryDevicesCount; i++) {
          // Extract the id for each memoryDevice
          const uri = response.data.members[i];
          const memoryDeviceId = JSON.stringify(uri)
            .split("/")
            .pop()
            .slice(0, -2);
          // Get memoryDevice by id
          const detailsResponse = await defaultApi.hostsGetMemoryDeviceById(
            this.hostIdForMemory,
            memoryDeviceId
          );
          // Pick up memoryDevice id
          // TODO: Here is a bug in cfm-service, should return id (39-00) not the whole path (/cfm/v1/hosts/host-0083/memorydevices/39-00)
          detailsResponse.data.id = memoryDeviceId;
          // Store memoryDevice in memoryDevice list
          if (detailsResponse) {
            this.memoryDeviceDetailsList.push(detailsResponse.data);
          }
        }
      } catch (error) {
        console.error("Error fetching memoryDevice:", error);
      }
    },
  },
};
</script>