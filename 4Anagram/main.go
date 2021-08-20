package main

import "fmt"

func main() {
	var words = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	var results = make([][]string, 0)

	for {
		if len(words) < 1 {
			break
		}

		var nWords = make([]string, 0)
		var tmp = []string{words[0]}

		for i := 1; i < len(words); i++ {
			if isValidAnagram(words[0], words[i]) {
				tmp = append(tmp, words[i])
			} else {
				nWords = append(nWords, words[i])
			}
		}

		results = append(results, tmp)
		words = nWords
	}

	fmt.Println(results)
}

func isValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tmp := [26]rune{}

	for _, val := range s {
		tmp[val-'a']++
	}

	for _, val := range t {
		tmp[val-'a']--
	}

	for _, val := range tmp {
		if val != 0 {
			return false
		}
	}

	return true
}
