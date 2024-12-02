package myfunc

import "fmt"

func (ll *LinkedList) Display() {
	if ll.isEmpty() {
		fmt.Println("Linked list is empty!!!")
	} else {
		temp := ll.Head
		for temp != nil {
			fmt.Printf("%d->", temp.Data)
			temp = temp.Next
		}
		fmt.Print("Nil")
	}
}
