package main

/*
Solution summary
To recap, the solution to this problem can be divided into the following parts:

Traverse the input string.

Use a hash map to store elements along with their respective indexes.

If the current element is present in the hash map, check whether it’s already present in the current window. If it is, we have found the end of the current window and the start of the next. We check if it’s longer than the longest window seen so far and update it accordingly.

Store the current element in the hash map with the key as the element and the value as the current index.

At the end of the traversal, we have the length of the longest substring with all distinct characters.

Time complexity
We have to iterate over all the n
n
 elements in the string. Therefore, the time complexity is O(n)
O(n)
.

Space complexity
We need extra space to store the last occurrence of each element. In the worst-case scenario, all of the elements can be unique and we need to store all n
n
 elements. Therefore, the space complexity will be O(n)
O(n)
.
*/

func findLongestSubstring(s string) int {
	maxLen := 0
	characters := make(map[rune]int, 0)
	l := 0

	for r, c := range s {
		if v, ok := characters[c]; ok && v >= l {
			l = v + 1
		}

		characters[c] = r

		if (r - l + 1) > maxLen {
			maxLen = r - l + 1
		}

	}

	if maxLen == 0 {
		return 0
	}

	return maxLen
}

func main() {
	println(findLongestSubstring("abcdbea"))
	println(findLongestSubstring("pwwkew"))
}
