// API service for microservices communication
const API_BASE_URL = 'http://localhost:8080'

export interface User {
  id: string
  email: string
  name: string
  created_at: string
}

export interface LoginRequest {
  email: string
  password: string
}

export interface SignupRequest extends LoginRequest {
  name: string
}

export interface AuthResponse {
  user:  Omit <User, "created_at">
}

export interface Subscription {
  id: string
  user_id: string
  plan: string
  created_at: string
  updated_at: string
}

export interface Invoice {
  id: string
  amount: number
  status: string
  created_at: string
  updated_at: string
}

class ApiService {
  private baseURL: string

  constructor(baseURL: string = API_BASE_URL) {
    this.baseURL = baseURL
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    
    const config: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        ...options.headers,
      },
      ...options,
      credentials:'include'
    }

    const response = await fetch(`${this.baseURL}${endpoint}`, config)
    
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    return response.json()
  }

  // Auth endpoints
  async login(credentials: LoginRequest): Promise<AuthResponse> {
    return this.request<AuthResponse>('/api/token/', {
      method: 'POST',
      body: JSON.stringify(credentials),
    })
  }

  async signup(userData: SignupRequest): Promise<AuthResponse> {
    return this.request<AuthResponse>('/api/signup/', {
      method: 'POST',
      body: JSON.stringify(userData),
    })
  }

  async refreshToken(): Promise<AuthResponse> {
    const refreshToken = localStorage.getItem('refresh_token')
    return this.request<AuthResponse>('/api/token/refresh/', {
      method: 'POST',
      body: JSON.stringify({ refresh: refreshToken }),
    })
  }

  // User endpoints
  async getCurrentUser(): Promise<User> {
    return this.request<User>('/api/users/me/')
  }

  async getUsers(): Promise<User[]> {
    return this.request<User[]>('/api/users/')
  }

  // Subscription endpoints
  async getSubscriptions(): Promise<Subscription[]> {
    return this.request<Subscription[]>('/api/billing/subscriptions')
  }

  async createSubscription(subscription: Partial<Subscription>): Promise<Subscription> {
    return this.request<Subscription>('/api/billing/subscriptions/', {
      method: 'POST',
      body: JSON.stringify(subscription),
    })
  }

  async deleteSubscription(id: string): Promise<void> {
    return this.request<void>(`/api/billing/subscriptions/${id}`, {
      method: 'DELETE',
    })
  }

  // Invoice endpoints
  async getInvoices(): Promise<Invoice[]> {
    return this.request<Invoice[]>('/api/billing/invoices')
  }

  async getInvoice(id: string): Promise<Invoice> {
    return this.request<Invoice>(`/api/billing/invoices/${id}`)
  }

  async getInvoiceStatus(id: string): Promise<{ invoice_status: string }> {
    return this.request<{ invoice_status: string }>(`/api/billing/invoices/${id}/status`)
  }
}

export const apiService = new ApiService()