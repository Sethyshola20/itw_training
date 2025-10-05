<template>
  <div class="flex flex-col space-y-8">
    <div class="text-center mb-5">
      <h1 class="text-4xl font-bold text-gray-800 mb-2">Settings</h1>
      <p class="text-gray-600">Manage your account settings and preferences</p>
    </div>

    <div class="flex flex-col space-y-8">
      <!-- User Profile -->
      <div>
        <h2 class="text-xl font-semibold text-gray-800 mb-5">User Profile</h2>
        <div class="bg-white rounded-xl p-6 shadow-sm flex justify-between items-center">
          <div class="flex items-center space-x-4">
            <div class="w-15 h-15 bg-gradient-to-br from-primary-500 to-secondary-500 text-white rounded-full flex items-center justify-center text-2xl font-bold">
              {{ user?.name?.charAt(0)?.toUpperCase() || 'U' }}
            </div>
            <div>
              <h3 class="text-lg font-semibold text-gray-800">{{ user?.name || 'Unknown User' }}</h3>
              <p class="text-gray-600">{{ user?.email || 'No email' }}</p>
              <span class="text-xs text-gray-500">
                Member since {{ formatDate(user?.created_at) }}
              </span>
            </div>
          </div>
          <button @click="editProfile = true" class="px-5 py-2 bg-gray-100 text-primary-500 border border-gray-300 rounded-lg font-semibold hover:bg-gray-200 transition-colors">
            Edit Profile
          </button>
        </div>
      </div>

      <!-- Account Settings -->
      <div>
        <h2 class="text-xl font-semibold text-gray-800 mb-5">Account Settings</h2>
        <div class="bg-white rounded-xl p-6 shadow-sm space-y-6">
          <div class="flex justify-between items-center py-4 border-b border-gray-100">
            <div>
              <h3 class="text-base font-semibold text-gray-800">Email Notifications</h3>
              <p class="text-sm text-gray-600">Receive email updates about your account</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" v-model="settings.emailNotifications" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>
          
          <div class="flex justify-between items-center py-4 border-b border-gray-100">
            <div>
              <h3 class="text-base font-semibold text-gray-800">Billing Alerts</h3>
              <p class="text-sm text-gray-600">Get notified about upcoming payments</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" v-model="settings.billingAlerts" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>
          
          <div class="flex justify-between items-center py-4">
            <div>
              <h3 class="text-base font-semibold text-gray-800">Two-Factor Authentication</h3>
              <p class="text-sm text-gray-600">Add an extra layer of security to your account</p>
            </div>
            <label class="relative inline-flex items-center cursor-pointer">
              <input type="checkbox" v-model="settings.twoFactorAuth" class="sr-only peer">
              <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-primary-300 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-primary-500"></div>
            </label>
          </div>
        </div>
      </div>

      <!-- Billing Information -->
      <div>
        <h2 class="text-xl font-semibold text-gray-800 mb-5">Billing Information</h2>
        <div class="grid grid-cols-1 md:grid-cols-3 gap-5">
          <div class="bg-white rounded-xl p-6 shadow-sm flex flex-col space-y-3">
            <h3 class="text-lg font-semibold text-gray-800">Payment Methods</h3>
            <p class="text-sm text-gray-600 leading-relaxed">Manage your payment methods and billing preferences</p>
            <button class="px-4 py-2 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold hover:-translate-y-0.5 transition-transform self-start">
              Manage Payment Methods
            </button>
          </div>
          
          <div class="bg-white rounded-xl p-6 shadow-sm flex flex-col space-y-3">
            <h3 class="text-lg font-semibold text-gray-800">Billing Address</h3>
            <p class="text-sm text-gray-600 leading-relaxed">Update your billing address for invoices</p>
            <button class="px-4 py-2 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold hover:-translate-y-0.5 transition-transform self-start">
              Update Address
            </button>
          </div>
          
          <div class="bg-white rounded-xl p-6 shadow-sm flex flex-col space-y-3">
            <h3 class="text-lg font-semibold text-gray-800">Tax Information</h3>
            <p class="text-sm text-gray-600 leading-relaxed">Configure tax settings for your region</p>
            <button class="px-4 py-2 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold hover:-translate-y-0.5 transition-transform self-start">
              Configure Taxes
            </button>
          </div>
        </div>
      </div>

      <!-- Danger Zone -->
      <div class="border-t-2 border-red-200 pt-8">
        <h2 class="text-xl font-semibold text-red-600 mb-5">Danger Zone</h2>
        <div class="bg-white rounded-xl p-6 shadow-sm">
          <div class="flex justify-between items-center">
            <div>
              <h3 class="text-base font-semibold text-red-600">Delete Account</h3>
              <p class="text-sm text-gray-600">Permanently delete your account and all associated data</p>
            </div>
            <button @click="showDeleteConfirm = true" class="px-5 py-2 bg-red-50 text-red-600 border border-red-200 rounded-lg font-semibold hover:bg-red-100 transition-colors">
              Delete Account
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Edit Profile Modal -->
    <div v-if="editProfile" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click="editProfile = false">
      <div class="bg-white rounded-xl w-full max-w-md max-h-screen overflow-y-auto" @click.stop>
        <div class="flex justify-between items-center p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-800">Edit Profile</h2>
          <button @click="editProfile = false" class="text-gray-400 hover:text-gray-600 text-2xl">×</button>
        </div>
        
        <form @submit.prevent="saveProfile" class="p-6">
          <div class="space-y-2 mb-5">
            <label for="name" class="block text-sm font-semibold text-gray-700">Full Name</label>
            <input
              id="name"
              v-model="profileForm.name"
              type="text"
              required
              placeholder="Enter your full name"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-lg focus:outline-none focus:border-primary-500"
            />
          </div>
          
          <div class="space-y-2 mb-5">
            <label for="email" class="block text-sm font-semibold text-gray-700">Email</label>
            <input
              id="email"
              v-model="profileForm.email"
              type="email"
              required
              placeholder="Enter your email"
              class="w-full px-4 py-3 border-2 border-gray-200 rounded-lg focus:outline-none focus:border-primary-500"
            />
          </div>

          <div class="flex space-x-3 justify-end">
            <button type="button" @click="editProfile = false" class="px-5 py-2 bg-gray-100 text-gray-600 border border-gray-300 rounded-lg font-semibold">
              Cancel
            </button>
            <button type="submit" class="px-5 py-2 bg-gradient-to-r from-primary-500 to-secondary-500 text-white rounded-lg font-semibold">
              Save Changes
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteConfirm" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click="showDeleteConfirm = false">
      <div class="bg-white rounded-xl w-full max-w-md max-h-screen overflow-y-auto" @click.stop>
        <div class="flex justify-between items-center p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-800">Delete Account</h2>
          <button @click="showDeleteConfirm = false" class="text-gray-400 hover:text-gray-600 text-2xl">×</button>
        </div>
        
        <div class="p-6">
          <div class="text-center py-5">
            <div class="text-5xl mb-4">⚠️</div>
            <h3 class="text-xl font-semibold text-red-600 mb-3">Are you sure?</h3>
            <p class="text-sm text-gray-600 leading-relaxed mb-6">This action cannot be undone. This will permanently delete your account and remove all data from our servers.</p>
            
            <div class="flex space-x-3 justify-center">
              <button @click="showDeleteConfirm = false" class="px-5 py-2 bg-gray-100 text-gray-600 border border-gray-300 rounded-lg font-semibold">
                Cancel
              </button>
              <button @click="deleteAccount" class="px-5 py-2 bg-red-50 text-red-600 border border-red-200 rounded-lg font-semibold hover:bg-red-100 transition-colors">
                Yes, Delete Account
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const user = computed(() => authStore.user)
const editProfile = ref(false)
const showDeleteConfirm = ref(false)

const profileForm = ref({
  name: '',
  email: ''
})

const settings = ref({
  emailNotifications: true,
  billingAlerts: true,
  twoFactorAuth: false
})

const saveProfile = () => {
  // Mock save functionality
  console.log('Saving profile:', profileForm.value)
  editProfile.value = false
}

const deleteAccount = () => {
  // Mock delete functionality
  console.log('Deleting account')
  authStore.logout()
  router.push('/auth')
}

const formatDate = (dateString?: string) => {
  if (!dateString) return 'Unknown'
  return new Date(dateString).toLocaleDateString()
}

onMounted(() => {
  if (user.value) {
    profileForm.value = {
      name: user.value.name,
      email: user.value.email
    }
  }
})
</script>
