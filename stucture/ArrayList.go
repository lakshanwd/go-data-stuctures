package stucture

import "errors"

//ArrayList - list implementation
type ArrayList struct {
	Size int
	Docs []interface{}
}

//NewArrayList -
func NewArrayList(initialSize int) (*ArrayList, error) {
	if initialSize < 0 {
		return nil, errors.New("invalid size")
	}
	list := ArrayList{Size: initialSize, Docs: make([]interface{}, initialSize)}
	return &list, nil
}

//Get -
func (list ArrayList) Get(index int) (*interface{}, error) {
	e := list.Docs[index]
	return &e, nil
}

//Set -
func (list ArrayList) Set(index int, element *interface{}) error {
	if index < 0 || (list.Size-1) < index {
		return errors.New("index out of range")
	}
	list.Docs[index] = element
	return nil
}

//Add -
func (list ArrayList) Add(index int, element *interface{}) {
	docs := make([]interface{}, list.Size+1)
	list.Docs = append(docs, list.Docs[:index-1], element, list.Docs[index:])
	list.Size = list.Size + 1
}

//Remove -
func (list ArrayList) Remove(index int) error {
	docs := make([]interface{}, list.Size-1)
	list.Docs = append(docs, list.Docs[:index-1], list.Docs[index+1:])
	list.Size = list.Size - 1
	return nil
}
