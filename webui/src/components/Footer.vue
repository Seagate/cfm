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
    &copy; 2024 Seagate Technology LLC | CFM Service Version:
    <span :style="{ color: serviceVersion ? 'inherit' : 'red', margin: '8px' }">
      {{ serviceVersion || "Not Found" }}
    </span>
    | CFM Web UI Version:
    <span :style="{ margin: '8px' }"> {{ uiVersion }} </span>
  </v-card>
</template>

<script>
import { computed } from "vue";
import packageJson from "/package.json";
import { useServiceStore } from "./Stores/ServiceStore";

export default {
  setup() {
    const serviceStore = useServiceStore();
    const serviceVersion = computed(() => serviceStore.serviceVersion);

    const uiVersion = packageJson.version;

    return {
      serviceVersion,
      uiVersion,
    };
  },
};
</script>