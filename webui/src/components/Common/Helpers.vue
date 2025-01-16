<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<script lang="ts">
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