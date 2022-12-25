package main

import (
	"fmt"
	"sort"
	"strconv"
)

func employeeFreeTime(schedule [][]Interval) []Interval {
	intervals := make([]Interval, 0)
	flattenedIntervals := flattenIntervals(schedule)
	sortedIntervals := sortIntervalsByStart(flattenedIntervals)
	maxEnd := sortedIntervals[0].end

	for i := 1; i < len(sortedIntervals); i++ {
		a := sortedIntervals[i-1]
		b := sortedIntervals[i]
		if !isOverlapingIntervals(a, b) {
			maxEnd = max(maxEnd, a.end)

			intervals = append(intervals, getDiff(Interval{a.start, maxEnd}, b))
		}
	}
	return intervals
}

func flattenIntervals(schedule [][]Interval) []Interval {
	intervals := make([]Interval, 0)
	for _, invs := range schedule {
		for _, v := range invs {
			intervals = append(intervals, v)
		}
	}
	return intervals
}

func sortIntervalsByStart(intervals []Interval) []Interval {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].start < intervals[j].start {
			return true
		} else if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		} else {

			return false
		}
	})

	return intervals
}

func isOverlapingIntervals(a, b Interval) bool {
	return (min(a.end, b.end) - max(a.start, b.start)) >= 0
}

func getDiff(a, b Interval) Interval {
	return Interval{a.end, b.start}
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
	inputs1 := [][]Interval{
		{Interval{1, 2}, Interval{5, 6}, Interval{1, 3}, Interval{4, 10}},
		{Interval{1, 3}, Interval{6, 7}, Interval{2, 4}, Interval{2, 5}, Interval{11, 12}},
		{Interval{2, 3}, Interval{7, 9}, Interval{1, 4}, Interval{6, 7}},
	}
	// {{Interval{3, 5}, Interval{8, 10}}, {Interval{4, 6}, Interval{9, 12}}, {Interval{5, 6}, Interval{8, 10}}},
	// {{Interval{1, 3}, Interval{6, 9}, Interval{10, 11}}, {Interval{3, 4}, Interval{7, 12}}, {Interval{1, 3}, Interval{7, 10}}, {Interval{1, 4}}, {Interval{7, 10}, Interval{11, 12}}},
	// {{Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}}, {Interval{2, 3}, Interval{4, 5}, Interval{6, 8}}},
	// {{Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}, {Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}, {Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}, {Interval{1, 2}, Interval{3, 4}, Interval{5, 6}, Interval{7, 8}, Interval{9, 10}, Interval{11, 12}}},

	inputs2 := [][]Interval{
		{Interval{2, 3}, Interval{7, 9}}, {Interval{1, 4}, Interval{6, 7}}}

	fmt.Printf("%v", employeeFreeTime(inputs1))
	println("")
	fmt.Printf("%v", employeeFreeTime(inputs2))
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
