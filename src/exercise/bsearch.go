package exercise

import (
	"fmt"
)

//	@TODO: Slice method

func bsearch(arr []int, t int) int {
	i, j := 0, len(arr)-1

	for i != j {
		mid := (j + i + 1) / 2

		if t < arr[mid] {
			j = mid - 1
		} else {
			i = mid
		}
	}

	if t == arr[i] {
		return i
	}

	return -1
}

func bsearch_recursive(arr []int, search, start, end int) int {
	for start != end {
		mid := (start + end + 1) / 2
		if search < arr[mid] {
			end = mid - 1
			return bsearch_recursive(arr, search, start, end)
		} else {
			start = mid
			return bsearch_recursive(arr, search, start, end)
		}
	}
	return start
}

func BSearch() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(bsearch_recursive(a, 8, 0, len(a)-1))
	fmt.Println(bsearch(a, 5))

}
