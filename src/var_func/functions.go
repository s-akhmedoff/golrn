package var_func

import "fmt"

func repaet(a int, e error) (int, error) {
	return a, e
}

func temp(some ...int) ([]int) {
	a := make([]int, 3)

	a = some[:2]

	for i := 1; i != 0; i-- {
		a[i] = i
	}

	return some
}

func Function() {
	fmt.Println(repaet(3, nil))
	fmt.Println(temp(1,2,3,4))
}