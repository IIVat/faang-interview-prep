package main

import (
	"fmt"
	"sort"
)

func intersectionOfTwoArrays(nums1 []int, nums2 []int) []int {
	intersection := []int{}
	visited := make(map[int]bool)

	minArray := []int{}
	maxArray := []int{}

	if len(nums1) <= len(nums2) {
		minArray = nums1
		maxArray = nums2
	} else {
		minArray = nums2
		maxArray = nums1
	}

	sort.Ints(minArray)
	sort.Ints(maxArray)

	minArrayLen := len(minArray)
	maxArrayLen := len(maxArray)

	for i := 0; i < minArrayLen; i++ {
		insert := -1
		currElem := minArray[i]
		for j := 0; j < maxArrayLen; j++ {
			if currElem == maxArray[j] {
				if visited[j] == false {
					visited[j] = true
					insert = maxArray[j]
					break
				}
			}
		}

		if insert != -1 {
			intersection = append(intersection, insert)
		}
	}

	return intersection
}

func main() {
	fmt.Println(intersectionOfTwoArrays([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersectionOfTwoArrays([]int{1, 2}, []int{1, 1}))
	fmt.Println(intersectionOfTwoArrays([]int{1, 1}, []int{1, 2}))
}
