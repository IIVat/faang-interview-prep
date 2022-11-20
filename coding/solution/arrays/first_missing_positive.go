package main

import (
	"fmt"
)

//applied
//https://emre.me/coding-patterns/cyclic-sort/

func firstMissingPositive(nums []int) int {

	i := 0
	for i < len(nums) {
		if nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			Swap(nums, i, nums[i]-1)
		} else {
			i++
		}
	}

	for i, v := range nums {
		if i != v-1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
func Swap(nums []int, i, j int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 3}))
	fmt.Println(firstMissingPositive([]int{-1, 1, 2, 3}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{6, 4, 3, 2, 8, 5}))
}
