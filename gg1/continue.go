/*package main

import "fmt"

func main() {
	a := 1
	for a < 10 {
		if a == 5 {
			a = a + 1
			continue
		}
		fmt.Printf("a: %d\n", a)
		a++
	}
}*/

package main

import "fmt"

func main() {
	var a int = 1
	var b int = 1
	for a = 1; a < 3; a++ {
		for b = 1; b < 3; b++ {
			if a == 2 && b == 2 {
				continue
			}
			fmt.Printf("a and b is %d %d\n", a, b)
		}
		fmt.Printf(" a and b is %d %d\n", a, b)
	}
}
