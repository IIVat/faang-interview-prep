package main

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if circularArrayLoopWithIndex(nums, i) {
			return true
		}
	}
	return false
}

func circularArrayLoopWithIndex(nums []int, index int) bool {
	if len(nums) <= 1 {
		return false
	}

	slow, fast, firstFastPosition := index, index, -1
	for {
		firstFastPosition = nextPosition(fast, nums)

		if nums[fast]*nums[firstFastPosition] < 0 {
			return false
		}

		slow, fast = nextPosition(slow, nums), nextPosition(firstFastPosition, nums)
		if nums[fast]*nums[firstFastPosition] < 0 || fast == firstFastPosition {
			return false
		}
		if slow == fast {
			return true
		}
	}
}

func nextPosition(currPos int, arr []int) int {

	newPosition := (currPos + arr[currPos]) % len(arr)
	if newPosition < 0 {
		newPosition += len(arr)
	}
	return newPosition
}

// abs function takes absolute of the given int
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	println(circularArrayLoop([]int{7, -1, 6, 7, 7}))
}

// func nextPos(curr, step, len int) int {
// 	next := curr + step
// 	if int(math.Abs(float64(step))) >= len {
// 		if step > 0 {
// 			return curr - (step % len)
// 		}
// 		return curr + (step % len)
// 	}

// 	if next < 0 {
// 		return len + next
// 	}

// 	if next >= len {
// 		return next - len
// 	}

// 	return next
// }
