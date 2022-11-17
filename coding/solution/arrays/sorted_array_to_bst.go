package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	root := &TreeNode{Val: nums[i]}
	root.Left = sortedArrayToBST(nums[:i])
	root.Right = sortedArrayToBST(nums[i+1:])
	return root
}

//beat 100% cpu&mem:
func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	root := &TreeNode{Val: nums[i]}
	if i != 0 {
		root.Left = sortedArrayToBST(nums[:i])
	}
	if i != len(nums)-1 {
		root.Right = sortedArrayToBST(nums[i+1:])
	}
	return root
}

func main() {
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}))
}
