package main

import (
	"fmt"
	"mylinkedlist/myfunc"
)

func main() {
	ll := &myfunc.LinkedList{}
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)
	fmt.Print("LinkedList is : ")
	ll.Display()
	fmt.Println("")

	//pop()
	fmt.Print("\nPoped value is  : ", ll.Pop().Data)
	fmt.Print("\nAfter Pop : ")
	ll.Display()
	fmt.Println("")

	//prepend()
	ll.Prepend(0)
	fmt.Print("\nAfter Prepend 0 : ")
	ll.Display()
	fmt.Println("")

	//popFirst()
	fmt.Print("\nPoped value is : ", ll.PopFirst().Data)
	fmt.Print("\nAfter PopFirst : ")
	ll.Display()
	fmt.Println("")

	//get()
	fmt.Print("\nGetting the 2 index value : ", ll.Get(2).Data)
	fmt.Println("")

	//set()
	fmt.Print("\nSetting the 1 index value to 5 ")
	result := ll.Set(1, 5)
	if !result {
		fmt.Print("\nError to set the value!!!")
	} else {
		fmt.Print("\nAfter Updation : ")
		ll.Display()
		fmt.Println("")
	}

	//insert()
	fmt.Print("\nInserting 2 at 1st index ")
	result = ll.Insert(1, 2)
	if !result {
		fmt.Print("\nError to insert value!!!")
	} else {
		fmt.Print("\nAfter Insertion : ")
		ll.Display()
		fmt.Println("")
	}

	//remove()
	fmt.Print("\nRemoving 2nd index node : ", ll.Remove(2).Data)
	fmt.Print("\nAfter Deletion : ")
	ll.Display()
	fmt.Println("")

	//reverse()
	fmt.Print("\nAfter reversing the list : ")
	ll.Reverse()
	ll.Display()
	fmt.Println()
}
