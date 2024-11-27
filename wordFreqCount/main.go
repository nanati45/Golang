package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func counterFunc() map[string]int {
	
	reader := bufio.NewReader(os.Stdin)
	
	s , _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	fmt.Println(words)

	
	counter := make(map[string]int)
	for _, word := range words {
		for _, ch := range word {
			ch := string(ch)
			ch = strings.ToLower(ch)
			counter[ch]++
		}
	}
	
	fmt.Println(counter) 
	return counter
}

