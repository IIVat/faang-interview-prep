package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	output := [][]int{intervals[0]}

	// fmt.Printf("%v", intervals)

	for i := 1; i < len(intervals); i++ {
		i1, i2 := output[len(output)-1], intervals[i]

		if i1[1] >= i2[0] {
			output[len(output)-1] = []int{i1[0], max(i1[1], i2[1])}
		} else {
			output = append(output[:len(output)-1], [][]int{i1, i2}...)
		}
	}
	return output
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}

	return i2
}

func main() {
	fmt.Printf("%v", merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
