package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func palindromeCheck(s string) bool {
	
	s = strings.ToLower(s)
	
	var ss []rune
	for _ , word := range s {
		
		if unicode.IsLetter(word) || unicode.IsNumber(word) {
			ss = append(ss, word)
		}

	}

	var left , right int 

	left = 0
	right = len(ss) - 1
	for left < right {
		if ss[left] != ss[right] {
			return false
		} else {
			left++
			right--
		}

	}
	return true

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	s , _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	fmt.Println("answer " , palindromeCheck(s))
}