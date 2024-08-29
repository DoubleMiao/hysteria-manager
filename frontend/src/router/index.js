import { createRouter, createWebHistory } from 'vue-router'
import InstanceList from '../components/InstanceList.vue'
import CreateInstance from '../components/CreateInstance.vue'

const routes = [
  { path: '/', component: InstanceList },
  { path: '/create', component: CreateInstance },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
