package main

import "fmt"

func removeDuplicates(nums []int) int {
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[count-1] != nums[i] {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(removeDuplicates([]int{-1, 1, 1, 1, 2, 4, 4, 4, 9}))
	fmt.Println(removeDuplicates([]int{1}))
	fmt.Println(removeDuplicates([]int{1, 1, 1}))
}
