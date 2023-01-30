package main

import "fmt"

func main() {
	fmt.Println(reverse(120))
}
func reverse(x int) int {
	r := 0
	for x > 0 || x < 0 {
		reminder := x % 10
		x /= 10
		r = (r * 10) + reminder

	}
	return r
}
