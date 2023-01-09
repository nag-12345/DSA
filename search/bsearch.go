package main

import (
	"fmt"
	// "math"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(binarySearchAsc(arr, 4))

}
func binarySearchAsc(arr []int, key int) int {
	s := 0
	e := len(arr) - 1
	var mid int

	for s <= e {
		mid = int((s + e) / 2)
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return -1
}
