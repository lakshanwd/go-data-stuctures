package stucture

import "errors"

//ArrayList - list implementation
type ArrayList struct {
	Size int
	Docs []interface{}
}

//NewArrayList - returns new array list
func NewArrayList(initialSize int) *ArrayList {
	list := ArrayList{Size: initialSize, Docs: make([]interface{}, initialSize)}
	return &list
}

//Get - returns element which locates in given index
func (list ArrayList) Get(index int) (*interface{}, error) {
	e := list.Docs[index]
	return &e, nil
}

//Set - replaces element which locates in given index with given element
func (list ArrayList) Set(index int, element *interface{}) error {
	if 0 <= index && index < list.Size {
		list.Docs[index] = element
		return nil
	}
	return errors.New("index out of range")
}

//Add - adds given element to given index
func (list ArrayList) Add(index int, element *interface{}) {
	docs := make([]interface{}, list.Size+1)
	list.Docs = append(docs, list.Docs[:index-1], element, list.Docs[index:])
	list.Size = list.Size + 1
}

//Remove - removes element which located in given index from list
func (list ArrayList) Remove(index int) error {
	docs := make([]interface{}, list.Size-1)
	list.Docs = append(docs, list.Docs[:index-1], list.Docs[index+1:])
	list.Size = list.Size - 1
	return nil
}
