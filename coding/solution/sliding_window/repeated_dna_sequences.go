package main

import (
	"fmt"
	"math"
)

func findRepeatedSequences(s string, k int) []string {
	res := make([]string, 0)
	dnaSet := make(map[string]int)

	for i := 0; i <= len(s)-k; i++ {
		dnaSubseq := s[i : i+k]
		// println(dnaSubseq)
		if count, ok := dnaSet[dnaSubseq]; ok && count < 2 {
			res = append(res, dnaSubseq)
		}
		dnaSet[dnaSubseq]++
	}
	return res
}

func findRepeatedSequencesOptimised(s string, k int) []string {
	windowSize := k
	stringOutput := make([]string, 0)

	if len(s) <= windowSize {
		return stringOutput
	}

	// Parameters of rolling hash
	base := 4
	maxPow := int(math.Pow(float64(base), float64(windowSize)))

	// Mapping of a character into an integer
	mapping := map[rune]int{'A': 1, 'C': 2, 'G': 3, 'T': 4}
	numbers := make([]int, 0)

	for _, ch := range s {
		numbers = append(numbers, mapping[ch])
	}

	hashing := 0
	substringHashes, output := make(map[int]bool, 0), make(map[string]int, 0)

	for i := 0; i <= len(s)-windowSize; i++ {
		if i != 0 {
			hashing = hashing * base
			hashing -= numbers[i-1] * maxPow
			hashing += numbers[i+windowSize-1]
			println(hashing)
		} else {
			for j := 0; j < windowSize; j++ {
				hashing = hashing*base + numbers[j]
			}
		}

		if _, ok := substringHashes[hashing]; ok {
			output[s[i:i+windowSize]] = 0
		}
		substringHashes[hashing] = true
	}

	for str, _ := range output {
		stringOutput = append(stringOutput, str)
	}
	return stringOutput
}

func main() {
	// fmt.Printf("%v", findRepeatedSequencesOptimised("TTTTTGGGTTTTCCA", 14))
	fmt.Printf("%v", findRepeatedSequencesOptimised("ACTGC", 3))
	println()

	nucleotides := map[rune]uint32{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	w := uint32(0)
	for _, r := range "ACTTTTGGGTTTTCCA"[:10] {
		w <<= 2
		fmt.Printf("%b\n", w)
		w |= nucleotides[r]
		fmt.Printf("%b\n", w<<2)
	}

	fmt.Printf("%b\n", (w<<2)&(1024*1024-1))
}
