package main

import "fmt"

func main() {
	arr := []int{4, 1, 8, 2, 7}
	fmt.Println("Sorted array : ")
	arr = sortArray(arr)
	fmt.Println(arr)
}

func sortArray(arr []int) []int {
	var t int
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := i; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				t = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = t
			}
		}
	}
	return arr
}
