package queue

type Message struct {
	ID       string
	Priority int8
	Data     map[string]string
}

type PriorityQueue interface {
	Enqueue(message *Message) error

	Dequeue() (*Message, error)

	DequeuePriority(priority int) (*Message, error)
}
