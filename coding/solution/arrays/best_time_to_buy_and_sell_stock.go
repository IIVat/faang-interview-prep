package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	min := math.MaxInt
	profit := 0

	for _, v := range prices {

		if v < min {
			min = v
		}

		profit = int(math.Max(float64(v-min), float64(profit)))
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
