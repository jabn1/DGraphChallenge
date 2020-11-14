import Vue from 'vue'
import VueRouter from 'vue-router'
import Restaurant from '../components/Restaurant.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: Restaurant
  }
]

const router = new VueRouter({
  routes
})

export default router
