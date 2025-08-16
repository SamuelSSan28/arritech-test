import { createRouter, createWebHistory } from 'vue-router'
import UsersList from '@/views/UsersList.vue'

const routes = [
  {
    path: '/',
    name: 'UsersList',
    component: UsersList,
    meta: {
      title: 'Users Management'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Update page title based on route meta
router.beforeEach((to, from, next) => {
  document.title = to.meta.title ? `${to.meta.title} - Arritech User Management` : 'Arritech User Management'
  next()
})

export default router 