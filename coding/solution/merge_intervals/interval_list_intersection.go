package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	maxList, minList := getMaxLenList(firstList, secondList)

	for _, v := range minList {
		maxList = insertInterval(maxList, v)
	}
	result := [][]int{}

	// fmt.Printf("%v\n", maxList)

	for i := 1; i < len(maxList); i++ {
		var currInterval []int
		t := maxList[i-1]

		for i < len(maxList) && isOverlapingIntervals(t, maxList[i]) {
			currInterval = getIntersection(maxList[i], t)
			if t[1] < maxList[i][1] {
				t = maxList[i]
			}
			result = append(result, currInterval)
			// fmt.Printf("%v\n", currInterval)
			i++
		}

	}

	return result
}

func getMaxLenList(firstList [][]int, secondList [][]int) ([][]int, [][]int) {
	if len(firstList) > len(secondList) {
		return firstList, secondList
	}

	return secondList, firstList
}

func isOverlapingIntervals(a, b []int) bool {
	return (min(a[1], b[1]) - max(a[0], b[0])) >= 0
}

func getIntersection(a, b []int) []int {
	return []int{max(a[0], b[0]), min(a[1], b[1])}
}

//binary search
func searchInsertPosition(list [][]int, a []int) int {
	ans := len(list)
	start := 0
	end := len(list) - 1

	for start <= end {
		mid := (start + end) / 2

		if a[0] < list[mid][0] {
			ans = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return ans
}

func insertInterval(list [][]int, a []int) [][]int {
	insertPos := searchInsertPosition(list, a)

	if insertPos == len(list) {
		return append(list, a)
	}

	list = append(list[:insertPos+1], list[insertPos:]...)
	list[insertPos] = a

	return list
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Printf("%v\n", intervalIntersection(
		[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
		[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}))

	fmt.Printf("%v\n", intervalIntersection(
		[][]int{{0, 4}, {5, 7}, {8, 12}, {13, 15}, {16, 18}},
		[][]int{{0, 18}}))

}
