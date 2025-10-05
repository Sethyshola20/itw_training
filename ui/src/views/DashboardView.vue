<template>
  <div class="flex min-h-screen bg-gray-50">
    <main class="flex-1 flex flex-col">
      <header class="bg-white px-8 py-5 border-b border-gray-200 flex justify-between items-center shadow-sm">
        <h1 class="text-2xl font-semibold text-gray-800">{{ pageTitle }}</h1>
        <div class="text-gray-600 text-sm">
          <span>{{ user?.email }}</span>
        </div>
      </header>
      
      <div class="flex-1 p-8 overflow-y-auto">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import Sidebar from '@/components/Sidebar.vue'
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import DashboardContent from '@/components/DashboardContent.vue'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const user = computed(() => authStore.user)

const pageTitle = computed(() => {
  const titles: Record<string, string> = {
    '/dashboard': 'Dashboard',
    '/billing': 'Billing Management',
    '/subscriptions': 'Subscriptions',
    '/invoices': 'Invoices',
    '/settings': 'Settings'
  }
  return titles[route.path] || 'Dashboard'
})

const handleLogout = () => {
  authStore.logout()
  router.push('/auth')
}
</script>
