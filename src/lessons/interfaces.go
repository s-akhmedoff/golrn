package lessons

import (
	"fmt"
)

type some interface {
	upgrade(delta []interface{})
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

func (f *foo) upgrade(delta []interface{}) {
	for _, v := range delta {
		switch v.(type) {
		case string:
			fmt.Println(v)

		}
	}
}

func (b *bar) upgrade(delta []interface{}) {
	fmt.Println(delta)
}

func somer(s some, multiply []interface{}) {
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

	somer(&f, []interface{}{1, "str", 2.2})
	somer(&b, []interface{}{1, 2.2, f})

}
