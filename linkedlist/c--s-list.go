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

// count of the list
func (obj *List) count() int {
	return obj.len
}

//display each node in the list
func (obj *List) display() {
	comm := obj.head
	for i := 0; i < obj.len; i++ {
		fmt.Printf("%v==>", comm.element)
		comm = comm.next
	}
	fmt.Println()
}

// create a node and append in list
func (obj *List) adding(val int) {
	newnode := Node{element: val}

	if obj.len == 0 {
		obj.head = &newnode
	} else {
		newnode.next = obj.tail.next
		obj.tail.next = &newnode
	}
	obj.tail = &newnode
	obj.len++
}

//append a node at beginning
func (obj *List) insertbegg(val int) {
	newnode := Node{element: val}
	if obj.len == 0 {
		obj.head = &newnode
		obj.tail = &newnode
		newnode.next = &newnode
	} else {
		obj.tail.next = &newnode
		newnode.next = obj.head
		obj.head = &newnode
	}
	obj.len++

}

//insert a node in between nodes
func (obj *List) insert(pos, val int) {
	newnode := Node{element: val}
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	newnode.next = comm.next
	comm.next = &newnode
	obj.len++
}

//delete node at starting
func (obj *List) deletebegg() {
	obj.head = obj.head.next
	obj.tail.next = obj.head
	obj.len--
}

//delete node at end of the list
func (obj *List) deleteend() {
	comm := obj.head
	for i := 1; i < obj.len-2; i++ {
		comm = comm.next
	}
	add := comm.next
	add.next = obj.head
	obj.len--
}

//delete a node in between nodes
func (obj *List) deleteat(pos int) {
	comm := obj.head
	for i := 1; i < pos-1; i++ {
		comm = comm.next
	}
	comm.next = comm.next.next
	obj.len--

}
func main() {
	obj := List{}
	obj.adding(10)
	obj.adding(20)
	obj.adding(30)
	obj.adding(40)
	fmt.Println(obj.count())
	obj.display()

	obj.insertbegg(1000)
	obj.display()

	obj.adding(3000) //adding node at end
	obj.display()

	obj.insert(4, 2000)
	obj.display()

	obj.deletebegg()
	obj.display()

	obj.deleteend()
	obj.display()

	obj.deleteat(3)
	obj.display()

}
