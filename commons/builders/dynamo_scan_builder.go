package builders

import (
	"ejercicio/commons/infrastructure"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoScan struct {
	 scan	*dynamodb.ScanInput
}

type scanBuilder struct {
	dynamoScan	*dynamoScan
}

func NewDynamoScan() *scanBuilder{
	return &scanBuilder{dynamoScan: &dynamoScan{scan: &dynamodb.ScanInput{}}}
}

func (s *scanBuilder) Execute() (*dynamodb.ScanOutput, error){
	return infrastructure.NewClientAws().GetDynamoClient().Scan(s.dynamoScan.scan)
}

//builder

func (s *scanBuilder) ProjectionExpression(projectionExpression *string) *scanBuilder {
	s.dynamoScan.scan.ProjectionExpression = projectionExpression
	return s
}

func (s *scanBuilder) ExpressionAttributeNames(expressionAttributeNames map[string]*string) *scanBuilder {
	s.dynamoScan.scan.ExpressionAttributeNames = expressionAttributeNames
	return s
}

func (s *scanBuilder) ExpressionAttributeValues(expressionAttributeValues map[string]*dynamodb.AttributeValue) *scanBuilder {
	s.dynamoScan.scan.ExpressionAttributeValues = expressionAttributeValues
	return s
}

func (s *scanBuilder) FilterExpression(filterExpression *string) *scanBuilder {
	s.dynamoScan.scan.FilterExpression = filterExpression
	return s
}

func (s *scanBuilder) TableName(tableName *string) *scanBuilder {
	s.dynamoScan.scan.TableName = tableName
	return s
}
