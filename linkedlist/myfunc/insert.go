package myfunc

func (ll *LinkedList) Insert (index, value int)bool{
	if ll.Head == nil{
		ll.Append(value)
		return true
	}else {
		temp := ll.Get(index-1)
		if temp == nil{
			return false
		}
		newNode := &Node{Data: value}
		newNode.Next = temp.Next
		temp.Next = newNode
		return true
	}
}