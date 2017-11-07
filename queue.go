package datastuctures

import "errors"

//Queue - queue implementation
type Queue struct {
	Front int
	Rear  int
	Size  int
	Docs  []interface{}
}

//NewQueue - makes a new queue
func NewQueue(size int) *Queue {
	q := Queue{Front: -1, Rear: -1, Docs: make([]interface{}, size), Size: size}
	return &q
}

//Enqueue - puts an element at the end of the queue
func (q *Queue) Enqueue(e interface{}) error {
	if ((q.Rear + 1) % q.Size) == q.Front {
		return errors.New("Full")
	} else if q.IsEmpty() {
		q.Front = 0
		q.Rear = 0
	} else {
		q.Rear = (q.Rear + 1) % q.Size
	}
	q.Docs[q.Rear] = e
	return nil
}

//Dequeue - removes an element from the start of the queue
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Empty List")
	} else if q.Front == q.Rear {
		e := q.Docs[q.Front]
		q.Front = -1
		q.Rear = -1
		return e, nil
	} else {
		e := q.Docs[q.Front]
		q.Front = (q.Front + 1) % q.Size
		return e, nil
	}
}

//IsEmpty - checks whether queue is empty or not
func (q *Queue) IsEmpty() bool {
	return q.Front == -1 && q.Rear == -1
}

//IsFull - checks whether queue is full or not
func (q *Queue) IsFull() bool {
	return false
}
