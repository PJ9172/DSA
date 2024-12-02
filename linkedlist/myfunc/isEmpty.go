package myfunc

func (ll *LinkedList) isEmpty()bool{
	if ll.Head == nil{
		return true
	}else{
		return false
	}
}