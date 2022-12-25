package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

//mein
func hasCycleMein(head *ListNode) bool {
	slow := head

	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	fast := head.Next.Next

	for fast != nil {
		if fast == slow {
			return true
		}

		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
