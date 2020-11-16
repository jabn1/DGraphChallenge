import Vue from 'vue'
import VueRouter from 'vue-router'
import Restaurant from '../components/Restaurant.vue'
import Buyerlist from '../components/Buyerlist.vue'
import Buyerhistory from '../components/Buyerhistory.vue'
import Loaddate from '../components/Loaddate.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/home',
    component: Restaurant
  },
  {
    path: '/buyers',
    component: Buyerlist
  },
  {
    path: '/buyer/:id',
    component: Buyerhistory
  },
  {
    path: '/buyer',
    component: Buyerhistory
  },
  {
    path: '/load',
    component: Loaddate
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
