package myfunc

func (ll *LinkedList) Append(value int) {
	var newNode Node
	newNode.Data = value
	newNode.Next = nil
	if ll.Head == nil {
		ll.Head = &newNode
		ll.Tail = &newNode
	} else {
		ll.Tail.Next = &newNode
		ll.Tail = &newNode
	}
}