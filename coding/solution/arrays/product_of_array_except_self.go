package main

import "fmt"

/*
 Example:

   Left:     1   |   2   |  2*3  | 2*3*4|
   Nums:     2   |   3   |   4   |   5  |
   Right   3*4*5 |  4*5  |   5   |   1  |

   Result:  Left*Right

   Explanation: https://leetcode.com/problems/product-of-array-except-self/discuss/65622/Simple-Java-solution-in-O(n)-without-extra-space

*/

func productExceptSelfOptimal(nums []int) []int {

	left := 1
	result := []int{left}

	for i := 1; i < len(nums); i++ {
		result = append(result, result[i-1]*nums[i-1])
	}

	right := 1

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * right
		right = right * nums[i]
	}

	return result
}

func productExceptSelf(nums []int) []int {
	var result []int
	forwardProduct := []int{nums[0]}
	backwardProduct := []int{nums[len(nums)-1]}

	for i := 1; i < len(nums); i++ {
		forwardProduct = append(forwardProduct, forwardProduct[i-1]*nums[i])
		backwardProduct = append([]int{backwardProduct[0] * nums[len(nums)-i-1]}, backwardProduct...)
	}

	for k := 0; k < len(nums); k++ {
		if k-1 < 0 {
			result = append(result, backwardProduct[k+1])
		} else if k+1 > len(nums)-1 {
			result = append(result, forwardProduct[k-1])
		} else {
			result = append(result, backwardProduct[k+1]*forwardProduct[k-1])
		}

	}

	return result
}

func main() {
	fmt.Println(productExceptSelfOptimal([]int{2, 3, 4, 5}))
	fmt.Println(productExceptSelfOptimal([]int{-2, 2, 3, 5, -3}))
}
