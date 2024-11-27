package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Println("Enter a word : ")
	fmt.Scanln(&word)
	counter := make(map[string]int)

	for i := 0 ; i < len(word) ; i++ {
		ch := rune(word[i])
		if ch >= 65 && ch <= 122 {
			char := string(ch)
			_, exists := counter[char]
			if  exists {
				counter[char] += 1
	
			} else {
				counter[char] = 1
			}
		}
	}
	fmt.Println("The counter dictionary is " , counter)

}