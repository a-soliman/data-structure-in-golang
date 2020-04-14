package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// Item genseric item
type Item generic.Type

// BST type
type BST struct {
	value int
	left  *BST
	right *BST
}

// Insert insert a new bst node
func (bst *BST) Insert(value int) {
	if value < bst.value {
		if bst.left == nil {
			newNode := BST{value, nil, nil}
			bst.left = &newNode
		} else {
			bst.left.Insert(value)
		}
	} else {
		if bst.right == nil {
			newNode := BST{value, nil, nil}
			bst.left = &newNode
		} else {
			bst.right.Insert(value)
			return
		}
	}
}

func main() {
	bst := BST{10, nil, nil}
	bst.Insert(9)
	bst.Insert(8)
	fmt.Print(bst.left)
}
