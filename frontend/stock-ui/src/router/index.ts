import { createRouter, createWebHistory } from 'vue-router'
import StockTable from '@/components/StockTable.vue'  // Componente existente
import Recommendations from '@/components/Recommendation.vue'  // Componente existente

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/stocks'
    },
    {
      path: '/stocks',
      name: 'Stocks',
      component: StockTable // Usando el componente directamente
    },
    {
      path: '/recommendations',
      name: 'Recommendations',
      component: Recommendations // Usando el componente directamente
    }
  ]
})

export default router