<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="min-h-screen bg-gray-100">
    <!-- Barra de navegación -->
    <nav class="bg-white shadow">
      <div class="container mx-auto px-6 py-4">
        <div class="flex items-center justify-between">
          <h1 class="text-xl font-bold text-gray-800">Stock Analyzer</h1>
          <div class="flex space-x-4">
            <RouterLink 
              to="/stocks"
              class="px-4 py-2 text-gray-600 hover:text-blue-600 transition-colors"
              active-class="text-blue-600 border-b-2 border-blue-600"
            >
              Stocks
            </RouterLink>
            <RouterLink 
              to="/recommendations"
              class="px-4 py-2 text-gray-600 hover:text-blue-600 transition-colors"
              active-class="text-blue-600 border-b-2 border-blue-600"
            >
              Recommendations
            </RouterLink>
          </div>
        </div>
      </div>
    </nav>

    <!-- Contenido de la página -->
    <main class="container mx-auto p-6">
      <RouterView />
    </main>
  </div>
</template>