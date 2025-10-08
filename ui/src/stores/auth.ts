import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { apiService, type User, type LoginRequest, type SignupRequest } from '@/services/api'

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(null)
  const isLoading = ref(false)
  const error = ref<string | null>(null)


  const isAuthenticated = computed(() => !!user.value)


  const login = async (credentials: LoginRequest) => {
    isLoading.value = true
    error.value = null

    try {
      await apiService.login(credentials)
      await fetchCurrentUser()
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Login failed'
      throw err
    } finally {
      isLoading.value = false
    }
  }
  
  const signup = async (userData: SignupRequest) => {
    isLoading.value = true
    error.value = null

    try {
      await apiService.signup(userData)
      await fetchCurrentUser()
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Signup failed'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchCurrentUser = async () => {
    try {
      const res = await apiService.getCurrentUser()
      console.log('user ? : ', res)
      user.value = res
      return res
    } catch (err) {
      console.error('Failed to fetch user:', err)
      logout()
    }
  }

  const logout = async () => {
    user.value = null
    error.value = null
    try {
      // Optional: call backend endpoint to clear cookies
      // await apiService.logout()
    } catch {}
  }

  const clearError = () => {
    error.value = null
  }

  fetchCurrentUser()

  return {
    user,
    isLoading,
    error,
    isAuthenticated,
    login,
    signup,
    logout,
    clearError,
    fetchCurrentUser,
  }
}, { persist: true })
