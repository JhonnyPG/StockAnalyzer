<script setup lang="ts">
import { useStockStore } from '@/stores/stockStore'
import { onMounted } from 'vue'

const store = useStockStore()

onMounted(() => {
  store.fetchStocks()
})
</script>

<template>
  <div class="overflow-x-auto">
    <div v-if="store.loading" class="text-center py-4">Loading...</div>
    <div v-else-if="store.error" class="text-red-500 p-4">{{ store.error }}</div>
    <table v-else class="min-w-full divide-y divide-gray-200">
      <thead class="bg-gray-50">
        <tr>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Ticker</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Company</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Action</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Brokerage</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Target Price</th>
          <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Rating From</th>
          <th class = "px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase" > Rating To </th>
        </tr>
      </thead>
      <tbody class="bg-yellow-100 divide-y divide-gray-200">
        <tr v-for="stock in store.stocks" :key="stock.ticker">
          <td class="px-6 py-4 whitespace-nowrap font-medium">{{ stock.ticker }}</td>
          <td class="px-6 py-4 whitespace-nowrap">{{ stock.company }}</td>
          <td class="px-6 py-4 whitespace-nowrap">{{ stock.action }}</td>
          <td class="px-6 py-4 whitespace-nowrap">{{ stock.brokerage }}</td>
          <td class="px-6 py-4 whitespace-nowrap">
            {{ stock.target_from }} → {{ stock.target_to }}
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            <span class="px-2 py-1 text-sm rounded-full" 
                  :class="{
                    'bg-green-100 text-green-800': stock.rating_from === 'Buy',
                    'bg-yellow-100 text-orange-800': stock.rating_from === 'Hold',
                    'bg-red-100 text-red-800': stock.rating_from === 'Sell'
                  }">
              {{ stock.rating_from }}
            </span>
          </td>
          <td class="px-6 py-4 whitespace-nowrap">
            <span class="px-2 py-1 text-sm rounded-full" 
                  :class="{
                    'bg-green-100 text-green-800': stock.rating_to === 'Buy',
                    'bg-yellow-100 text-orange-800': stock.rating_to === 'Hold',
                    'bg-red-100 text-red-800': stock.rating_to === 'Sell'
                  }">
              {{ stock.rating_to }}
            </span>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>