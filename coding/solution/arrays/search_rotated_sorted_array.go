package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// until both comes not equal
	for left <= right {
		mid := (left + right) / 2 //Calculating mid point
		//checking if mid is target then return  index
		if target == nums[mid] {
			return mid
		}
		//checking first half array is sorted or not
		if nums[mid] >= nums[left] {
			//checking target is exist in first half or not
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//checking target is exist in first half or not
			if target <= nums[right] && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 10, 1, 2}, 9))
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	// fmt.Println(search([]int{0, 1, 2}, 2))
	// fmt.Println(search([]int{3, 5, 1}, 3))
	// fmt.Println(search([]int{6, 0, 1, 2, 3, 4, 5}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 5))

	// fmt.Println(search([]int{4, 5, 6, 7}, 6))

}
