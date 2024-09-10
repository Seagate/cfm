<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-container>
    <h2>CFM Help Center</h2>
    <br />
    <v-row>
      <v-col v-for="(variant, i) in variants" :key="i" cols="6" md="3">
        <v-card
          class="mx-auto custom-size"
          max-width="344"
          :href="variant.url"
          target="_blank"
          rel="noopener"
          variant="tonal"
        >
          <v-card-title>
            <v-icon left>{{ variant.icon }}</v-icon>
            {{ variant.title }}
          </v-card-title>
          <v-card-text>{{ variant.content }}</v-card-text>
          <template v-slot:append>
            <v-icon icon="mdi-open-in-new"></v-icon>
          </template>
        </v-card>
      </v-col>
    </v-row>
    <br />

    <h3>Contact Support</h3>
    <p>
      If you need further assistance, please contact your local Seagate
      representative or our support team.
    </p>
    <v-icon color="#6ebe4a">mdi-email-outline</v-icon>
    Email:
    <a href="mailto:Mag.Software@seagate.com">Mag.Software@seagate.com</a>
  </v-container>
</template>

<style scoped>
.custom-size {
  height: 150px;
}
</style>

<script>
export default {
  data() {
    return {
      variants: [
        {
          id: 0,
          title: "WebUI User Guide",
          content:
            "For instructions on our product, please refer to the CFM User Guide.",
          url: "https://github.com/Seagate/cfm/wiki/WebUI-User-Guide",
          icon: "mdi-book-open-page-variant",
        },
        {
          id: 1,
          title: "CFM GitHub Repository",
          content: "For more details, please visit our GitHub repository.",
          url: "https://github.com/Seagate/cfm",
          icon: "mdi-github",
        },
        {
          id: 2,
          title: "Troubleshooting",
          content:
            "For solutions to common issues, please see our troubleshooting.",
          url: "https://github.com/Seagate/cfm/wiki/Troubleshooting",
          icon: "mdi-tools",
        },
      ],
    };
  },
};
</script>
