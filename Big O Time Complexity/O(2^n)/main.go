package main

import "fmt"

func main() {
	n, steps := 0, 1
	fmt.Print("Enter the number of disks : ")
	fmt.Scan(&n)
	for ; n > 0; n-- {
		steps = steps * 2
	}
	fmt.Print("Total no. of steps to solve Towe of Hanoi : ",steps-1)
}
