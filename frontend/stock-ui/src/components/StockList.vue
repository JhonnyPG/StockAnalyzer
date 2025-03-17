<template>
    <div class = "container" mix-auto p-4>
        <h1 class = "text-3x1 font-bold mb-6">Stock List</h1>

        <div class = "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div 
                v-for = "stock in stocks"
                :key = "stock.ticker"
                class = "bg-white p-4 rounded-lg shadow-md" >

                <h2 class = "text-x1 font-semibold">{{ stock.ticker }}</h2>
                <p class = "text-gray-600">{{ stock.company }}</p>
                <div class = "mt-2">
                    <span class = "text-green-600">{{ stock.targetTo }}</span>
                    <span class = "mx-2">→</span>
                    <span class = "text-red-600">{{ stock.targetFrom }}</span>
                </div>
            </div>
        </div>

    </div>
</template>

<script setup lang="ts">

import { useStockStore } from '@/stores/stockStore';
import { onMounted } from 'vue';

const store = useStockStore();

onMounted(() => {
    store.fetchStocks();
})
</script>
