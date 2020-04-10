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

func main() {
	ll := LinkedList{}
	ll.Append(10)
	ll.Append(12)
	fmt.Print(ll.head.next)

}
