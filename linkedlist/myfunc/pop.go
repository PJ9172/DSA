package myfunc

import "fmt"

func (ll *LinkedList) Pop() *Node {
	if ll.isEmpty() {
		fmt.Println("Linkedlist is empty!!!")
		return nil
	} else {
		pre := ll.Head
		for pre.Next != ll.Tail {
			pre = pre.Next
		}
		temp := pre
		pre = ll.Tail
		ll.Tail = temp
		ll.Tail.Next = nil
		return pre
	}
}
