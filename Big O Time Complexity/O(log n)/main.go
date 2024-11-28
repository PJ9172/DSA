package main

import "fmt"

func main() {
	var arr []int = []int{10, 20, 30, 40, 50}
	var n int
	fmt.Println("Enter number to search in array : ")
	fmt.Scan(&n)
	binarySearch(arr, n)
	fmt.Println(arr)
}

func binarySearch(arr []int, n int) {
	l := 0
	u := len(arr) - 1
	for l <= u {
		mid := (l + u) / 2
		if arr[mid] == n {
			fmt.Println(n, "found in array!!!")
			return
		}
		if n < arr[mid] {
			u = mid - 1
		} else {
			l = mid + 1
		}
	}
	fmt.Println(n, "not found in array!!!")
}
