package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}

func removeElement2(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	return n
}

func main() {
	fmt.Println(removeElement([]int{3, 3, 2, 3, 4, 6}, 3))
	fmt.Println(removeElement2([]int{3, 3, 2, 3, 4, 6}, 3))
}
