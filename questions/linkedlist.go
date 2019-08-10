package questions

import (
	"fmt"
	"math/rand"
)

// LinkedList interface
type LinkedList interface {
	Append(int)
	Prepend(int)
	Delete(int)
	HasElem(int) bool
	ToArray() []int
}

// DoublyLinkedList is a singly linked list
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// NewDoublyLinkedList returns a new Singly Linked List
func NewDoublyLinkedList(val int) *DoublyLinkedList {
	list := &DoublyLinkedList{&Node{value: val}, nil}
	list.Tail = list.Head
	return list
}

// Append a value to the end of the linked list
func (l *DoublyLinkedList) Append(val int) {
	if l.Head == nil {
		l.Head = &Node{value: val}
		l.Tail = l.Head
		return
	}

	l.Tail.right = &Node{value: val, left: l.Tail}
	l.Tail = l.Tail.right
}

// Prepend a value before the start of the linked list
func (l *DoublyLinkedList) Prepend(val int) {
	if l.Head == nil {
		l.Head = &Node{value: val}
		l.Tail = l.Head
		return
	}

	tmp := &Node{value: val, right: l.Head}
	l.Head.left = tmp
	l.Head = tmp
}

// Delete a node
func (l *DoublyLinkedList) Delete(val int) {
	tmp := l.Head

	for tmp != nil {
		if tmp.value == val {
			if tmp.left == nil {
				l.Head = tmp.right
				l.Head.left = nil
			} else if tmp.right == nil {
				l.Tail = tmp.left
				l.Tail.right = nil
			} else {
				tmp.left.right = tmp.right
				tmp.right.left = tmp.left
			}
		}
		tmp = tmp.right
	}
}

// ToArray converts linked list to an array
func (l DoublyLinkedList) ToArray() []int {
	arr := make([]int, 0)

	tmp := l.Head
	for tmp != nil {
		arr = append(arr, tmp.value)
		tmp = tmp.right
	}

	return arr
}

// HasElem checks if the Linked List contains the value
func (l DoublyLinkedList) HasElem(val int) bool {
	tmp := l.Head

	for tmp != nil {
		if tmp.value == val {
			return true
		}

		tmp = tmp.right
	}

	return false
}

func (l DoublyLinkedList) String() string {
	str := ""

	tmp := l.Head
	for tmp.right != nil {
		str += fmt.Sprintf("%v -> ", tmp.value)
		tmp = tmp.right
	}
	str += fmt.Sprintf("%v", tmp.value)

	return str
}

// GetIntersection finds the intersection between two linked list
func GetIntersection(list1, list2 DoublyLinkedList) DoublyLinkedList {
	tmp := list1.Head
	res := DoublyLinkedList{}

	for tmp != nil {
		if list2.HasElem(tmp.value) {
			res.Append(tmp.value)
		}
		tmp = tmp.right
	}

	return res
}

// GetRandomNode returns a random node from the linked list
func GetRandomNode(list DoublyLinkedList) *Node {
	arr := list.ToArray()

	index := rand.Intn(len(arr)) - 1
	tmp := list.Head

	for index > 0 {
		tmp = tmp.right
		index--
	}

	return tmp
}

// SwapEveryTwo swaps every two nodes in a linked list
func (l *DoublyLinkedList) SwapEveryTwo() {
	l.Head = swapEveryTwo(l.Head)

	newTail := l.Tail

	for newTail.right != nil {
		newTail = newTail.right
	}

	l.Head.left = nil
	l.Tail = newTail
}

func swapEveryTwo(n *Node) *Node {
	if n == nil || n.right == nil {
		return n
	}

	tmp := n.right
	afterTmp := tmp.right

	tmp.right = n
	tmp.left = n.left
	n.left = tmp
	afterTmp.left = n
	n.right = swapEveryTwo(afterTmp)
	return tmp
}

// GetUnion gets the union of two linked list
func GetUnion(list1, list2 DoublyLinkedList) DoublyLinkedList {
	res := DoublyLinkedList{}

	curr := list1.Head
	for curr != nil {
		if !res.HasElem(curr.value) {
			res.Append(curr.value)
		}

		curr = curr.right
	}

	curr = list2.Head
	for curr != nil {
		if !res.HasElem(curr.value) {
			res.Append(curr.value)
		}

		curr = curr.right
	}

	return res
}
