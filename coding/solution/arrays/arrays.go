package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				res = []int{i, j}
			}
		}
	}
	return res
}

func TwoSumOptimal(nums []int, target int) []int {
	m := make(map[int]int)
	for currIdx, currNum := range nums {
		if requiredIdx, isPresent := m[target-currNum]; isPresent {
			fmt.Println(requiredIdx)
			return []int{requiredIdx, currIdx}
		}
		m[currNum] = currIdx
	}
	return []int{}
}

func main() {
	fmt.Println(TwoSum([]int{5, 3, 9, 1}, 6))
	fmt.Println(TwoSumOptimal([]int{5, 3, 9, 1}, 6))

}
