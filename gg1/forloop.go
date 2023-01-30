package main

import "fmt"

/*nested for loop#####################33/*/
func main() {
	for a := 0; a < 3; a++ {
		for b := 3; b > 0; b-- {
			fmt.Print(a, " ", b, "\n")
		}
	}
}

/*for loop#############################
func main() {
	for a := 0; a < 11; a++ {
		fmt.Print(a, "\n")
	}
}*/

/*infinitive for loop
func main() {
	for true {
		fmt.Printf("This is babu \n")
	}
}*/
