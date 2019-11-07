package queue

import (
	"fmt"
)

const (
	minQSize = 16
)

// Queue defines a queue data structure.
type Queue struct {
	buf  []interface{}
	head int
	tail int
	size int
}

// New constructor
func New() *Queue {

	return &Queue{
		buf: make([]interface{}, minQSize),
	}
}

func (q *Queue) increment() {

	//double size.
	newBuf := make([]interface{}, q.size<<1)
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.size
	q.buf = newBuf
}

// Enqueue an element.
func (q *Queue) Enqueue(e interface{}) {

	if q.size == len(q.buf) {
		q.increment()
	}

	q.buf[q.tail] = e
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.size++
}

// Peek a head element without removing
func (q *Queue) Peek() *interface{} {
	if q.size <= 0 {
		fmt.Println("Peek: Empty queue")
		return nil
	}

	return &q.buf[q.head]
}

// Dequeue an element
func (q *Queue) Dequeue() *interface{} {

	if q.size <= 0 {
		fmt.Println("Pop: Empty queue")
		return nil
	}

	ele := q.buf[q.head]
	q.head++
	q.size--
	return &ele
}

// Size returns the length of queue.
func (q *Queue) Size() int {
	return q.size
}
