package main

import (
	"fmt"
	"math"
)

/*
 Explanation:
 https://leetcode.com/problems/maximum-product-subarray/discuss/1608862/JAVA-or-3-Solutions-or-Detailed-Explanation-Using-Image

*/

func maxProduct(nums []int) int {
	left, right := 1, 1
	ans := nums[0]
	len := len(nums)

	for i := 0; i < len; i++ {
		if left == 0 {
			left = 1
		}

		if right == 0 {
			right = 1
		}

		left = left * nums[i]
		right = right * nums[len-1-i]

		ans = int(math.Max(float64(ans), math.Max(float64(left), float64(right))))

	}

	return ans
}

func maxProduct2(nums []int) int {
	max, min, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		temp := max

		max = Max(Max(max*nums[i], min*nums[i]), nums[i])
		min = Min(Min(min*nums[i], temp*nums[i]), nums[i])

		if max > ans {
			ans = max
		}
	}

	return ans
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-1, 2, -3, 4, -1, 3, -1}))
	// fmt.Println(maxProduct([]int{1}))
	fmt.Println(maxProduct([]int{-1, -2, -3}))
	fmt.Println(maxProduct([]int{-3, -2, -1}))
	fmt.Println(maxProduct([]int{-2, 0, 1}))
	fmt.Println(maxProduct2([]int{3, -1, 4}))

	// fmt.Println(maxProduct([]int{0, 0}))
}
