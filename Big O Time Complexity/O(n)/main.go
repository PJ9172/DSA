package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println(linearSearch(arr, 30))
	fmt.Println(linearSearch(arr, 70))
}

func linearSearch(arr []int, n int) bool {
	size := len(arr)
	for i := 0; i < size; i++ {
		if arr[i] == n {
			return true
		}
	}
	return false
}
