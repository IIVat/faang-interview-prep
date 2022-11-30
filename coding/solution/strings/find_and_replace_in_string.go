package main

import "strings"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	res := ""
	hmap := make(map[int]([]string))

	for i, idx := range indices {
		hmap[idx] = []string{sources[i], targets[i]}
	}

	i := 0

	for i < len(s) {

		if val, ok := hmap[i]; ok && strings.HasPrefix(s[i:], val[0]) {
			res += val[1]
			i += len(val[0])

		} else {
			res += string(s[i])
			i++
		}

	}

	return res
}

type Pair struct {
	src    string
	target string
}

func main() {
	println(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
	println(findReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
	println(findReplaceString("abc", []int{0, 1}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
}
