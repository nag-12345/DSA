package main

import "fmt"

/*func main() {
	var i int
	fmt.Print("enter i value")
	fmt.Scanln(&i)
	if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
		fmt.Print("leeap")
	} else {
		fmt.Print("not leeap")
	}
}*/
//#########nested if condition ######################
func main() {
	var y int
	fmt.Print("enter y value")
	fmt.Scanln(&y)
	if y%4 == 0 {
		if y%100 == 0 {
			if y%400 == 0 {
				fmt.Print(" leeap")
			} else {
				fmt.Println("not leep")
			}
		} else {
			fmt.Println("  leap")
		}
	} else {
		fmt.Println("not leap")
	}
}
