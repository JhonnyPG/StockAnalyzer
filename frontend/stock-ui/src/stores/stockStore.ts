import { defineStore } from 'pinia'
import axios from 'axios'

interface Stock {
    ticker: string
    company: string
    brokerage: string
    action: string
    target_from: string
    target_to: string
    rating_from: string
    rating_to: string
    time: string
}

export const useStockStore = defineStore('stocks', {
    state: () => ({
        stocks: [] as Stock[],
        loading: false,
        error: null as string | null
    }),
    actions: {
        async fetchStocks() {
            this.loading = true
            this.error = null
            try {
                const response = await axios.get('/api/stocks')
                this.stocks = response.data
            } catch (error) {
                this.error = error instanceof Error ? error.message : 'Failed to fetch stocks'
            } finally {
                this.loading = false
            }
        }
    }
})