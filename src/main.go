package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

//Tree is struct
type Tree struct {
	Right *Tree
	Value int
	Left  *Tree
}

// Walk through Tree
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		fmt.Println()
	}

	if t.Right != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value
	defer close(ch)

	if t.Left != nil {
		Walk(t.Left, ch)
	}
}

func main() {

	fmt.Println()
}
