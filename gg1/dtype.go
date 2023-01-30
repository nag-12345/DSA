package main

import "fmt"

func main() {
	var a bool = true
	var b int = 6
	c := 100
	f := "hello"
	y := 1000908986857.8

	fmt.Print(a)
	fmt.Printf("\n%T\n", c)
	fmt.Println(b)
	fmt.Printf("\n%T  %v", f, f)
	fmt.Printf("\n%T\n", y)

}
