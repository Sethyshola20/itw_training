<template>
  <div class="flex flex-col space-y-6">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-gray-800">Invoices</h1>
      <div class="flex space-x-3">
        <button @click="fetchInvoices" :disabled="isLoading" class="flex items-center space-x-2 px-4 py-2 bg-gray-100 text-gray-600 border border-gray-300 rounded-lg font-semibold transition-colors hover:bg-gray-200 disabled:opacity-70 disabled:cursor-not-allowed">
          <span class="text-sm">🔄</span>
          <span>Refresh</span>
        </button>
      </div>
    </div>

    <div v-if="isLoading" class="flex flex-col items-center justify-center py-16 text-gray-600">
      <div class="w-10 h-10 border-4 border-gray-200 border-t-primary-500 rounded-full animate-spin mb-4"></div>
      <p>Loading invoices...</p>
    </div>

    <div v-else-if="error" class="bg-red-50 text-red-700 p-5 rounded-lg text-center">
      <p>{{ error }}</p>
      <button @click="fetchInvoices" class="mt-3 px-4 py-2 bg-red-600 text-white rounded-md">Retry</button>
    </div>

    <div v-else class="space-y-6">
      <div v-if="invoices.length === 0" class="text-center py-16 text-gray-600">
        <div class="text-5xl mb-4">🧾</div>
        <h3 class="text-xl font-semibold text-gray-800 mb-2">No invoices found</h3>
        <p>There are no invoices to display at the moment.</p>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
        <div v-for="invoice in invoices" :key="invoice.id" class="bg-white rounded-xl p-6 shadow-sm transition-transform hover:-translate-y-0.5">
          <div class="flex justify-between items-start mb-4">
            <div>
              <h3 class="text-lg font-semibold text-gray-800">Invoice #{{ invoice.id.slice(-8) }}</h3>
              <span class="inline-block px-3 py-1 rounded-full text-xs font-semibold uppercase" :class="getStatusClass(invoice.status)">
                {{ invoice.status }}
              </span>
            </div>
            <div class="text-2xl font-bold text-gray-800">
              ${{ invoice.amount.toFixed(2) }}
            </div>
          </div>
          
          <div class="space-y-2 mb-5">
            <div class="flex justify-between items-center">
              <span class="text-sm font-medium text-gray-600">Created:</span>
              <span class="text-sm text-gray-800">{{ formatDate(invoice.created_at) }}</span>
            </div>
            <div class="flex justify-between items-center">
              <span class="text-sm font-medium text-gray-600">Updated:</span>
              <span class="text-sm text-gray-800">{{ formatDate(invoice.updated_at) }}</span>
            </div>
          </div>

          <div class="flex space-x-3 justify-end">
            <button @click="viewInvoiceDetails(invoice.id)" class="px-4 py-2 bg-gray-100 text-primary-500 border border-gray-300 rounded-md hover:bg-gray-200 transition-colors">
              View Details
            </button>
            <button @click="checkInvoiceStatus(invoice.id)" class="px-4 py-2 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-md hover:-translate-y-0.5 transition-transform">
              Check Status
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Invoice Details Modal -->
    <div v-if="selectedInvoice" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click="selectedInvoice = null">
      <div class="bg-white rounded-xl w-full max-w-lg max-h-screen overflow-y-auto" @click.stop>
        <div class="flex justify-between items-center p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-800">Invoice Details</h2>
          <button @click="selectedInvoice = null" class="text-gray-400 hover:text-gray-600 text-2xl">×</button>
        </div>
        
        <div class="p-6">
          <div class="space-y-4">
            <div class="flex justify-between items-center py-3 border-b border-gray-100">
              <span class="text-sm font-semibold text-gray-600">Invoice ID:</span>
              <span class="text-sm text-gray-800">{{ selectedInvoice.id }}</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-100">
              <span class="text-sm font-semibold text-gray-600">Amount:</span>
              <span class="text-sm text-gray-800">${{ selectedInvoice.amount.toFixed(2) }}</span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-100">
              <span class="text-sm font-semibold text-gray-600">Status:</span>
              <span class="text-sm px-3 py-1 rounded-full text-xs font-semibold uppercase" :class="getStatusClass(selectedInvoice.status)">
                {{ selectedInvoice.status }}
              </span>
            </div>
            <div class="flex justify-between items-center py-3 border-b border-gray-100">
              <span class="text-sm font-semibold text-gray-600">Created:</span>
              <span class="text-sm text-gray-800">{{ formatDate(selectedInvoice.created_at) }}</span>
            </div>
            <div class="flex justify-between items-center py-3">
              <span class="text-sm font-semibold text-gray-600">Updated:</span>
              <span class="text-sm text-gray-800">{{ formatDate(selectedInvoice.updated_at) }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useBillingStore, type Invoice } from '@/stores/billing'

const billingStore = useBillingStore()

const selectedInvoice = ref<Invoice | null>(null)
const isLoading = computed(() => billingStore.isLoading)
const error = computed(() => billingStore.error)
const invoices = computed(() => billingStore.invoices)

const fetchInvoices = async () => {
  try {
    await billingStore.fetchInvoices()
  } catch (err) {
    console.error('Failed to fetch invoices:', err)
  }
}

const viewInvoiceDetails = async (id: string) => {
  try {
    const invoice = await billingStore.getInvoice(id)
    selectedInvoice.value = invoice
  } catch (err) {
    console.error('Failed to fetch invoice details:', err)
  }
}

const checkInvoiceStatus = async (id: string) => {
  try {
    const status = await billingStore.getInvoiceStatus(id)
    alert(`Invoice ${id} status: ${status.invoice_status}`)
  } catch (err) {
    console.error('Failed to check invoice status:', err)
  }
}

const getStatusClass = (status: string) => {
  const statusClasses: Record<string, string> = {
    'pending': 'bg-yellow-100 text-yellow-800',
    'paid': 'bg-green-100 text-green-800',
    'failed': 'bg-red-100 text-red-800',
    'cancelled': 'bg-gray-100 text-gray-800'
  }
  return statusClasses[status.toLowerCase()] || 'bg-gray-100 text-gray-600'
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

onMounted(() => {
  fetchInvoices()
})
</script>
