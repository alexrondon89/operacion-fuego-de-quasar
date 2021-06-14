package utils

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func OutputIsEmpty(response *dynamodb.QueryOutput) bool{
	return len(response.Items) == 0 || response.Items[0] == nil
}

func OutputListIsEmpty(response *dynamodb.ScanOutput) bool{
	return len(response.Items) == 0
}

func GenerateModel(input map[string]*dynamodb.AttributeValue, model interface{}) error{
	if err := dynamodbattribute.UnmarshalMap(input, model); err != nil {
		return err
	}
	return nil
}