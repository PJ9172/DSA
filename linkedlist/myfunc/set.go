package myfunc

func (ll *LinkedList) Set(index, value int)bool{
	temp := ll.Head
	for i:=0; i<index; i++{
		temp = temp.Next
		if temp == nil{
			return false
		}
	}
	temp.Data = value
	return true
}