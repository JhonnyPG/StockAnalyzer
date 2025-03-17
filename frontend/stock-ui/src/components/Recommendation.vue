<template>
    <div class = "container mx-auto p-4">
        <h2 class = "text-2x1 font-bold mb-4">Recommendation</h2>
        <div v-if = "Stock.recommendation.targetTo" class ="bg-blue-50 p-6 rounded-lg">
            <h3 class = "text-x1 font-semibold"> {{ recommendation.ticker }}</h3>
            <p class = "text-gray-600">{{ recommendation.company }}</p>
            <div class="mt-4">
        <span class="font-medium">Potencial:</span>
        <span class="ml-2 text-green-600">
          {{ recommendation.target_to }} ({{ calculateDifference() }})
        </span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStockStore } from '@/stores/stockStore'
import { Stock } from '@/types/Stocks'
import { onMounted } from 'vue'

const store = useStockStore()

onMounted(() => {
    store.fetchRecommendation()
})

const calculateDifference = () => {
    const to = parseFloat(store.recommendation.target_to.replace('$', ''))
    const from = parseFloat(store.recommendation.target_from.replace('$', ''))
    return `+${(to - from).toFixed(2)}`
}
</script>