package queue

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var _ PriorityQueue = (*dynamoPriorityQueue)(nil)

type dynamoPriorityQueue struct {
	client    dynamodbiface.DynamoDBAPI
	tableName string
}

type dynamoItem struct {
	ID                string            `dynamodbav:"id"`
	Priority          int8              `dynamodbav:"Priority"`
	EnqueuedTimestamp int64             `dynamodbav:"enqueued_timestamp"`
	Data              map[string]string `dynamodbav:"Data"`
}

func NewDynamoPriorityQueue(tableName, region string) (*dynamoPriorityQueue, error) {
	sess, err := session.NewSession(aws.NewConfig().WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &dynamoPriorityQueue{
		client:    dynamodb.New(sess),
		tableName: tableName,
	}, nil
}

func (d *dynamoPriorityQueue) Enqueue(message *Message) error {
	item := &dynamoItem{
		ID:                message.ID,
		Priority:          message.Priority,
		EnqueuedTimestamp: time.Now().Unix(),
		Data:              message.Data,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	request := &dynamodb.PutItemInput{
		TableName: aws.String(d.tableName),
		Item:      av,
	}

	_, err = d.client.PutItem(request)
	return err
}

func (d *dynamoPriorityQueue) Dequeue() (*Message, error) {
	panic("implement me")
}

func (d *dynamoPriorityQueue) DequeuePriority(priority int) (*Message, error) {
	panic("implement me")
}
