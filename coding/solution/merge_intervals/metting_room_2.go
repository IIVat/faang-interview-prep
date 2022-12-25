package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findSets(intervals []Interval) int {
	rooms := []int{intervals[0].end}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	for i := 1; i < len(intervals); i++ {
		rooms = addToRoom(intervals[i], rooms)
	}

	return len(rooms)
}

func addToRoom(interval Interval, rooms []int) []int {
	lenR := len(rooms)

	for i := 0; i < lenR; i++ {
		a, b := rooms[i], interval
		if a <= b.start {
			return rooms

		}

	}

	return append(rooms, interval.end)

}

//clever solution!!!

func findSetsBetter(intervals []Interval) int {
	start := []int{}
	end := []int{}
	for _, v := range intervals {
		start = append(start, v.start)
		end = append(end, v.end)
	}

	sortIntervals(start)
	sortIntervals(end)

	res, count := 0, 0
	i, j := 0, 0

	for i < len(intervals) {
		if start[i] < end[j] {
			i += 1
			count += 1
		} else {
			j += 1
			count -= 1
		}

		res = max(count, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func sortIntervals(intervals []int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i] < intervals[j]
	})
}

func main() {
	inputs1 := []Interval{{2, 8}, {3, 4}, {3, 9}, {5, 11}, {8, 20}, {11, 15}}

	fmt.Printf("%v\n", findSetsBetter(inputs1))
}

type Interval struct {
	start int
	end   int
}

func (i *Interval) IntervalInit(start int, end int) {
	i.start = start
	i.end = end
}

func (i *Interval) str() string {
	out := "(" + strconv.Itoa(i.start) + ", " + strconv.Itoa(i.end) + ")"
	return out
}
