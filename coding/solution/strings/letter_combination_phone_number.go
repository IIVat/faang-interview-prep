package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	nums := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	queue := []string{}

	for _, ch := range nums[rune(digits[0])] {
		queue = append(queue, string(ch))
	}

	for _, d := range digits[1:] {

		len := len(queue)

		for i := 0; i < len; i++ {
			for _, ch := range nums[d] {
				queue = append(queue, queue[i]+string(ch))
			}

		}
		queue = queue[len:]

	}
	return queue
}

func enqueue(queue []int, element int) []int {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []int) []int {
	element := queue[0] // The first element is the one to be dequeued.
	fmt.Println("Dequeued:", element)
	return queue[1:] // Slice off the element once it is dequeued.
}

func main() {
	fmt.Printf("%v", letterCombinations("23245"))
}
