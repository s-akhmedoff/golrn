package lessons

import (
	"fmt"
	"reflect"
)

type A struct {
	a int
	c string
	e *A
}

func (a A) getProperties() (int, string, *A) {

	if a.e == nil {
		return a.a, a.c, nil
	}

	return a.a, a.c, a.e
}

func (a *A) setProperties(t interface{}) reflect.Kind {

	switch t.(type) {
	case int:
		(*a).a = t.(int)
		return reflect.Int
	case string:
		(*a).c = t.(string)
		return reflect.String
	case *A:
		(*a).e = t.(*A)
		return reflect.Struct
	}

	return reflect.Invalid
}

func (a *A) conProperties(t interface{}) (kind reflect.Kind, old_val, new_val interface{}) {

	switch t.(type) {
	case int:
		old_val = (*a).a
		(*a).a += t.(int)
		return reflect.Int, old_val, (*a).a
	case string:
		old_val = (*a).c
		(*a).c += t.(string)
		return reflect.String, old_val, (*a).c
	case *A:
		fmt.Println("cannot concrete struct")
	}

	return reflect.Invalid, nil, nil
}

func (a A) showOnlyInt() int {
	return a.a
}

func (a A) showOnlyString() string {
	return a.c
}

func (a A) showOnlyStruct() *A {
	return a.e
}

func (a *A) incIntField() {
	a.a++
}

func Methods() {
	a := new(A)

	a.setProperties(1)
	a.setProperties("foo")
	a.setProperties(&A{2, "ping", nil})

	fmt.Println(a.getProperties())
	a.conProperties(4)
	fmt.Println(a.conProperties("bar"))

	fmt.Println(a.getProperties())
	fmt.Println(a.showOnlyInt())
	fmt.Println(a.showOnlyString())
	fmt.Println(a.showOnlyStruct())
}
