package main

import (
	"fmt"
	"math/rand"
)

var depth int = 10

//Tree - Binary Tree Type Structure
type Tree struct {
	Right *Tree
	Value int
	Left  *Tree
}

//New - Function To Create New Binary Tree Recursively
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

	for _, v := range rand.Perm(depth) {
		t = f(t, (1+v)*k)
	}
	return t
}

//Walk - Function To Recursively Walk And Read Values To Channel From Binary Tree
func Walk(t *Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

//Same - Function To Compare Two Tree By Their Values
func Same(t1, t2 *Tree) bool {

	if t1 == nil || t2 == nil {
		return false
	}

	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < depth; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

//String - Formatting Method
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

func main() {

	ch1 := make(chan int)
	tree1 := New(2)

	go Walk(tree1, ch1)

	fmt.Println(tree1)

	fmt.Println(Same(New(2), New(2))) // true
	fmt.Println(Same(New(1), New(2))) // false
}
