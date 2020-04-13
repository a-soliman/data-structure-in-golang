package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// Item genseric item
type Item generic.Type

// Node a linked list node
type Node struct {
	value Item
	next  *Node
}

// LinkedList an ll struct
type LinkedList struct {
	head *Node
	size int
}

// ListValues returns a slice contains the linkedlist values
func (ll *LinkedList) ListValues() []Item {
	var values []Item
	currentNode := ll.head
	for {
		values = append(values, currentNode.value)
		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
	}
	return values
}

// Append appends a new node to linked list
func (ll *LinkedList) Append(value Item) {
	newNode := Node{value, nil}
	if ll.head == nil {
		ll.head = &newNode
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &newNode
	}
	ll.size++
}

// Insert inserts a new node at a given position in the linkedlist
func (ll *LinkedList) Insert(i int, value Item) {
	if i > ll.size || i < 1 {
		return
	}
	if i == 1 {
		ll.Append(value)
		return
	}
	newNode := Node{value, nil}
	nodeBefore := ll.head
	step := 1
	for {
		if i == step+1 {
			break
		}
		nodeBefore = nodeBefore.next
		step++
	}
	newNode.next = nodeBefore.next
	nodeBefore.next = &newNode
	ll.size++
}

// RemoveAt removes the node at a given index
func (ll *LinkedList) RemoveAt(i int) {
	if i < 1 || i > ll.size {
		return
	}
	if i == 1 {
		ll.head = ll.head.next
	} else {
		currentNode := ll.head.next
		prevNode := ll.head
		nextNode := currentNode.next
		step := 2
		for {
			if i == step {
				break
			}
			prevNode = currentNode
			currentNode = currentNode.next
			nextNode = currentNode.next
			step++
		}
		prevNode.next = nextNode
	}
	ll.size--
}

// IndexOf returns the position of the item
func (ll *LinkedList) IndexOf(item Item) int {
	if ll.size < 1 {
		return 0
	}
	currentNode := ll.head
	currentIdx := 1

	for {
		if currentNode == nil {
			currentIdx = 0
			break
		}
		if item == currentNode.value {
			break
		}
		currentNode = currentNode.next
		currentIdx++
	}
	return currentIdx
}

func main() {
	ll := LinkedList{}
	ll.Append(10)
	ll.Append(12)
	ll.Insert(2, 5)
	ll.RemoveAt(4)
	// fmt.Print(ll.ListValues())
	fmt.Print(ll.IndexOf(12))
}
