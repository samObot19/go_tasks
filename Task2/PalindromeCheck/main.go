package main

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return -1
	}, s)

	
	
	i := 0
	j := len(s) - 1

	for{
		if i >= j{
			break
		}

		if s[i] != s[j]{
			return false
		}
		i++
		j--

	}

	return true
}

func main() {
	//test
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(IsPalindrome("race a car")) // false
}
