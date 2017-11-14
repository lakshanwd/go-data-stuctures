package stucture

type node struct {
	left  *node
	right *node
	e     *interface{}
}

//LinkedList - linkedlist implementation
type LinkedList struct {
}

//NewLinkedList - makes a new linkedlist
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
