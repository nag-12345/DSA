package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	//s2 := []string{"abc", "bcd"}
	fmt.Println("slice:", s1)
	fmt.Println("len:", len(s1))
	fmt.Println("cap:", cap(s1))

	// fmt.Println("slice:", s2)
	// fmt.Println("len:", len(s2))
	// fmt.Println("cap:", cap(s2))

	res1 := append(s1, 4)
	//res2 := append(s2, "cda")
	fmt.Println("slice:", res1)
	fmt.Println("len:", len(res1))
	fmt.Println("cap:", cap(res1))

	// fmt.Println("slice:", res2)
	// fmt.Println("len:", len(res2))
	// fmt.Println("cap:", cap(res2))

	res3 := append(res1, 5, 6, 7)
	//res4 := append(res2, "jh", "kl")
	fmt.Println("slice:", res3)
	fmt.Println("len:", len(res3))
	fmt.Println("cap:", cap(res3))

	// fmt.Println("slice:", res4)
	// fmt.Println("len:", len(res4))
	// fmt.Println("cap:", cap(res4))

}
