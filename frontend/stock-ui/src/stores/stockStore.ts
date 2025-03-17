import { defineStore } from 'pinia'
import axios from 'axios'

interface stock {
  Ticker: string
  Company: string
  Brockerage: string
  TargetTo: string
  TargetFrom: string
  RatingFrom: string
  RatingTo: string
  Time: string

}

export const useStockStore = defineStore('stock', {
  state: () => ({
    stocks: [] as stock[],
    recommendation: {} as stock[]
  }),
  actions: {
    async fetchStocks() {
      const response = await axios.get('http//localhost:8080/stocks')
      this.stocks = response.data
    },
    async fetchRecommendation() {
      const response = await axios.get('http//localhost:8080/recommendation')
      this.recommendation = response.data
    }
  }


})
