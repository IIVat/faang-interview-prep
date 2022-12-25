package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverse(slow)

	for slow != nil {
		if slow.Val != head.Val {
			return false
		}

		slow = slow.Next
		head = head.Next
	}

	return true
}

func reverse(half *ListNode) *ListNode {
	reverse := new(ListNode)
	reverse = nil
	for half != nil {
		next := half.Next
		half.Next = reverse
		reverse = half
		half = next
	}
	return reverse
}

func main() {

}
