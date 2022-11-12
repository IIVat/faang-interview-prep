package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxVolume := 0
	i := 0
	j := len(height) - 1

	for i < j {
		if height[i] < height[j] {
			maxVolume = Max(maxVolume, height[i]*(j-i))
			i++
		} else {
			maxVolume = Max(maxVolume, height[j]*(j-i))
			j--
		}

	}

	return maxVolume
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2, 1, 2}))
}
