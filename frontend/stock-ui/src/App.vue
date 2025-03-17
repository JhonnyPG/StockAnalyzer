<script setup lang="ts">
import StockList  from './components/StockList.vue'
import Recommendation from './components/Recommendation.vue'
</script>

<template>
  <nav class="bg-gray-800 p-4">
    <div class="container mx-auto">
      <router-link to="/" class="text-white mr-4">Inicio</router-link>
      <router-link to="/recommendation" class="text-white">Recomendación</router-link>
    </div>
  </nav>
  <router-view />
</template>
