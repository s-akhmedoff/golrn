package intership

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	temp := string(runes)
	return temp
}

func reverseInt(i int) int {
	revrs := 0
	for i > 0 {
		lastDigit := i % 10
		revrs = (revrs * 10) + lastDigit
		i = i / 10
	}

	return revrs
}

func fib(n int) int {
	if n < 2 {
		return 1
	}

	return fib(n-2) + fib(n-1)
}

func fizzbuzz(arr []int) {
	for i, v := range arr {
		if v%5 == 0 && v%3 == 0 {
			fmt.Printf("[%v]: %v = FIZZBUZZ\n", i, v)
		} else if v%5 == 0 {
			fmt.Printf("[%v]: %v = BUZZ\n", i, v)
		} else if v%3 == 0 {
			fmt.Printf("[%v]: %v = FIZZ\n", i, v)
		}
	}
}

func palindrom(data interface{}) bool {
	switch data.(type) {
	case int:
		revrs := reverseInt(data.(int))
		if revrs == data.(int) {
			return true
		}
	case string:
		data = strings.Replace(data.(string), " ", "", -1)
		data = strings.ToLower(data.(string))
		revrs := reverseString(data.(string))
		if revrs == data.(string) {
			return true
		}
	}
	return false
}

func oddEvenSum(arr []int) (oddSum, evenSum int) {
	for _, v := range arr {
		if v%2 == 0 {
			evenSum += v
		}
		oddSum += v
	}
	return
}

func Exercise() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i))
	}

	fizzbuzz([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 15})

	fmt.Println(palindrom(13131))

	fmt.Println("Odd and Even sum: ")
	fmt.Println(oddEvenSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
