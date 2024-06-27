<!-- Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates -->
<template>
  <v-layout class="rounded rounded-md">
    <v-app-bar app density="compact">
      <Header />
    </v-app-bar>

    <v-footer app density="compact">
      <Footer />
    </v-footer>

    <v-navigation-drawer
      app
      permanent
      image="@/assets/iStock-1090069346_white_bg_green_dots.jpg"
      theme="light"
    >
      <SideBar />
    </v-navigation-drawer>

    <v-main
      class="d-flex align-center justify-center"
      style="min-height: 300px"
    >
      <router-view />
    </v-main>
  </v-layout>
</template>

<script setup>
import SideBar from "./components/SideBar.vue";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
</script>
