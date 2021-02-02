package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		connect(root.Right)
		connect(root.Left)

	}
	return root
}

func main() {
	root := &Node{
		Val: 1,
	}
	root.Left = &Node{
		Val: 2,
	}
	root.Right = &Node{
		Val: 3,
	}
	root.Left.Left = &Node{
		Val: 4,
	}
	root.Left.Right = &Node{
		Val: 5,
	}
	root.Right.Left = &Node{
		Val: 6,
	}
	root.Right.Right = &Node{
		Val: 7,
	}
	root.Left.Left.Left = &Node{
		Val: 8,
	}
	root.Left.Left.Right = &Node{
		Val: 9,
	}
	root.Left.Right.Left = &Node{
		Val: 10,
	}
	root.Left.Right.Right = &Node{
		Val: 11,
	}
	root.Right.Left.Left = &Node{
		Val: 12,
	}
	root.Right.Left.Right = &Node{
		Val: 13,
	}
	root.Right.Right.Left = &Node{
		Val: 14,
	}
	root.Right.Right.Right = &Node{
		Val: 15,
	}

	root = connect(root)
	fmt.Println(root)
}
