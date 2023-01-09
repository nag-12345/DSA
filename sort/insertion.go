package main

import (
	"fmt"
)

func insert(a []int) []int {
	res := make([]int, len(a)) //creat new slice equal to original
	copy(res, a)
	for i := 1; i < len(res); i++ {
		key := res[i] //hold present value
		j := i - 1    // hold previous value
		for j >= 0 && res[j] > key {
			res[j+1] = res[j] //previous value grater then so it forward 1 position
			j--
		}
		res[j+1] = key
	}
	return res

}
func main() {
	a := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(a)
	fmt.Println(insert(a))
}
