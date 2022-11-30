package main

func numJewelsInStones(jewels string, stones string) int {
	set := make(map[rune]bool, 0)
	count := 0

	for _, v := range jewels {
		set[v] = true
	}

	for _, stone := range stones {
		if _, ok := set[stone]; ok {
			count++
		}
	}
	return count
}

func main() {
	println(numJewelsInStones("a", "AAbbbbz"))
}
