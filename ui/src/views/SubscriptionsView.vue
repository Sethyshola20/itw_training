<template>
  <div class="flex flex-col space-y-6">
    <div class="flex justify-between items-center">
      <h1 class="text-3xl font-bold text-gray-800">Subscriptions</h1>
      <button @click="showCreateModal = true" class="flex items-center space-x-2 px-5 py-3 bg-gradient-to-r from-primary-500 to-secondary-500 text-white font-semibold rounded-lg transition-transform hover:-translate-y-0.5">
        <span class="text-lg">➕</span>
        <span>Create Subscription</span>
      </button>
    </div>

    <div v-if="isLoading" class="flex flex-col items-center justify-center py-16 text-gray-600">
      <div class="w-10 h-10 border-4 border-gray-200 border-t-primary-500 rounded-full animate-spin mb-4"></div>
      <p>Loading subscriptions...</p>
    </div>

    <div v-else-if="error" class="bg-red-50 text-red-700 p-5 rounded-lg text-center">
      <p>{{ error }}</p>
      <button @click="fetchSubscriptions" class="mt-3 px-4 py-2 bg-red-600 text-white rounded-md">Retry</button>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
      <div v-for="subscription in subscriptions" :key="subscription.id" class="bg-white rounded-xl p-6 shadow-sm transition-transform hover:-translate-y-0.5">
        <div class="flex justify-between items-start mb-4">
          <h3 class="text-lg font-semibold text-gray-800 capitalize">{{ subscription.plan }}</h3>
          <span class="text-xs text-gray-500 font-mono">#{{ subscription.id.slice(-8) }}</span>
        </div>
        
        <div class="space-y-2 mb-5">
          <div class="flex justify-between items-center">
            <span class="text-sm font-medium text-gray-600">User ID:</span>
            <span class="text-sm text-gray-800">{{ subscription.user_id }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm font-medium text-gray-600">Created:</span>
            <span class="text-sm text-gray-800">{{ formatDate(subscription.created_at) }}</span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm font-medium text-gray-600">Updated:</span>
            <span class="text-sm text-gray-800">{{ formatDate(subscription.updated_at) }}</span>
          </div>
        </div>

        <div class="flex justify-end">
          <button @click="deleteSubscription(subscription.id)" class="px-4 py-2 bg-red-50 text-red-600 border border-red-200 rounded-md hover:bg-red-100 transition-colors">
            Delete
          </button>
        </div>
      </div>
    </div>

    <!-- Create Subscription Modal -->
    <div v-if="showCreateModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click="showCreateModal = false">
      <div class="bg-white rounded-xl w-full max-w-md max-h-screen overflow-y-auto" @click.stop>
        <div class="flex justify-between items-center p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-800">Create New Subscription</h2>
          <button @click="showCreateModal = false" class="text-gray-400 hover:text-gray-600 text-2xl">×</button>
        </div>
        
        <form @submit.prevent="handleCreateSubscription" class="p-6">
          <div class="space-y-2 mb-5">
            <label for="user_id" class="block text-sm font-semibold text-gray-700">User ID</label>
            <input
              id="user_id"
              v-model="newSubscription.user_id"
              type="text"
              required
              placeholder="Enter user ID"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-lg focus:outline-none focus:border-primary-500"
            />
          </div>
          
          <div class="space-y-2 mb-5">
            <label for="plan" class="block text-sm font-semibold text-gray-700">Plan</label>
            <select id="plan" v-model="newSubscription.plan" required class="w-full px-4 py-3 border-2 border-gray-200 rounded-lg focus:outline-none focus:border-primary-500">
              <option value="">Select a plan</option>
              <option value="free">Free</option>
              <option value="pro">Pro</option>
              <option value="enterprise">Enterprise</option>
            </select>
          </div>

          <div class="flex space-x-3 justify-end">
            <button type="button" @click="showCreateModal = false" class="px-5 py-2 bg-gray-100 text-gray-600 border border-gray-300 rounded-lg font-semibold">
              Cancel
            </button>
            <button type="submit" :disabled="isLoading" class="px-5 py-2 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold disabled:opacity-70">
              {{ isLoading ? 'Creating...' : 'Create Subscription' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useBillingStore } from '@/stores/billing'

const billingStore = useBillingStore()

const showCreateModal = ref(false)
const isLoading = computed(() => billingStore.isLoading)
const error = computed(() => billingStore.error)
const subscriptions = computed(() => billingStore.subscriptions)

const newSubscription = ref({
  user_id: '',
  plan: ''
})

const fetchSubscriptions = async () => {
  try {
    await billingStore.fetchSubscriptions()
  } catch (err) {
    console.error('Failed to fetch subscriptions:', err)
  }
}

const handleCreateSubscription = async () => {
  try {
    await billingStore.createSubscription(newSubscription.value)
    showCreateModal.value = false
    newSubscription.value = { user_id: '', plan: '' }
  } catch (err) {
    console.error('Failed to create subscription:', err)
  }
}

const deleteSubscription = async (id: string) => {
  if (confirm('Are you sure you want to delete this subscription?')) {
    try {
      await billingStore.deleteSubscription(id)
    } catch (err) {
      console.error('Failed to delete subscription:', err)
    }
  }
}

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString()
}

onMounted(() => {
  fetchSubscriptions()
})
</script>
