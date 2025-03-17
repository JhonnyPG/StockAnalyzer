import { createRouter, createWebHistory } from 'vue-router'
import StockList from '../components/StockList.vue'
import Recommendation from '../components/Recommendation.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: StockList },
    { path: '/recommendation', component: Recommendation }
  ],
});

export default router;