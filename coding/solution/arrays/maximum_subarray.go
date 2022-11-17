package main

import "fmt"

/*
	Explanation:

	Given array: [-2, 1, -3, 4, -1, 2, 1, -5, 4]

	i        nums[i]         sum                 max
	0         -2             -2                  -2
	1          1          -2 < 0 = 1           -2 < 1 = 1
    2         -3          1 > 0 = 1 - 3 = -2    1 > -2 = 1
	3          4          -2 < 0 = 4            1 < 4 = 4
	4         -1 		  4 >0 = 4 - 1= 3	    4 > 3 = 4
	5          2          3 >0 = 3 + 2 = 5      5 > 4 = 5
	6          1          5 > 0 = 5 + 1 = 6     5 < 6 = 6
	7         -5          6 > 0 = 6 - 5 = 1     6 > 1 = 6
	8         4           1 > 0 = 1 + 4 = 5     6 > 5 = 6


	Result: 6
*/

func maxSubArrayConcise(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]

	currSum := 0

	for i := 0; i < len(nums); i++ {

		if currSum <= -nums[i] {
			currSum = 0
		} else {
			currSum = currSum + nums[i]
		}

		if currSum == 0 {
			if nums[i] > maxSum {
				maxSum = nums[i]
			}
		} else if currSum > maxSum {
			maxSum = currSum
		}

	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-1, -2, -3}))
	fmt.Println(maxSubArray([]int{-3, -2, -1}))
	fmt.Println(maxSubArray([]int{-1, 0}))
	m := []int{-3, -2, -1}

	fmt.Println(m[1:])
}
