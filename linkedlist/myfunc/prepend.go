package myfunc

func (ll *LinkedList) Prepend(value int){
	newNode := &Node{Data: value}
	newNode.Next = ll.Head
	ll.Head = newNode
}