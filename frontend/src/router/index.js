import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Send from '../views/Send.vue'
import Receive from '../views/Receive.vue'
import Admin from '../views/Admin.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/send', component: Send },
  { path: '/receive', component: Receive },
  { path: '/admin', component: Admin }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
