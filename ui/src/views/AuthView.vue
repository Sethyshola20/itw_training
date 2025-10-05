<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-primary-500 to-secondary-500 p-5">
    <div class="bg-white rounded-xl shadow-2xl p-10 w-full max-w-md">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-gray-800 mb-2">{{ isLogin ? 'Sign In' : 'Sign Up' }}</h1>
        <p class="text-gray-600">{{ isLogin ? 'Welcome back!' : 'Create your account' }}</p>
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-5">
        <div v-if="!isLogin" class="space-y-2">
          <label for="name" class="block text-sm font-semibold text-gray-700">Full Name</label>
          <input
            id="name"
            v-model="form.name"
            type="text"
            required
            placeholder="Enter your full name"
            :disabled="isLoading"
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-lg text-base transition-colors focus:outline-none focus:border-primary-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
          />
        </div>

        <div class="space-y-2">
          <label for="email" class="block text-sm font-semibold text-gray-700">Email</label>
          <input
            id="email"
            v-model="form.email"
            type="email"
            required
            placeholder="Enter your email"
            :disabled="isLoading"
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-lg text-base transition-colors focus:outline-none focus:border-primary-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
          />
        </div>

        <div class="space-y-2">
          <label for="password" class="block text-sm font-semibold text-gray-700">Password</label>
          <input
            id="password"
            v-model="form.password"
            type="password"
            required
            placeholder="Enter your password"
            :disabled="isLoading"
            class="w-full px-4 py-3 border-2 border-gray-200 rounded-lg text-base transition-colors focus:outline-none focus:border-primary-500 disabled:bg-gray-100 disabled:cursor-not-allowed"
          />
        </div>

        <div v-if="error" class="bg-red-50 text-red-700 p-3 rounded-lg text-sm text-center">
          {{ error }}
        </div>

        <button 
          type="submit" 
          :disabled="isLoading"
          class="w-full py-3 bg-blue-500 cursor-pointer text-white font-semibold rounded-lg transition-transform hover:transform hover:-translate-y-0.5 disabled:opacity-70 disabled:cursor-not-allowed disabled:transform-none"
        >
          <span v-if="isLoading">Loading...</span>
          <span v-else>{{ isLogin ? 'Sign In' : 'Sign Up' }}</span>
        </button>
      </form>

      <div class="text-center mt-8">
        <p class="text-gray-600 text-sm">
          {{ isLogin ? "Don't have an account?" : 'Already have an account?' }}
          <button @click="toggleMode" class="cursor-pointer text-primary-500 font-semibold underline hover:text-secondary-500">
            {{ isLogin ? 'Sign Up' : 'Sign In' }}
          </button>
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const router = useRouter()
const authStore = useAuthStore()

const isLogin = ref(true)
const isLoading = computed(() => authStore.isLoading)
const error = computed(() => authStore.error)

const form = ref({
  name: '',
  email: '',
  password: ''
})

const toggleMode = () => {
  isLogin.value = !isLogin.value
  authStore.clearError()
  form.value = { name: '', email: '', password: '' }
}

const handleSubmit = async () => {
  try {
    if (isLogin.value) {
      await authStore.login({
        email: form.value.email,
        password: form.value.password
      })
    } else {
      await authStore.signup({
        name: form.value.name,
        email: form.value.email,
        password: form.value.password
      })
    }
    
    router.push('/dashboard')
  } catch (err) {
    // Error is handled by the store
  }
}
</script>
