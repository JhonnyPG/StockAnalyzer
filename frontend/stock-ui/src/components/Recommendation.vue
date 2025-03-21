<script setup lang="ts">
import { useRecommendationStore } from '@/stores/recommendationStore'
import { onMounted } from 'vue'

const store = useRecommendationStore()

onMounted(() => {
    if (!store.bestStock) {
        store.fetchBestStock()
    }
})
</script>

<template>
    <div class="bg-white rounded-lg shadow-lg p-6 mb-6">
      <h2 class="text-2xl font-bold text-gray-800 mb-4">🚀 Best Daily Recommendation</h2>
  
      <!-- Estado de carga -->
      <div v-if="store.loading" class="animate-pulse space-y-4">
        <div class="h-4 bg-gray-200 rounded w-1/2"></div>
        <div class="h-4 bg-gray-200 rounded"></div>
        <div class="h-4 bg-gray-200 rounded w-3/4"></div>
      </div>
  
      <!-- Error -->
      <div v-else-if="store.error" class="p-4 bg-red-50 border border-red-200 rounded">
        <p class="text-red-600 font-medium">{{ store.error }}</p>
      </div>
  
      <!-- Datos -->
      <div v-else-if="store.bestStock" >
        <div class="flex items-center gap-3">
          <div class="flex-1">
            <h3 class="text-lg font-semibold text-gray-900">
              {{ store.bestStock.ticker }}
              <span class="text-gray-600 ml-2">{{ store.bestStock.company }}</span>
            </h3>
            <p class="text-sm text-gray-500">{{ store.bestStock.brokerage }}</p>
          </div>
          <span class="px-3 py-1 text-sm rounded-full bg-blue-100 text-blue-800">
            {{ store.bestStock.rating_to }}
          </span>
        </div>
  
        <div class="grid grid-cols-2 gap-4">
          <div class="p-3 bg-gray-50 rounded">
            <p class="text-sm text-gray-500 mb-1">Target Price</p>
            <p class="font-medium">
              <span class="text-gray-700">{{ store.bestStock.target_from }}</span>
              <span class="mx-2 text-gray-400">→</span>
              <span class="text-green-600">{{ store.bestStock.target_to }}</span>
            </p>
          </div>
          
          <div class="p-3 bg-gray-50 rounded">
            <p class="text-sm text-black-500 mb-1">Rating Change</p>
            <div class="flex items-center gap-2">
              <span class="px-2 py-1 text-sm rounded-full bg-red-100 text-red-800">
                {{ store.bestStock.rating_from }}
              </span>
              <span class="text-gray-400">→</span>
              <span class="px-2 py-1 text-sm rounded-full bg-green-100 text-green-800">
                {{ store.bestStock.rating_to }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>