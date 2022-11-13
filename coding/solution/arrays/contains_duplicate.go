package main

import "fmt"

func containsDuplicate(nums []int) bool {
	numbers := make(map[int]int)

	for _, v := range nums {
		if numbers[v]+1 > 1 {
			return true
		}

		numbers[v]++
	}

	return false
}

//
func main() {
	fmt.Println(containsDuplicate([]int{4, 9, 5}))
	fmt.Println(containsDuplicate([]int{1, 1, 2}))
	fmt.Println(containsDuplicate([]int{1, 1}))
}
