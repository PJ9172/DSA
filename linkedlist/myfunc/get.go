package myfunc

func (ll *LinkedList) Get(index int) *Node{
	temp := ll.Head
	for i:=0; i<index; i++{
		temp = temp.Next
		if temp == nil{
			return nil
		}
	}
	return temp
}