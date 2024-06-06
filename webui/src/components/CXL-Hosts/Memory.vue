<template>
  <v-container>
    <v-data-table
      :headers="headers"
      :items="memoryDetailsList"
      :sort-by="[{ key: 'memoryId', order: 'asc' }]"
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
          title: "MemoryId",
          align: "start",
          key: "id",
        },
        { title: "Type", key: "type" },
        { title: "SizeGiB", key: "sizeMiB" },
      ],
      memoryDetailsList: [],
      search: "",
      memoryCount: 0,
    };
  },
  mounted() {
    this.loadData();
  },
  methods: {
    async loadData() {
      try {
        await this.fetchMemory();
      } catch (error) {
        console.error("Error loading data:", error);
      }
    },
    async fetchMemory() {
      try {
        // Get all memory
        const defaultApi = new DefaultApi(undefined, API_BASE_PATH);
        const response = await defaultApi.hostGetMemory(this.hostIdForMemory);
        this.memoryCount = response.data.memberCount;
        for (let i = 0; i < this.memoryCount; i++) {
          // Extract the id for each memory
          const uri = response.data.members[i];
          const memoryId = JSON.stringify(uri).split("/").pop().slice(0, -2);
          // Get memory by id
          const detailsResponse = await defaultApi.hostsGetMemoryById(
            this.hostIdForMemory,
            memoryId
          );
          // Pick up memory id
          // TODO: Here is a bug in cfm-service, should return id (node0) not the whole path (/cfm/v1/hosts/host-0083/memory/node0)
          detailsResponse.data.id = memoryId;
          // Store memory in memory list
          if (detailsResponse) {
            // change the unit of memory size from MiB to GiB
            detailsResponse.data.sizeMiB = (
              detailsResponse.data.sizeMiB / 1024
            ).toFixed(0);
            this.memoryDetailsList.push(detailsResponse.data);
          }
        }
      } catch (error) {
        console.error("Error fetching memory:", error);
      }
    },
  },
};
</script>