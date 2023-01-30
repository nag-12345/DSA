package main

import (
	"fmt"
)

/*func main() {
	i := 1
	for i <= 10 {
		product := 4 * i
		fmt.Printf("4*%d=%d\n", i, product)

		i++
	}
}*/

//for loop#################

func main() {
	/*var n int
	  fmt.Print("enter int")
	  fmt.Scan(&n)*/
	n := 6
	var i = 5
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, "x", i, "=", n*i)
		i++
	}
}
