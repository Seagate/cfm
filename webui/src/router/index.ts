// Copyright (c) 2024 Seagate Technology LLC and/or its Affiliates
// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [

  {
    path: '/',
    component: () => import(/* webpackChunkName: "appliances" */ '@/views/Appliances.vue'),
  },
  {
    path: '/appliances',
    name: 'HomeAppliances',
    component: () => import(/* webpackChunkName: "appliances" */ '@/views/Appliances.vue'),
  },
  {
    path: '/appliances/:appliance_id',
    name: 'Appliances',
    component: () => import(/* webpackChunkName: "appliances" */ '@/views/Appliances.vue'),
    children: [
      {
        // Optional blade_id parameter
        path: 'blades/:blade_id?',
        name: 'ApplianceWithBlade',
        component: () => import(/* webpackChunkName: "appliance" */ '@/views/Appliances.vue'),
        props: true
      }
    ]
  },
  {
    path: '/hosts',
    name: 'HomeHosts',
    component: () => import(/* webpackChunkName: "cxl-hosts" */ '@/views/CXL-Hosts.vue'),
  },
  {
    path: '/hosts/:host_id',
    name: 'CXL-Hosts',
    component: () => import(/* webpackChunkName: "cxl-hosts" */ '@/views/CXL-Hosts.vue'),
  },
  //  Hide the Alert page
  //{
  //  path: '/alerts',
  //  name: 'Alerts',
  //  component: () => import(/* webpackChunkName: "blades" */ '@/views/Alerts.vue'),
  // },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFoundPage',
    component: () => import(/* webpackChunkName: "NotFoundPage" */ '@/views/NotFoundPage.vue'),
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
