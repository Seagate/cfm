<template>
  <v-toolbar app color="#6ebe4a" density="compact">
    <v-img
      class="mx-auto"
      max-width="100"
      src="@/assets/seagate_1c_blk_horizontal.png"
    >
    </v-img>

    <v-toolbar-title>Composer & Fabric Manager Console</v-toolbar-title>

    <v-toolbar-items>
      <v-btn>{{ new Date().toLocaleDateString() }}</v-btn>

      <!-- 
  // Hide the under construction Help and Settings page
  
      <v-divider vertical></v-divider>

      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn icon="mdi-dots-vertical" v-bind="props" variant="text"></v-btn>
        </template>

        <v-list>
          <v-list-item>
            <v-list-item-title
              ><v-btn
                variant="text"
                prepend-icon="mdi-help-circle-outline"
                @click="help = true"
                >Help</v-btn
              >
            </v-list-item-title>

            <v-list-item-title>
              <v-btn
                variant="text"
                prepend-icon="mdi-cog-outline"
                @click="settings = true"
                >Settings</v-btn
              ></v-list-item-title
            >
          </v-list-item>
        </v-list>
      </v-menu>
       -->
    </v-toolbar-items>
  </v-toolbar>

  <v-dialog v-model="help" max-width="600px">
    <v-container>
      <v-card>
        <v-card-item>
          <div class="text-h5 mb-1">
            Help page is under construction, thanks for your patience 😘
          </div>
        </v-card-item>
        <v-card-actions>
          <v-btn color="#6ebe4a" variant="text" @click="help = false">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-container>
  </v-dialog>

  <v-dialog v-model="settings" max-width="600px">
    <v-container>
      <v-card>
        <v-card-item>
          <div class="text-h5 mb-1">
            Settings page is under construction, thanks for your patience 😘
          </div>
        </v-card-item>
        <v-card-actions>
          <v-btn color="#6ebe4a" variant="text" @click="settings = false">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-container>
  </v-dialog>
</template>

<script>
export default {
  data() {
    return {
      help: false,
      settings: false,
    };
  },
};
</script>