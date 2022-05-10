package main

import "fmt"

var count int = 0

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}

	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}
func (n *Node) Search(k int) bool {
	count = count + 1
	if n == nil {
		return false
	}
	if n.Key < k {
		return n.Right.Search(k)

	} else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {

	tree := &Node{Key: 100}
	fmt.Println(tree)
	tree.Insert(100)
	tree.Insert(200)
	tree.Insert(176)
	fmt.Println(tree.Search(176))
	fmt.Println(count)
}
