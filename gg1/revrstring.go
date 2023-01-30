package main

import "fmt"

func main() {
	//s := []string("hello")
	s := []string{"h", "e", "l", "l", "o"}
	fmt.Println(reverseString(s))
}
func reverseString(s []string) []string {
	empty := make([]string, len(s))

	k := 0
	for i := len(s) - 1; i >= 0; i-- {
		empty[k] = s[i]
		k++
	}

	copy(s, empty)

	return empty

}

///////////
/*package main

import "fmt"
func main() {
	s := []byte("hello")
	//s:=[]string{"h","e","l","l","o"}
	fmt.Println(reverseString(s))
}
func reverseString(s []byte) []string {
	empty := make([]byte, len(s))

	k := 0
	for i := len(s) - 1; i >= 0; i-- {
		empty[k] = s[i]
		k++
	}

	copy(s, empty)
	res := []string{}
	for _, j := range empty {
		res = append(res, string(j))
	}
	return res

}*/
