package main

import (
	"fmt"
	"sort"
)

//ugly, but faster than intersectionOfTwoArrays1 and uses less memory
func intersectionOfTwoArrays0(nums1 []int, nums2 []int) []int {
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

//counts frequency of elems in first map
// then iterates through nums2 and if elem was found
//pushes the elem to result array and dicrement frequency in Map
// if nums2 < nums1 they must be swaped
func intersectionOfTwoArrays1(nums1 []int, nums2 []int) []int {
	if len(nums2) < len(nums1) {
		intersectionOfTwoArrays1(nums2, nums1)
	}

	intersection := []int{}

	nums1Frequency := make(map[int]int)

	for _, v := range nums1 {
		if _, ok := nums1Frequency[v]; ok {
			nums1Frequency[v]++
		} else {
			nums1Frequency[v] = 1
		}
	}

	for _, v := range nums2 {
		if nums1Frequency[v] > 0 {
			intersection = append(intersection, v)
			nums1Frequency[v]--
		}
	}

	return intersection
}

//The approch is the MOST EFFICIENT
//uses 2 Pointers approach on sorted input arrays
// compares elems from arrays and
//dicrement/increment pointers based on comparison result
func intersectionOfTwoArrays2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	intersection := []int{}
	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			intersection = append(intersection, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return intersection
}

func main() {
	fmt.Println(intersectionOfTwoArrays2([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersectionOfTwoArrays2([]int{1, 2}, []int{1, 1}))
	fmt.Println(intersectionOfTwoArrays2([]int{1, 1}, []int{1, 2}))
}
