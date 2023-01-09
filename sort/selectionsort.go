package main

import "fmt"

func selectionsort(n []int) []int {
	length := len(n)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if n[j] < n[i] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return n
}
func main() {
	k := []int{3, 66, 1, 100, 78, 4, 99, 1, -88}
	fmt.Println(k)
	selectionsort(k)
	fmt.Println(k)

}
