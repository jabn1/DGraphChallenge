import Vue from 'vue'
import VueRouter from 'vue-router'
import Restaurant from '../components/Restaurant.vue'
import Buyerlist from '../components/Buyerlist.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/home',
    component: Restaurant
  },
  {
    path: '/buyers',
    component: Buyerlist
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
