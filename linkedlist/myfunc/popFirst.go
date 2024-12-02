package myfunc

import "fmt"

func (ll *LinkedList) PopFirst() *Node {
	if ll.isEmpty() {
		fmt.Println("LinkedList is empty!!!")
		return nil
	}else{
		temp := ll.Head
		ll.Head = ll.Head.Next
		return temp
	}
}
