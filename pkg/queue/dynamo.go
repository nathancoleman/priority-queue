package queue

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var _ PriorityQueue = (*dynamoPriorityQueue)(nil)

type dynamoPriorityQueue struct {
	client dynamodbiface.DynamoDBAPI
}

func NewDynamoPriorityQueue(region string) (*dynamoPriorityQueue, error) {
	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &dynamoPriorityQueue{
		client: dynamodb.New(sess),
	}, nil
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
