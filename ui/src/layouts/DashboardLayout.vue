<template>
    <div>
      <main class="flex">
        <Sidebar />
        <div class="p-6 component">
            <router-view />
        </div>
      </main>
    </div>
  </template>
  
  <script setup>
  import Sidebar from '@/components/Sidebar.vue'
  </script>
  