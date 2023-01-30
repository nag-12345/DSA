package main

import "fmt"

/*func main() {
	a := 1
	for a < 10 {
		fmt.Print(" a is ", a, "\n")
		a++
		if a > 5 {
			       /* terminate the loop using break statement
			break
		}
	}

}*/

func main() {
	var a int
	var b int
	for a = 1; a <= 3; a++ {
		for b = 1; b <= 3; b++ {
			if a == 2 && b == 2 {
				break
			}
			fmt.Print(a, " ", b, "\n")
		}
	}

}
