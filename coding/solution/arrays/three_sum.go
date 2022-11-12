package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}

	tripletMap := make(map[[3]int][]int)

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				triplet := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplet[:])
				if (nums[i] + nums[j] + nums[k]) == 0 {
					tripletMap[triplet] = []int{nums[i], nums[j], nums[k]}
				}
			}
		}
	}

	// fmt.Println(result)

	for _, triplet := range tripletMap {
		result = append(result, triplet)
	}

	return result
}

func threeSumOptimal(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target, left, right := -nums[i], i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]

			if target == sum {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				if left < right && nums[left] == nums[left-1] {
					left++
				}
				if left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if target > sum {
				left++
			} else if sum > target {
				right--
			}
		}

	}

	return result
}

func main() {
	fmt.Println(threeSumOptimal([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSumOptimal([]int{0, 0, 0, 0}))
	fmt.Println(threeSumOptimal([]int{-2, 0, 0, 2, 2}))
}
