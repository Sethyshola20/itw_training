<template>
      <main class="flex min-h-screen">
        <Sidebar />
        <div class="p-6 component">
            <router-view />
        </div>
      </main>
</template>
  
<script setup>
import Sidebar from '@/components/Sidebar.vue'
</script>
  