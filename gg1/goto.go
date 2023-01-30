package main

import (
	"fmt"
)

func main() {
	var input int
Loop:
	fmt.Print("You are not eligible to vote ")
	fmt.Print("Enter your age ")
	fmt.Scanln(&input)
	if input <= 18 {
		goto Loop
	} else {
		fmt.Print("You can vote ")
	}
}
