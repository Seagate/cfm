// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [

  {
    path: '/',
    component: () => import(/* webpackChunkName: "appliances" */ '@/views/Appliances.vue'),
  },
  {
    path: '/appliances',
    name: 'Appliances',
    component: () => import(/* webpackChunkName: "appliances" */ '@/views/Appliances.vue'),
  },
  {
    path: '/hosts',
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
