package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	ranges := []string{}
	k := 0
	length := len(nums)
	for k < length {
		start := nums[k]
		end := start
		j := k
		for j < length-1 && nums[j]+1 == nums[j+1] {
			end = nums[j+1]
			j++
			k = j
		}
		arrow := ""
		if k == j {
			k++
		}

		if start < end {
			arrow = "->" + strconv.FormatInt(int64(end), 10)
		}
		rangeStr := strconv.FormatInt(int64(start), 10) + arrow
		ranges = append(ranges, rangeStr)
	}

	return ranges
}

func summaryRangesNotMineSolution(nums []int) []string {
	var ans []string
	for i, j := 0, 0; j < len(nums); j++ {
		i = j
		for j+1 < len(nums) && nums[j]+1 == nums[j+1] {
			j++
		}
		if i == j {
			ans = append(ans, fmt.Sprint(nums[i]))
		} else {
			ans = append(ans, fmt.Sprintf("%v->%v", nums[i], nums[j]))
		}
	}
	return ans
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
