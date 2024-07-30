package main

import (
	"strings"
	"fmt"
)



func GetWordFrequency(inputString string) map[string] int{
	WordFrequency := make(map[string] int)

	inputString = strings.ToLower(strings.ReplaceAll(inputString, ",", ""))
	inputString = strings.ReplaceAll(inputString, ".", "")
	inputString = strings.ReplaceAll(inputString, "!", "")
	inputString = strings.ReplaceAll(inputString, "?", "")

	words := strings.Split(inputString, " ")

	for _, word := range words{
		WordFrequency[word]++
	}

	return WordFrequency 

}

func main(){
	test := "The quick brown fox jumps over the lazy dog. The dog barks at the fox."
	fmt.Println(GetWordFrequency(test))
}
