package main

import "fmt"

func containsDuplicate(nums []int) bool {
    numbers := make(map[int]bool)

	for _, v := range nums {
		if numbers[v] {
			return true
		}

		numbers[v] = true
	}

	return false
}

//
func main() {
	fmt.Println(containsDuplicate([]int{4, 9, 5}))
	fmt.Println(containsDuplicate([]int{1, 1, 2}))
	fmt.Println(containsDuplicate([]int{1, 1}))
}
