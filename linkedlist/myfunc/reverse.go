package myfunc

import "fmt"

func (ll *LinkedList) Reverse() {
	if ll.isEmpty() {
		fmt.Println("Linkedlist is empty!!!")
	} else {
		temp := ll.Head
		ll.Head = ll.Tail
		ll.Tail = temp
		pre, nex := &Node{}, &Node{}
		pre,nex = nil, nil
		for temp != nil {
			nex = temp.Next
			temp.Next = pre
			pre = temp
			temp = nex
		}
	}
}
