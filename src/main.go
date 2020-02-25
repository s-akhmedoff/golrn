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

}

func main() {

	fmt.Println()
}
