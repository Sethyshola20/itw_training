// router/index.ts
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import DashboardLayout from '@/layouts/DashboardLayout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    // Protected area (all behind DashboardLayout)
    {
      path: '/',
      component: DashboardLayout,
      meta: { requiresAuth: true },
      children: [
        { path: '', redirect: { name: 'dashboard' } }, // 👈 default redirect
        { path: 'dashboard', name:'dashboard', component: () => import('@/views/DashboardView.vue')  },
        { path: 'billing', name: 'billing', component: () => import('@/views/BillingView.vue') },
        { path: 'subscriptions', name: 'subscriptions', component: () => import('@/views/SubscriptionsView.vue') },
        { path: 'invoices', name: 'invoices', component: () => import('@/views/InvoicesView.vue') },
        { path: 'settings', name: 'settings', component: () => import('@/views/SettingsView.vue') },
      ],
    },

    // Auth page
    {
      path: '/auth',
      name: 'auth',
      component: () => import('@/views/AuthView.vue'),
      meta: { requiresGuest: true },
    },
  ],
})

// Navigation guards
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return next({ name: 'auth' })
  }

  if (to.meta.requiresGuest && authStore.isAuthenticated) {
    return next({ name: 'billing' }) // 👈 or whichever child is your "home"
  }

  next()
})

export default router