package main

func reverseWords(sentence string) string {

	strRes := []byte{}
	i := 0
	j := len(sentence) - 1

	for {
		if sentence[i] != ' ' {
			break
		}
		i++
	}

	for {
		if sentence[j] != ' ' {
			break
		}
		j--
	}

	for i <= j {
		if sentence[i] == ' ' && strRes[len(strRes)-1] != ' ' {

			strRes = append(strRes, sentence[i])
		} else if sentence[i] != ' ' {
			strRes = append(strRes, sentence[i])
		}
		i++
	}

	//reverse the whole string
	i = 0
	j = len(strRes) - 1
	for i < j {
		strRes[i], strRes[j] = strRes[j], strRes[i]
		j--
		i++
	}

	left, right := 0, 0
	for i, v := range strRes {
		if v == ' ' || i == len(strRes)-1 {
			if i == len(strRes)-1 {
				right = i
			} else {
				right = i - 1
			}

			for left <= right {
				strRes[left], strRes[right] = strRes[right], strRes[left]
				right--
				left++
			}
			left = i + 1
		}
	}

	return string(strRes)
}

func main() {
	println(reverseWords("   fffff ff gg ee"))
}
