// import { createRouter, createWebHistory } from 'vue-router'
import { createRouter, createWebHistory } from '@/router/vue-router.mjs'
import HomeView from '../views/HomeView.vue'
import Login from '../views/Login.vue'
import axios from 'axios'
import { useData } from '../composables/useData.js'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/private/another',
      name: 'another',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AnotherView.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/private/new',
      name: 'new',
      component: () => import('../views/NewView.vue')
    }
  ]
})

router.beforeEach(async (to, from) => {
  const data  = await useData(to.path, router)
  window.data = data

})



export default router
