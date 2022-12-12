package main

import (
	"fmt"
	"math"
)

/*
Solution summary

1. Initialize two pointers, indexS1 and indexS2, to zero for iterating both strings.
2. If the character pointed by indexS1 in str1 is the same as the character pointed by indexS2 in str2, increment both pointers.
3. Create two new pointers (start and end) once indexS2 reaches the end of str2. With these two pointers, we’ll slide the window in the opposite direction.
4. Set the start to the value of indexS1 and end to start + 1.
5. ecrement the start pointer until indexS2 becomes less than zero.
6. Decrement indexS2 only if the character pointed by the start pointer in str1 is equal to the character pointed to by indexS2 in str2.
7. Calculate the length of a substring by subtracting values of the end and start variables.
8. If this length is less than the current minimum length, update the length variable and minSubsequence string.
9. Repeat until indexS1 reaches the end of str1.

Time complexity
The time complexity of this solution is O(m×n), where m is the length of str2 and n is the length of str1.

Space complexity
Since we are not using any extra space apart from a few variables, the space complexity is O(1)
O(1)
.
*/

func minWindow(str1 string, str2 string) string {
	idxS1, idxS2 := 0, 0
	minSubsequence := ""

	length := math.MaxInt64

	// Process every every character of str1
	for idxS1 < len(str1) {
		// Check if the character pointed by indexS1 in str1 is the same as the character pointed by indexS2 in str2
		if str1[idxS1] == str2[idxS2] {
			// If the pointed character is the same in both strings increment indexS2
			idxS2++

			// Check if indexS2 has reached the end of str2
			if idxS2 == len(str2) {
				start, end := idxS1, idxS1+1

				// Initialize start to the index where all characters of str2 were present in str1
				idxS2--

				// Decrement pointer indexS2 and start a reverse loop
				for idxS2 >= 0 {
					// Decrement pointer indexS2 and start a reverse loop
					if str1[start] == str2[idxS2] {
						idxS2--
					}
					start--
				}

				start++

				// Update minimum subsequence string to this new shorter string
				if end-start < length {
					// Update minimum subsequence string to this new shorter string
					length = end - start
					// Update minimum subsequence string to this new shorter string
					minSubsequence = str1[start:end]
				}
				// Set indexS1 to start to continue checking in str1 after this discovered subsequence
				idxS1 = start
				idxS2 = 0
			}
		}
		idxS1++

	}

	return minSubsequence
}

func main() {
	fmt.Println(minWindow("abcbdebdde", "bde"))
	fmt.Println(minWindow("fgrqsqsnodwmxzkzxwqegkndaa", "kzed"))
	fmt.Println(minWindow("afgegrwgwga", "aa"))
}

/*
"abcdebdde" , "bde"
"fgrqsqsnodwmxzkzxwqegkndaa" , "kzed"
"michmznaitnjdnjkdsnmichmznait" , "michmznait"
"afgegrwgwga" , "aa"
"abababa" , "ba"
*/
