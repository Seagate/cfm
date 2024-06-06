<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="portsDetailsList"
      :sort-by="[{ key: 'portId', order: 'asc' }]"
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
      hostIdForPorts: this.hostId,

      headers: [
        {
          title: "PortId",
          align: "start",
          key: "id",
        },
        { title: "GCxlId", key: "gCxlId" },
        { title: "LinkedApplianceBlade", key: "linkedPortUri" },
        { title: "LinkStatus", key: "linkStatus" },
      ],
      portsDetailsList: [],
      search: "",
      portsCount: 0,
    };
  },

  mounted() {
    this.loadData();
  },

  methods: {
    async loadData() {
      try {
        await this.fetchPorts();
      } catch (error) {
        console.error("Error loading data:", error);
      }
    },

    async fetchPorts() {
      try {
        // Get all ports
        const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
        const response = await defaultApi.hostsGetPorts(this.hostIdForPorts);
        this.portsCount = response.data.memberCount;
        for (let i = 0; i < this.portsCount; i++) {
          // Extract the id for each ports
          const uri = response.data.members[i];
          const portId = JSON.stringify(uri).split("/").pop().slice(0, -2);
          // Get port by id
          const detailsResponse = await defaultApi.hostsGetPortById(
            this.hostIdForPorts,
            portId
          );
          // Store port in ports list
          if (detailsResponse) {
            if (detailsResponse.data.linkedPortUri != "NOT_FOUND") {
              // Fetch linked appliance and blade id from LinkedPortUri
              const uri = JSON.stringify(
                detailsResponse.data.linkedPortUri
              ).split("/");
              const applianceId = uri[4];
              const hostId = uri[6];
              detailsResponse.data.linkedPortUri = applianceId + "/" + hostId;
            }
            this.portsDetailsList.push(detailsResponse.data);
          }
        }
      } catch (error) {
        console.error("Error fetching ports:", error);
      }
    },
  },
};
</script>