<template>
  <nav class="w-72 bg-black sticky h-[100vh] text-white flex flex-col shadow-lg">
      <div class="p-8 border-b border-white/10">
        <h2 class="text-xl font-bold mb-2">Scaleway</h2>
        <p class="text-sm opacity-80">Welcome, {{ user?.name || 'User' }}</p>
      </div>
      <ul class="flex-1 py-5 bg-black">
        <li>
          <router-link to="/dashboard" class="flex items-center px-5 py-4 text-white transition-colors hover:bg-white/10 border-l-4 border-transparent hover:border-white router-link-active:bg-white/15 router-link-active:border-white">
            <span class="mr-3 text-lg">📊</span>
            Dashboard
          </router-link>
        </li>
        <li>
          <router-link to="/billing" class="flex items-center px-5 py-4 text-white transition-colors hover:bg-white/10 border-l-4 border-transparent hover:border-white router-link-active:bg-white/15 router-link-active:border-white">
            <span class="mr-3 text-lg">💳</span>
            Billing
          </router-link>
        </li>
        <li>
          <router-link to="/subscriptions" class="flex items-center px-5 py-4 text-white transition-colors hover:bg-white/10 border-l-4 border-transparent hover:border-white router-link-active:bg-white/15 router-link-active:border-white">
            <span class="mr-3 text-lg">📋</span>
            Subscriptions
          </router-link>
        </li>
        <li>
          <router-link to="/invoices" class="flex items-center px-5 py-4 text-white transition-colors hover:bg-white/10 border-l-4 border-transparent hover:border-white router-link-active:bg-white/15 router-link-active:border-white">
            <span class="mr-3 text-lg">🧾</span>
            Invoices
          </router-link>
        </li>
        <li>
          <router-link to="/settings" class="flex items-center px-5 py-4 text-white transition-colors hover:bg-white/10 border-l-4 border-transparent hover:border-white router-link-active:bg-white/15 router-link-active:border-white">
            <span class="mr-3 text-lg">⚙️</span>
            Settings
          </router-link>
        </li>
      </ul>
      <div class="p-5 border-t border-white/10">
        <button @click="handleLogout" class="flex items-center w-full px-4 py-3 bg-white/10 border border-white/20 rounded-lg text-white cursor-pointer transition-colors hover:bg-white/20">
          <span class="mr-3 text-lg">🚪</span>
          Logout
        </button>
      </div>
    </nav>
</template>
<script>

</script>