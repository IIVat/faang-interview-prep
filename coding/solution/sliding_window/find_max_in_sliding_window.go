package main

import (
	"fmt"
)

func findMaxSlidingWindow(nums []int, windowSize int) []int {
	result := make([]int, 0)
	// Initializing doubly ended que for storing indices
	window := NewDeque()

	// Return empty list
	if len(nums) == 0 {
		return result
	}

	// If window_size is greater than the array size,
	// set the windowSize to len(nums)
	if windowSize > len(nums) {
		windowSize = len(nums)
	}

	for i := 0; i < windowSize; i++ {
		// For every element, remove the previous smaller elements from window
		for !window.Empty() && nums[i] > nums[window.Back()] {
			window.PopBack()
		}

		// Add current element at the back of the window
		window.PushBack(i)

	}

	result = append(result, nums[window.Front()])

	for i := windowSize; i < len(nums); i++ {

		// Remove indexes of elements that are smaller than the current element
		// from the back of the window
		for !window.Empty() && nums[i] > nums[window.Back()] {
			window.PopBack()
		}

		// Remove first index from the window deque if
		// it doesn't fall in the current window anymore
		if !window.Empty() && window.Front() <= (i-windowSize) {
			window.PopFront()
		}

		// it doesn't fall in the current window anymore
		window.PushBack(i)
		// Appending the largest element from the window to the result
		result = append(result, nums[window.Front()])

	}

	return result
}

//Very concise version
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0)
	q := make([]int, 0)

	for i, num := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= num {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if i >= k-1 {
			result = append(result, nums[q[0]])
		}

		if q[0] == i-k+1 {
			q = q[1:]
		}
	}

	return result
}

func main() {
	fmt.Printf("%v", findMaxSlidingWindow([]int{10, 6, 9, -3, 23, -1, 34, 56, 67, -1, -4, -8, -2, 9, 10, 34, 67}, 2))
	fmt.Printf("%v", findMaxSlidingWindow([]int{1, 2}, 2))
}

type Deque struct {
	items []int
}

// NewDeque is a constructor that will declare and return the Deque type object
func NewDeque() *Deque {
	return new(Deque)
}

// PushFront will push an element at the front of the dequeue
func (s *Deque) PushFront(item int) {
	temp := []int{item}
	s.items = append(temp, s.items...)
}

// PushBack will push an element at the back of the dequeue
func (s *Deque) PushBack(item int) {
	s.items = append(s.items, item)
}

// PopFront will pop an element from the front of the dequeue
func (s *Deque) PopFront() int {
	defer func() {
		s.items = s.items[1:]
	}()
	return s.items[0]
}

// PopBack will pop an element from the back of the dequeue
func (s *Deque) PopBack() int {
	i := len(s.items) - 1
	defer func() {
		s.items = append(s.items[:i], s.items[i+1:]...)
	}()
	return s.items[i]
}

// Front will return the element from the front of the dequeue
func (s *Deque) Front() int {
	return s.items[0]
}

// Back will return the element from the back of the dequeue
func (s *Deque) Back() int {
	return s.items[len(s.items)-1]
}

// Empty will check if the dequeue is empty or not
func (s *Deque) Empty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

// Len will return the length of the dequeue
func (s *Deque) Len() int {
	return len(s.items)
}
