package datastructure

import "fmt"

// BinaryTree is a binary tree interface
type BinaryTree interface {
	InOrder()
	PreOrder()
	PostOrder()
	IsValid() bool
	BottomLeftNode() int
	Add(interface{})
}

// Node for binary tree; holds integer values
type Node struct {
	left  *Node
	right *Node
	value int
}

// NewTree returns a root node for a binary tree
func NewTree(rootValue interface{}) *Node {
	return &Node{nil, nil, rootValue.(int)}
}

// Add inserts a new value into binary tree
func (n *Node) Add(value interface{}) {
	add(n, value.(int))
}

func add(n *Node, value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{nil, nil, value}
		} else {
			add(n.left, value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{nil, nil, value}
		} else {
			add(n.right, value)
		}
	}
}

// InOrder prints all values in order
func (n *Node) InOrder() {
	inOrder(n)
}

func inOrder(n *Node) {
	if n == nil {
		return
	}

	inOrder(n.left)
	fmt.Printf("%d ", n.value)
	inOrder(n.right)
}

// PostOrder prints all values in postorder
func (n *Node) PostOrder() {
	postOrder(n)
}

func postOrder(n *Node) {
	if n == nil {
		return
	}

	inOrder(n.left)
	inOrder(n.right)
	fmt.Printf("%d ", n.value)
}

// PreOrder prints all values in preorder
func (n *Node) PreOrder() {
	preOrder(n)
}

func preOrder(n *Node) {
	if n == nil {
		return
	}

	fmt.Printf("%d ", n.value)
	inOrder(n.left)
	inOrder(n.right)
}

// IsValid checks if the BST is valid
func (n *Node) IsValid() bool {
	if n == nil {
		return true
	}

	if n.left != nil && n.value < n.left.value || n.right != nil && n.value > n.right.value {
		return false
	}

	return n.left.IsValid() && n.right.IsValid()
}

// BottomLeftNode gets the bottom leftmost value
func (n Node) BottomLeftNode() int {
	curr := &n
	queue := make([]*Node, 0)
	queue = append(queue, curr)

	for len(queue) > 0 {
		curr = queue[0]
		queue = queue[1:]

		if curr.right != nil {
			queue = append(queue, curr.right)
		}
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
	}

	return curr.value
}
