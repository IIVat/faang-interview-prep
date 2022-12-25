package main

import "fmt"

//should try binary search
func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0)
	result = append(result, merge(intervals[0], newInterval)...)

	if len(intervals) == 1 {
		return result
	}

	if len(result) == 2 && result[1][0] == intervals[0][0] && result[1][1] == intervals[0][1] {
		return append(result, intervals[1:]...)
	}

	fmt.Printf("%v\n", result)

	for i := 1; i < len(intervals); i++ {
		fmt.Printf("%v\n", result[:len(result)-1])
		result = append(result[:len(result)-1], merge(result[len(result)-1], intervals[i])...)

		a := result[len(result)-1]
		b := intervals[i]
		if a[0] == b[0] && a[1] == b[1] {
			return append(result, intervals[i+1:]...)
		}
	}

	return result

}

func merge(i1, i2 []int) [][]int {
	first, second := sort2Intervals(i1, i2)

	if first[1] < second[0] {
		return [][]int{first, second}
	}

	return [][]int{{first[0], max(first[1], second[1])}}

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

func sort2Intervals(i1, i2 []int) ([]int, []int) {
	if i1[0] <= i2[0] {
		return i1, i2
	}
	return i2, i1
}

//Binary Search

func insert2(intervals [][]int, newInterval []int) [][]int {
	intervals = insertInterval(intervals, newInterval)
	// fmt.Printf("%v\n", intervals)

	result := [][]int{}
	for i := 0; i < len(intervals); i++ {
		currInterval := intervals[i]

		for i < len(intervals) && isOverlapingIntervals(currInterval, intervals[i]) {
			// fmt.Printf("%v\n", currInterval)
			// fmt.Printf("%v\n", intervals[i])
			// println(i)
			currInterval = mergeIntervals(currInterval, intervals[i])
			// fmt.Printf("%v\n", currInterval)
			// println("------")
			i++
		}

		result = append(result, currInterval)
		// fmt.Printf("%v\n", result)
		i--
	}

	return result
}

func isOverlapingIntervals(a, b []int) bool {
	return min(a[1], b[1])-max(a[0], b[0]) >= 0
}

func mergeIntervals(a, b []int) []int {
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", []int{min(a[0], b[0]), max(a[1], b[1])})
	return []int{min(a[0], b[0]), max(a[1], b[1])}
}

func binarySearch(arr [][]int, target int) int {
	start := 0
	end := len(arr) - 1

	insertPos := len(arr)

	for start <= end {
		mid := (start + end) / 2

		if arr[mid][0] > target {
			insertPos = mid
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	// println(insertPos)

	return insertPos
}

func insertInterval(intervals [][]int, newInterval []int) [][]int {
	insertionPos := binarySearch(intervals, newInterval[0])

	if insertionPos == 0 {
		return append([][]int{newInterval}, intervals...)
	} else if insertionPos != len(intervals) {
		intervals = append(intervals[:insertionPos+1], intervals[insertionPos:]...)
		intervals[insertionPos] = newInterval

		return intervals
	}

	return append(intervals, newInterval)

}

func main() {
	fmt.Printf("%v\n", insert2([][]int{{1, 6}, {8, 9}, {10, 15}, {16, 18}}, []int{9, 10}))
	// fmt.Println(binarySearch([][]int{{1, 3}, {6, 9}, {9, 10}}, 11))
}
