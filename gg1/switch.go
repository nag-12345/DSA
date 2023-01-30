package main

import "fmt"

func main() {
	day := 2
	switch day {

	case 1, 3, 5:
		fmt.Println("odd")
	case 2, 4:
		fmt.Println("eve")
	case 6, 7:
		fmt.Print("weekend")
	default:
		fmt.Print("invalid number")

	}

}
