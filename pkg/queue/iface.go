package queue

import "io"

type Message struct {
	priority int
	payload  io.ReadCloser
}

type PriorityQueue interface {
	Enqueue(message Message) error

	Dequeue() (Message, error)

	DequeuePriority(priority int) (Message, error)
}
