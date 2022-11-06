package main

import "fmt"

func searchInsert(nums []int, target int) int {
	mid := -1
	begin, finish := 0, len(nums) // b = 0 ;  f = 4

	for begin < finish { // 0 < 4 => 0 < 1
		mid = (begin + finish) / 2 // (0 + 4) / 2 = 2; (0 + 1) / 2 = 0;

		if nums[mid] == target { // 5 != 4
			return mid
		}

		if nums[mid] < target { // 5 > 4
			begin = mid + 1
		} else {
			finish = mid // f= 2 - 1  = 1
		}
	}

	if nums[mid] < target {
		mid++
	}

	return mid
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 7, 8, 9, 11, 15}, 10))
	fmt.Println(searchInsert([]int{1, 3, 5, 7}, 4)) //Wrong
	fmt.Println(searchInsert([]int{1, 3, 5, 7}, 8))
	fmt.Println(searchInsert([]int{1, 3, 5, 7}, 0))
	fmt.Println(searchInsert([]int{1}, 2))
	fmt.Println(searchInsert([]int{1}, 0))
	fmt.Println(searchInsert([]int{1}, 1))
}
