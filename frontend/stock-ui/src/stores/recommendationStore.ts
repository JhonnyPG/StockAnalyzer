import { defineStore } from 'pinia'
import axios from 'axios'

interface BestStock {
    ticker: string
    company: string
    brokerage: string
    action: string
    rating_from: string
    rating_to: string
    target_from: string
    target_to: string
    time: string
}

export const useRecommendationStore = defineStore('recommendations', {
    state: () => ({
        bestStock: null as BestStock | null,
        loading: false,
        error: null as string | null
    }),
    actions: {
        async fetchBestStock() {
            this.loading = true
            this.error = null
            try {
                const response = await axios.get('/api/recommendations')
                this.bestStock = response.data.best_stock
            } catch (error) {
                this.error = error instanceof Error ? error.message : 'Failed to fetch recommendation'
            } finally {
                this.loading = false
            }
        }
    }
})