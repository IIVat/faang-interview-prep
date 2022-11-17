package main

import "fmt"

func nextPermutation(nums []int) []int {

	i, j := len(nums)-1, len(nums)-1

	//find point when numbers starting to decrease
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}

	i = i - 1

	//find first number which more than number under decreasing point
	for i > -1 && j > i {
		if nums[i] < nums[j] {
			Swap(nums, i, j)
			break
		}
		j--
	}

	fmt.Println(i)

	//reverse numbers which are placed after decreasing point
	Reverse(nums, i+1)

	return nums
}

func Reverse(nums []int, start int) {
	for i, j := start, len(nums)-1; i < j; i++ {
		Swap(nums, i, j)
		j--
	}
}

func Swap(nums []int, i, j int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}

func main() {
	fmt.Println(nextPermutation([]int{1, 9, 9, 8, 5}))
	fmt.Println(nextPermutation([]int{1, 3, 2}))
	fmt.Println(nextPermutation([]int{1, 2, 3}))
	fmt.Println(nextPermutation([]int{3, 2, 1}))
	fmt.Println(nextPermutation([]int{2, 3, 1}))
	fmt.Println(nextPermutation([]int{4, 2, 0, 2, 3, 2, 0}))

}
