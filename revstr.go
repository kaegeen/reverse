package main

import (
	"fmt"
)

// Function to reverse a string
func reverseString(s string) string {
	// Convert the string to a rune slice to handle Unicode characters
	runes := []rune(s)
	n := len(runes)

	// Reverse the rune slice
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	// Convert the rune slice back to a string
	return string(runes)
}

func main() {
	var input string

	// Prompt the user to enter a string
	fmt.Print("Enter a string to reverse: ")
	fmt.Scanln(&input)

	// Reverse the string and display the result
	reversed := reverseString(input)
	fmt.Printf("Reversed string: %s\n", reversed)
}
