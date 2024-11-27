package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func palindromeCheck(s string) bool {
	
	s = strings.ToLower(s)

	
	var ss string
	for _ , word := range s {
		
		ch := rune(word)
		if ch >= 65 && ch < 96 {
			ss += string(word)
		}

	}

	var left , right int 

	left = 0
	right = len(ss) - 1

	for left < right {
		if ss[left] != ss[right] {
			return false
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