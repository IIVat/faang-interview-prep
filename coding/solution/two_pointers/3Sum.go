package main

import (
	"fmt"
	"sort"
)

func findSumOfThreeOpt(nums []int, target int) bool {
	sort.Ints(nums)
	fmt.Printf("%v\n", nums)

	low, high := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		low = i + 1
		for low < high {
			sum3 := nums[i] + nums[low] + nums[high]
			if sum3 == target {
				return true
			} else if sum3 > target {
				high--
			} else {
				low++
			}
		}

	}
	return false
}

func main() {
	println(findSumOfThreeOpt([]int{3, 7, 1, 2, 8, 4, 5}, 20))
	println(findSumOfThreeOpt([]int{-1, 2, 1, -4, 5, -3}, -8))
	println(findSumOfThreeOpt([]int{-1, 1, 0}, -1))
}
