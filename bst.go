package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// Item genseric item
type Item generic.Type

// BST type
type BST struct {
	value Item
	left  *BST
	right *BST
}

func main() {
	bst := BST{10, nil, nil}
	fmt.Print(bst)
}
