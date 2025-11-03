package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	lowerdWord := strings.ToLower(s)
	lastWord := lowerdWord[len(lowerdWord)-1]

	var trimedWord []rune
	if !('a' <= lastWord && lastWord <= 'z') {
		trimedWord = []rune(strings.ReplaceAll(lowerdWord[:len(lowerdWord)-1], " ", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), "'", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), ",", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), ".", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), ":", ""))
	} else {
		trimedWord = []rune(strings.ReplaceAll(lowerdWord, " ", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), "'", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), ",", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), ".", ""))
		trimedWord = []rune(strings.ReplaceAll(string(trimedWord), ":", ""))

	}

	i := 0

	halfCount := len(trimedWord) / 2
	for i < halfCount {
		if trimedWord[i] != trimedWord[len(trimedWord)-1-i] {
			return false
		}
		i++
	}

	return true
}

func main() {

	fmt.Println(isPalindrome(".a"))
	fmt.Println(isPalindrome("Was it a car or a cat I saw?"))
	fmt.Println(isPalindrome("Madam, in Eden, I'm Adam"))
}
