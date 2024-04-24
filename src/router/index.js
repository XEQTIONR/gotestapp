// import { createRouter, createWebHistory } from 'vue-router'
import { createRouter, createWebHistory } from '@/router/vue-router.mjs'
import HomeView from '../views/HomeView.vue'
import axios from 'axios'

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
      path: '/another',
      name: 'another',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AnotherView.vue')
    }
  ]
})

// router.beforeEach((to, from) => {
//   console.log('before each from')
//   console.log('to:')
//   console.log(to)
//   console.log('from:')
//   console.log(from)


//   const fullPath = to.fullPath

//   axios.get(fullPath)
//     .then((res) => {
//       console.log('response')
//       console.log(res)
//     })
// })



export default router
