package main

import (
	"fmt"
)

func main() {

	var floatValue float64 = 5.45
	// type conversion from float to int
	var intValue int = int(floatValue)

	fmt.Printf("Float Value is %g\n", floatValue)
	fmt.Printf("Integer Value is %d\n", intValue)

	intValue = 5
	// type conversion from int to float
	floatValue = float64(intValue)

	fmt.Printf("%d\n", intValue)
	fmt.Printf("%f\n", floatValue)

	i := 5
	f := float64(i)
	fmt.Printf("%f\n", f)

}
