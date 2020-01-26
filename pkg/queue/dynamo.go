package queue

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

var _ PriorityQueue = (*dynamoPriorityQueue)(nil)

type dynamoPriorityQueue struct {
	client dynamodbiface.DynamoDBAPI
}

func NewDynamoPriorityQueue() (*dynamoPriorityQueue, error) {
	return &dynamoPriorityQueue{}, nil
}

func (d *dynamoPriorityQueue) Enqueue(message Message) error {
	panic("implement me")
}

func (d *dynamoPriorityQueue) Dequeue() (Message, error) {
	panic("implement me")
}

func (d *dynamoPriorityQueue) DequeuePriority(priority int) (Message, error) {
	panic("implement me")
}
