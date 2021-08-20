package main

import (
	"fmt"
	"strings"
)

const (
	openBracket  = "("
	closeBracket = ")"
)

func main() {
	fmt.Println(findFirstStringInBracket(")(ibit"))
	fmt.Println(findFirstStringInBracket("bibi)t"))
}

func findFirstStringInBracket(str string) string {
	openBracketIdx := strings.Index(str, openBracket)
	closeBracketIdx := strings.Index(str, closeBracket)

	if openBracketIdx >= 0 && openBracketIdx < closeBracketIdx {
		return str[openBracketIdx+1 : closeBracketIdx]
	}

	return ""
}
