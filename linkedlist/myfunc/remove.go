package myfunc

import "fmt"

func (ll *LinkedList) Remove(index int) *Node {
	if ll.isEmpty() {
		fmt.Println("Linkedlist is empty!!!")
		return nil
	} else {
		if index == 1{
			return ll.PopFirst()
		}else {
			temp := ll.Get(index)
			pre := ll.Get(index-1)
			pre.Next = temp.Next
			return temp
		}
	}
}
