package infrastructure

import (
	"ejercicio/commons/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type client struct {
	dynamoClient 	*dynamodb.DynamoDB
}

type clientBuilder struct {
	client *client
}

func NewClientAws() *clientBuilder{
	return &clientBuilder{client: &client{dynamoClient: &dynamodb.DynamoDB{}}}
}

func (c *clientBuilder) GetDynamoClient() *dynamodb.DynamoDB{
	session, _ := getSession()
	c.client.dynamoClient = dynamodb.New(session, aws.NewConfig().WithRegion(utils.DefaultRegion))
	return c.client.dynamoClient
}
