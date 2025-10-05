<template>
  <div class="flex flex-col space-y-8">
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-5">
      <div class="bg-white rounded-xl p-6 shadow-sm flex items-center space-x-4 transition-transform hover:-translate-y-0.5">
        <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl">
          👥
        </div>
        <div>
          <h3 class="text-sm font-semibold text-gray-600 uppercase tracking-wide mb-2">Total Users</h3>
          <p class="text-3xl font-bold text-gray-800">{{ userCount }}</p>
        </div>
      </div>
      
      <div class="bg-white rounded-xl p-6 shadow-sm flex items-center space-x-4 transition-transform hover:-translate-y-0.5">
        <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl">
          📋
        </div>
        <div>
          <h3 class="text-sm font-semibold text-gray-600 uppercase tracking-wide mb-2">Active Subscriptions</h3>
          <p class="text-3xl font-bold text-gray-800">{{ subscriptionCount }}</p>
        </div>
      </div>
      
      <div class="bg-white rounded-xl p-6 shadow-sm flex items-center space-x-4 transition-transform hover:-translate-y-0.5">
        <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl">
          🧾
        </div>
        <div>
          <h3 class="text-sm font-semibold text-gray-600 uppercase tracking-wide mb-2">Total Invoices</h3>
          <p class="text-3xl font-bold text-gray-800">{{ invoiceCount }}</p>
        </div>
      </div>
      
      <div class="bg-white rounded-xl p-6 shadow-sm flex items-center space-x-4 transition-transform hover:-translate-y-0.5">
        <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl">
          💰
        </div>
        <div>
          <h3 class="text-sm font-semibold text-gray-600 uppercase tracking-wide mb-2">Revenue</h3>
          <p class="text-3xl font-bold text-gray-800">${{ totalRevenue.toFixed(2) }}</p>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <div class="bg-white rounded-xl p-6 shadow-sm">
        <h2 class="text-xl font-semibold text-gray-800 mb-5">Recent Activity</h2>
        <div class="space-y-4">
          <div v-for="activity in recentActivity" :key="activity.id" class="flex items-center space-x-3 p-3 rounded-lg bg-gray-50">
            <div class="w-10 h-10 flex items-center justify-center bg-white rounded-lg shadow-sm text-xl">
              {{ activity.icon }}
            </div>
            <div class="flex-1">
              <p class="text-sm font-medium text-gray-800 mb-1">{{ activity.text }}</p>
              <span class="text-xs text-gray-500">{{ activity.time }}</span>
            </div>
          </div>
        </div>
      </div>

      <div class="bg-white rounded-xl p-6 shadow-sm">
        <h2 class="text-xl font-semibold text-gray-800 mb-5">Quick Actions</h2>
        <div class="space-y-3">
          <router-link to="/subscriptions" class="flex items-center space-x-3 p-4 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold transition-transform hover:-translate-y-0.5">
            <span class="text-lg">➕</span>
            Create Subscription
          </router-link>
          <router-link to="/invoices" class="flex items-center space-x-3 p-4 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold transition-transform hover:-translate-y-0.5">
            <span class="text-lg">📄</span>
            View Invoices
          </router-link>
          <router-link to="/billing" class="flex items-center space-x-3 p-4 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold transition-transform hover:-translate-y-0.5">
            <span class="text-lg">💳</span>
            Manage Billing
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useBillingStore } from '@/stores/billing'

const billingStore = useBillingStore()

const userCount = ref(0)
const subscriptionCount = ref(0)
const invoiceCount = ref(0)
const totalRevenue = ref(0)

const recentActivity = ref([
  {
    id: 1,
    icon: '👤',
    text: 'New user registered',
    time: '2 minutes ago'
  },
  {
    id: 2,
    icon: '📋',
    text: 'Subscription created',
    time: '15 minutes ago'
  },
  {
    id: 3,
    icon: '💰',
    text: 'Payment processed',
    time: '1 hour ago'
  },
  {
    id: 4,
    icon: '🧾',
    text: 'Invoice generated',
    time: '2 hours ago'
  }
])

onMounted(async () => {
  try {
    await Promise.all([
      billingStore.fetchSubscriptions(),
      billingStore.fetchInvoices()
    ])
    
    subscriptionCount.value = billingStore.subscriptions.length
    invoiceCount.value = billingStore.invoices.length
    totalRevenue.value = billingStore.invoices.reduce((sum, invoice) => sum + invoice.amount, 0)
  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  }
})
</script>
