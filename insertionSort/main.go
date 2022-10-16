package main

import (
	"fmt"
)

func main() {
	sliceToSort := []int{26, 32, 21, 33, 9, 8}

	sliceToSort = insertSort(sliceToSort)

	// fmt.Println(sliceToSort)

}

func insertSort(a []int) []int {
	var key, j, i int
	for j = 1; j < len(a); j++ {
		key = a[j]
		fmt.Println(a)
		//0~(j-1) is seq
		i = j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
	return a
}
