package main

/*
1. Left, right := 0
2. Iterate while left < right
4. Check if s[l] == s[r] => continue
6. If s[l] != s[r-1] && s[r] != s[l+1] => false
7. Else if s[l] == s[r-1] => l++, r-=2
8. Else if s[r] == s[l+1] => r—, l+=2
9. If l+1 == r return true
*/

func validPalindrome(s string) bool {

	isValid, i, j := isPalindrome(s, 0, len(s)-1)

	if isValid {
		return true
	}

	if ok, _, _ := isPalindrome(s, i+1, j); ok {
		return true
	}

	if ok, _, _ := isPalindrome(s, i, j-1); ok {
		return true
	}
	return false
}

func isPalindrome(s string, i, j int) (bool, int, int) {
	for i < j {
		if s[i] != s[j] {
			return false, i, j
		}
		i++
		j--
	}

	return true, 0, 0

}

func main() {
	println(validPalindrome("ececabbacec"))
}
