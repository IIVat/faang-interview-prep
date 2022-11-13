package main

import "fmt"

/*
	Explanation:

*/

func maxSubArray(nums []int) int {
	maxSum := nums[0]

	currSum := 0

	for i := 0; i < len(nums); i++ {

		if currSum <= -nums[i] {
			currSum = 0
		} else {
			currSum = currSum + nums[i]
		}

		if currSum == 0 {
			if nums[i] > maxSum {
				maxSum = nums[i]
			}
		} else if currSum > maxSum {
			maxSum = currSum
		}

	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-1, -2, -3}))
	fmt.Println(maxSubArray([]int{-3, -2, -1}))
	fmt.Println(maxSubArray([]int{-1, 0}))
	m := []int{-3, -2, -1}

	fmt.Println(m[1:])
}
