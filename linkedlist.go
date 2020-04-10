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

func main() {
	ll := LinkedList{}
	ll.Append(10)
	ll.Append(12)
	ll.Insert(2, 5)
	fmt.Print(ll.head.next)

}
