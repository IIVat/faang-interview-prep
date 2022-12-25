package main

func maxProfit(stockPrices []int) int {
	if len(stockPrices) == 0 {
		return 0
	}

	minPrice := stockPrices[0]
	maxProfit := 0

	for i := 1; i < len(stockPrices); i++ {
		if stockPrices[i] < minPrice {
			minPrice = stockPrices[i]
		} else {
			currMax := stockPrices[i] - minPrice
			if stockPrices[i]-minPrice > maxProfit {
				maxProfit = currMax
			}
		}
	}

	return maxProfit
}
