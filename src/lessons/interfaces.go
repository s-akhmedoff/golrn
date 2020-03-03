package lessons

import (
	"fmt"
)

type some interface {
	upgrade(delta interface{})
}

type foo struct {
	first  int
	second float64
	third  string
}

type bar struct {
	bra string
	bro int
	bru foo
}

func (f *foo) upgrade(delta interface{}) {
	switch delta.(type) {
	case int:
		f.first = delta.(int)
	case float64:
		f.second = delta.(float64)
	case string:
		f.third = delta.(string)
	}
}

func (b *bar) upgrade(delta interface{}) {
	switch delta.(type) {
	case int:
		b.bro = delta.(int)
	case string:
		b.bra = delta.(string)
	case foo:
		b.bru = delta.(foo)
	}
}

func somer(s some, multiply interface{}) {
	s.upgrade(multiply)
}

func Iface() {

	f := foo{
		first:  1,
		second: 2.2,
		third:  "hello ",
	}

	b := bar{
		bra: "world ",
		bro: 2,
		bru: f,
	}

	fmt.Println(f, b)

	somer(&f, 2)
	somer(&b, "hello")

}
