package main

import "fmt"

type Node struct {
	element int
	next    *Node
}
type List struct {
	head *Node
	tail *Node
	len  int
}

func (obj *List) display() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Println("--->", comm.element)
		comm = comm.next
	}
	fmt.Println()
}
func (obj *List) search(val int) bool {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		if comm.element == val {
			return true
		}
		comm = comm.next
	}
	return false
}
func (obj *List) addfirst(val int) {
	newnode := Node{element: val}
	temp := obj.head
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
	} else {
		obj.head = &newnode
		newnode.next = temp
	}
	obj.len++
}
func (obj *List) append(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode

	} else {
		obj.tail.next = &newnode

	}
	obj.tail = &newnode
	obj.len++
}
func (obj *List) addanypos(pos, val int) {
	newnode := Node{element: val}
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	temp := comm.next
	comm.next = &newnode
	newnode.next = temp
	obj.len++
}
func (obj *List) delfirst() {
	if obj.len == 0 {
		fmt.Println("no list")
		return
	}
	comm := obj.head
	obj.head = comm.next
	obj.len--
}
func (obj *List) delend() {
	if obj.len == 0 {
		fmt.Println(("no list"))
		return
	}
	obj.tail.next = nil
	obj.tail = nil
	obj.len--
}
func (obj *List) delanypos(pos int) {
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	comm.next = comm.next.next
	obj.len--
}
func main() {
	obj := List{}
	slice := []int{10, 20, 30}
	for _, v := range slice {
		obj.append(v)
	}
	a := obj.search(20)
	fmt.Println(a)

	obj.addfirst(100)
	obj.addfirst(200)
	obj.display()
	obj.addanypos(2, 500)
	obj.display()
	obj.delfirst()
	obj.display()
	obj.delend()
	obj.display()
	obj.delanypos(3)
	obj.display()
}
