package main

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (i *ListNode) str() string {
	if i == nil {
		return "nil"
	}
	out := strconv.Itoa(i.Val) + " -> " + i.Next.str()
	return out
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var res *ListNode
	length := getLength(head)

	kTimesInLen := 0

	//find how many k in length
	if k >= length {
		kTimesInLen = length
	} else {
		kTimesInLen = (length / k) * k
	}

	count := 0
	//iterate while count less than quantity of k in length
	for count < kTimesInLen {
		innerCount := 0
		//accumulate reversed k-group
		var tmp *ListNode
		//reverse group
		for innerCount < k {
			next := head.Next
			head.Next = tmp
			tmp = head
			head = next
			innerCount++
		}

		//add k-group to result
		res = add(res, tmp)
		count += k
	}

	//add remaining head to result
	res = add(res, head)

	return res
}

func getLength(head *ListNode) int {
	tmp := head
	length := 0
	for tmp != nil {
		tmp = tmp.Next
		length++
	}

	return length
}

func add(head *ListNode, newNode *ListNode) *ListNode {
	if head == nil {
		return &ListNode{newNode.Val, newNode.Next}
	}

	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = newNode

	return head

}

func main() {
	input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	// input := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	// res := add(input, &ListNode{4, &ListNode{3, nil}})
	// println(res.str())
	println(reverseKGroup(input, 3).str())
}
