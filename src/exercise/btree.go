package exercise

import (
	"fmt"
	"math/rand"
)

var height int = 10

//Tree is struct
type Tree struct {
	Right *Tree
	Value int
	Left  *Tree
}

//New - new binary tree
func New(k int) *Tree {

	var t *Tree

	var f func(t *Tree, kk int) *Tree

	f = func(t *Tree, kk int) *Tree {
		if t == nil {
			return &Tree{nil, kk, nil}
		}
		if kk < t.Value {
			t.Left = f(t.Left, kk)
		} else {
			t.Right = f(t.Right, kk)
		}
		return t
	}

	for _, v := range rand.Perm(height) {
		t = f(t, (1+v)*k)
	}
	return t
}

// Walk through Tree
func Walk(t *Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

//Same - is 2 trees is equal
func Same(t1, t2 *Tree) bool {

	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < height; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

//String - formatter
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + "L "
	}
	s += fmt.Sprint(t.Value) + "M"
	if t.Right != nil {
		s += " R" + t.Right.String()
	}
	return "(" + s + ")"
}

func BTree() {

	ch1 := make(chan int, height)
	tree1 := New(2)

	go Walk(tree1, ch1)

	fmt.Println(tree1)

	fmt.Println(Same(New(2), New(2)))  // true
	fmt.Println(Same(New(1), New(2)))  // false
	fmt.Println(Same(New(10), New(5))) // false
	fmt.Println(Same(New(5), New(5)))  // true
	fmt.Println(Same(New(3), New(9)))  // false
}
