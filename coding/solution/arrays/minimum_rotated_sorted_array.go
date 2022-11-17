package main

import (
	"fmt"
	"math"
)

/*
  Solution
   https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/solution/
*/

func findMinOptimal(nums []int) int {
	// If the list has just one element then return that element.
	if len(nums) == 1 {
		return nums[0]
	}

	// initializing left and right pointers.
	left, right := 0, len(nums)-1

	// if the last element is greater than the first element then there is no
	// rotation.
	// e.g. 1 < 2 < 3 < 4 < 5 < 7. Already sorted array.
	// Hence the smallest element is first element. A[0]
	if nums[right] > nums[0] {
		return nums[0]
	}

	// Binary search way
	for right >= left {
		mid := left + (right-left)/2

		// if the mid element is greater than its next element then mid+1 element is the
		// smallest
		// This point would be the point of change. From higher to lower value.
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		// if the mid element is lesser than its previous element then mid element is
		// the smallest
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		// if the mid elements value is greater than the 0th element this means
		// the least value is still somewhere to the right as we are still dealing with
		// elements
		// greater than nums[0]
		if nums[mid] > nums[0] {
			left = mid + 1
		} else {
			// if nums[0] is greater than the mid value then this means the smallest value
			// is somewhere to
			// the left
			right = mid - 1
		}
	}

	return math.MaxInt
}

//works perfectly but the code is ugly
func findMin(nums []int) int {
	mid := len(nums) / 2
	var min int
	if mid == 0 {
		return nums[mid]
	}

	if mid == len(nums)-1 {
		return Min(nums[mid], nums[mid-1])
	}

	if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
		return Min(nums[mid-1], nums[mid+1])
	}

	if nums[mid-1] > nums[mid] && nums[mid] < nums[mid+1] {
		return nums[mid]
	}

	if nums[0] > nums[len(nums)-1] {
		if nums[mid] < nums[len(nums)-1] {
			min = findMin(nums[:mid])
		} else {
			min = findMin(nums[mid+1:])
		}
	}

	if nums[0] < nums[len(nums)-1] {
		if nums[mid] < nums[0] {
			min = findMin(nums[mid+1:])
		} else {
			min = findMin(nums[:mid])
		}
	}

	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	// fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	// fmt.Println(findMin([]int{1, 2, 4}))
	// fmt.Println(findMin([]int{1}))
	fmt.Println(findMinOptimal([]int{4, 0, 2}))
	fmt.Println(findMinOptimal([]int{11, 13}))
	fmt.Println(findMinOptimal([]int{13, 11}))
	// fmt.Println(findMin([]int{2, 3, 4, 5, 1}))
	fmt.Println(findMinOptimal([]int{3, 4, 5, 6, 7, 1, 2}))
	fmt.Println(findMinOptimal([]int{5, 1, 2, 3, 4}))

	// m := []int{3, 4, 5, 1, 2}
	// fmt.Println(m[:3])
	// fmt.Println(m[3:])
}
