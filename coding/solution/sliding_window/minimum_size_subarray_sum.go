package main

/*
	1. Init "start" = 0,
	   "end" use array iterration,
	   "minSize" = 0
	   ""
	2. Increase r until r < len(arr)
	  -r++
	3. if currSum >= target => l = l + 2, r++
	4. minLen = l-r+1

*/
import "math"

func minSubArrayLen(target int, nums []int) int {
	start, minLen, currSum := 0, math.MaxInt64, 0

	for end, _ := range nums {
		currSum += nums[end]

		for currSum >= target {
			windowSize := end - start + 1

			if windowSize < minLen {
				minLen = windowSize
			}

			currSum = currSum - nums[start]
			// println(currSum)
			start++

		}
	}

	if minLen == math.MaxInt64 {
		return 0
	}

	return minLen
}

func main() {
	println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	println(minSubArrayLen(4, []int{1, 4, 4}))
	println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	println(minSubArrayLen(10, []int{1, 2, 3, 4}))
	println(minSubArrayLen(5, []int{1, 2, 1, 3}))
}
