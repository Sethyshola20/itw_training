import { defineStore } from 'pinia'
import { ref } from 'vue'
import { apiService, type Subscription, type Invoice } from '@/services/api'

export type { Subscription, Invoice }

export const useBillingStore = defineStore('billing', () => {
  const subscriptions = ref<Subscription[]>([])
  const invoices = ref<Invoice[]>([])
  const isLoading = ref(false)
  const error = ref<string | null>(null)

  const fetchSubscriptions = async () => {
    isLoading.value = true
    error.value = null
    
    try {
      subscriptions.value = await apiService.getSubscriptions()
      console.log("subs",subscriptions )
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch subscriptions'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const createSubscription = async (subscription: Partial<Subscription>) => {
    isLoading.value = true
    error.value = null
    
    try {
      const newSubscription = await apiService.createSubscription(subscription)
      subscriptions.value.push(newSubscription)
      return newSubscription
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to create subscription'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const deleteSubscription = async (id: string) => {
    isLoading.value = true
    error.value = null
    
    try {
      await apiService.deleteSubscription(id)
      subscriptions.value = subscriptions.value.filter(sub => sub.id !== id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to delete subscription'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const fetchInvoices = async () => {
    isLoading.value = true
    error.value = null
    
    try {
      invoices.value = await apiService.getInvoices()
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch invoices'
      throw err
    } finally {
      isLoading.value = false
    }
  }

  const getInvoice = async (id: string) => {
    try {
      return await apiService.getInvoice(id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch invoice'
      throw err
    }
  }

  const getInvoiceStatus = async (id: string) => {
    try {
      return await apiService.getInvoiceStatus(id)
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to fetch invoice status'
      throw err
    }
  }

  const clearError = () => {
    error.value = null
  }

  return {
    subscriptions,
    invoices,
    isLoading,
    error,
    fetchSubscriptions,
    createSubscription,
    deleteSubscription,
    fetchInvoices,
    getInvoice,
    getInvoiceStatus,
    clearError
  }
})