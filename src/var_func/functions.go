package var_func

import (
	"fmt"
	"strconv"
)

func _(a int, b int) int {
	return a + b
}

func _(a, b int) (int, int) {
	return b, a
}

func _(vars ...int) (sum int) {
	for _, v := range vars {
		sum += v
	}

	return
}

func _(str string, num int) string {
	return str + strconv.Itoa(num)
}

func _(variable float64) int {
	return func(a float64) int {
		return int(a)
	}(variable)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func workWithPointer(a *int) *int {
	temp := a
	*temp += *a
	return temp
}

func Func() {
	factorial(10)
	a := 3
	fmt.Println(*workWithPointer(&a))

}
