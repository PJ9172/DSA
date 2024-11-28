package main

import "fmt"

func main() {
	var arr []int = []int{14, 35, 2, 8, 5}
	mergeSort(&arr)
	fmt.Print("Sorted array : ",arr)
}

func mergeSort(arr *[]int){
	l := len(*arr)
	
}