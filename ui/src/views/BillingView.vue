<template>
  <div class="flex flex-col space-y-8">
    <div class="text-center mb-5">
      <h1 class="text-4xl font-bold text-gray-800 mb-2">Billing Management</h1>
      <p class="text-gray-600">Manage your subscriptions, invoices, and payment methods</p>
    </div>

    <div class="flex flex-col space-y-8">
      <div>
        <h2 class="text-xl font-semibold text-gray-800 mb-5">Overview</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-5">
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
              <h3 class="text-sm font-semibold text-gray-600 uppercase tracking-wide mb-2">Total Revenue</h3>
              <p class="text-3xl font-bold text-gray-800">${{ totalRevenue.toFixed(2) }}</p>
            </div>
          </div>
          
          <div class="bg-white rounded-xl p-6 shadow-sm flex items-center space-x-4 transition-transform hover:-translate-y-0.5">
            <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl">
              ✅
            </div>
            <div>
              <h3 class="text-sm font-semibold text-gray-600 uppercase tracking-wide mb-2">Paid Invoices</h3>
              <p class="text-3xl font-bold text-gray-800">{{ paidInvoices }}</p>
            </div>
          </div>
        </div>
      </div>

      <div>
        <h2 class="text-xl font-semibold text-gray-800 mb-5">Recent Activity</h2>
        <div class="bg-white rounded-xl p-6 shadow-sm">
          <div class="space-y-4">
            <div v-for="activity in recentActivity" :key="activity.id" class="flex items-center space-x-4 p-4 rounded-lg bg-gray-50">
              <div class="w-10 h-10 flex items-center justify-center bg-white rounded-lg shadow-sm text-xl">
                {{ activity.icon }}
              </div>
              <div class="flex-1">
                <p class="text-sm font-medium text-gray-800 mb-1">{{ activity.text }}</p>
                <span class="text-xs text-gray-500">{{ activity.time }}</span>
              </div>
              <div class="px-3 py-1 rounded-full text-xs font-semibold uppercase" :class="activity.statusClass">
                {{ activity.status }}
              </div>
            </div>
          </div>
        </div>
      </div>
      <div>
        <h2 class="text-xl font-semibold text-gray-800 mb-5">Quick Actions</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-5">
          <router-link to="/subscriptions" class="bg-white rounded-xl p-6 shadow-sm text-center cursor-pointer transition-all hover:-translate-y-1 hover:shadow-lg group">
            <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl mx-auto mb-4">
              📋
            </div>
            <h3 class="text-lg font-semibold text-gray-800 mb-2">Manage Subscriptions</h3>
            <p class="text-sm text-gray-600 leading-relaxed">View and manage your active subscriptions</p>
          </router-link>
          
          <router-link to="/invoices" class="bg-white rounded-xl p-6 shadow-sm text-center cursor-pointer transition-all hover:-translate-y-1 hover:shadow-lg group">
            <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl mx-auto mb-4">
              🧾
            </div>
            <h3 class="text-lg font-semibold text-gray-800 mb-2">View Invoices</h3>
            <p class="text-sm text-gray-600 leading-relaxed">Check your billing history and invoices</p>
          </router-link>
          
          <div class="bg-white rounded-xl p-6 shadow-sm text-center cursor-pointer transition-all hover:-translate-y-1 hover:shadow-lg group" @click="refreshData">
            <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl mx-auto mb-4">
              🔄
            </div>
            <h3 class="text-lg font-semibold text-gray-800 mb-2">Refresh Data</h3>
            <p class="text-sm text-gray-600 leading-relaxed">Update billing information</p>
          </div>
          
          <div class="bg-white rounded-xl p-6 shadow-sm text-center cursor-pointer transition-all hover:-translate-y-1 hover:shadow-lg group" @click="exportData">
            <div class="w-15 h-15 flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 rounded-xl text-3xl mx-auto mb-4">
              📊
            </div>
            <h3 class="text-lg font-semibold text-gray-800 mb-2">Export Data</h3>
            <p class="text-sm text-gray-600 leading-relaxed">Download billing reports</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useBillingStore } from '@/stores/billing'

const billingStore = useBillingStore()

const subscriptionCount = computed(() => billingStore.subscriptions.length)
const invoiceCount = computed(() => billingStore.invoices.length)
const totalRevenue = computed(() => 
  billingStore.invoices.reduce((sum, invoice) => sum + invoice.amount, 0)
)
const paidInvoices = computed(() => 
  billingStore.invoices.filter(invoice => invoice.status.toLowerCase() === 'paid').length
)

const recentActivity = ref([
  {
    id: 1,
    icon: '📋',
    text: 'New subscription created',
    time: '2 minutes ago',
    status: 'Active',
    statusClass: 'bg-green-100 text-green-800'
  },
  {
    id: 2,
    icon: '💰',
    text: 'Payment processed successfully',
    time: '1 hour ago',
    status: 'Completed',
    statusClass: 'bg-blue-100 text-blue-800'
  },
  {
    id: 3,
    icon: '🧾',
    text: 'Invoice generated',
    time: '3 hours ago',
    status: 'Pending',
    statusClass: 'bg-yellow-100 text-yellow-800'
  },
  {
    id: 4,
    icon: '📊',
    text: 'Monthly report generated',
    time: '1 day ago',
    status: 'Completed',
    statusClass: 'bg-blue-100 text-blue-800'
  }
])

const refreshData = async () => {
  try {
    await Promise.all([
      billingStore.fetchSubscriptions(),
      billingStore.fetchInvoices()
    ])
  } catch (error) {
    console.error('Failed to refresh data:', error)
  }
}

const exportData = () => {
  const data = {
    subscriptions: billingStore.subscriptions,
    invoices: billingStore.invoices,
    summary: {
      totalSubscriptions: subscriptionCount.value,
      totalInvoices: invoiceCount.value,
      totalRevenue: totalRevenue.value,
      paidInvoices: paidInvoices.value
    }
  }
  
  const blob = new Blob([JSON.stringify(data, null, 2)], { type: 'application/json' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `billing-data-${new Date().toISOString().split('T')[0]}.json`
  document.body.appendChild(a)
  a.click()
  document.body.removeChild(a)
  URL.revokeObjectURL(url)
}

onMounted(() => {
  refreshData()
})
</script>
