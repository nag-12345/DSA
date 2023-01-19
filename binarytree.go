package main

import "fmt"

type Node struct {
	element int
	left    *Node
	right   *Node
}

type BinaryTree struct {
	root *Node
}

func (b *BinaryTree) MakeTree(val int) {
	if b.root == nil {
		b.root = &Node{element: val, left: nil, right: nil}
	} else {
		b.root.insert(val)
	}
	return
}
func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.element {
		if n.left == nil {
			n.left = &Node{element: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{element: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}

}
func (B *BinaryTree) Preorder(troot *Node) {

	if troot != nil {
		fmt.Printf("%d ", troot.element)
		B.Preorder(troot.left)
		B.Preorder(troot.right)
	}
}
func (b *BinaryTree) Inorder(troot *Node) {
	if troot != nil {
		b.Inorder(troot.left)
		fmt.Printf("%d ", troot.element)
		b.Inorder(troot.right)
	}
}
func (b *BinaryTree) Postorder(troot *Node) {
	if troot != nil {
		b.Postorder(troot.left)
		b.Postorder(troot.right)
		fmt.Printf("%d ", troot.element)
	}
}
func (b *BinaryTree) Count(troot *Node) int {
	if troot != nil {
		x := b.Count(troot.left)
		y := b.Count(troot.right)
		return x + y + 1
	}
	return 0
}
func (b *BinaryTree) Height(troot *Node) int {
	if troot != nil {
		x := b.Height(troot.left)
		y := b.Height(troot.right)
		if x > y {
			return x + 1
		} else {
			return y + 1
		}
	}
	return -1
}
func (b *BinaryTree) Search(troot *Node, key int) {
	if troot != nil {
		if key == troot.element {
			fmt.Println("The key is there..!")
		} else if key < troot.element {
			b.Search(troot.left, key)
		} else if key > troot.element {
			b.Search(troot.right, key)
		}
	} else {
		fmt.Println("The key is not there..!")
	}
}

func (b *BinaryTree) levelorder(troot *Node) {
	p := []*Node{}
	a := troot
	fmt.Print(troot.element, " ")
	p = append(p, a)
	for len(p) > 0 {
		a = p[0]
		p = p[1:]
		if a.left != nil {
			fmt.Print(a.left.element, " ")
			p = append(p, a.left)
		}
		if a.right != nil {
			fmt.Print(a.right.element, " ")
			p = append(p, a.right)
		}
	}
}
func main() {
	a := BinaryTree{}
	slic := []int{80, 50, 40, 60, 200, 120, 250}

	for _, v := range slic {
		a.MakeTree(v)
	}

	fmt.Println("Preorder ---->")
	a.Preorder(a.root)
	fmt.Println()
	fmt.Println("iorder ---->")
	a.Inorder(a.root)
	fmt.Println()
	fmt.Println("Postorder ---->")
	a.Postorder(a.root)
	fmt.Println()
	fmt.Println()
	fmt.Println("Height of the Tree Is --> ", a.Height(a.root))
	fmt.Println("Count of the Tree Is --> ", a.Count(a.root))
	a.Search(a.root, 87)
	a.Search(a.root, 50)
	fmt.Println("Levelorder ---->")
	a.levelorder(a.root)

}

/*
x.maketree(40,a,a)
y.maketree(60,a,a)
z.maketree(20,x,a)
r.maketree(50,a,a)
s.maketree(30,r,y)
t.maketree(10,z,s)
*/
