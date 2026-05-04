package arrays

func (a *array) MaxProfitOptimized(prices []int) int {
  buyDay, maxProfit := prices[0], 0
  for i := 1 ; i<len(prices) ; i++ {
    currentPrice := prices[i]

    if currentPrice < buyDay {
        buyDay = currentPrice
    } else if (currentPrice-buyDay > maxProfit){
        maxProfit = currentPrice-buyDay
    }
  }
   return maxProfit
}