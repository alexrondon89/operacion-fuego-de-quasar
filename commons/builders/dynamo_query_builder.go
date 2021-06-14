package builders

import (
	"ejercicio/commons/infrastructure"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoQuery struct {
	 query	*dynamodb.QueryInput
}

type queryBuilder struct {
	dynamoQuery	*dynamoQuery
}

func NewDynamoQuery() *queryBuilder{
	return &queryBuilder{dynamoQuery: &dynamoQuery{query: &dynamodb.QueryInput{}}}
}

func (q *queryBuilder) Execute() (*dynamodb.QueryOutput, error){
	return infrastructure.NewClientAws().GetDynamoClient().Query(q.dynamoQuery.query)
}

//builder

func (q *queryBuilder) KeyConditionExpression(keyConditionExpression *string) *queryBuilder {
	q.dynamoQuery.query.KeyConditionExpression = keyConditionExpression
	return q
}

func (q *queryBuilder) ExpressionAttributeNames(expressionAttributeNames map[string]*string) *queryBuilder {
	q.dynamoQuery.query.ExpressionAttributeNames = expressionAttributeNames
	return q
}

func (q *queryBuilder) ExpressionAttributeValues(expressionAttributeValues map[string]*dynamodb.AttributeValue) *queryBuilder {
	q.dynamoQuery.query.ExpressionAttributeValues = expressionAttributeValues
	return q
}

func (q *queryBuilder) FilterExpression(filterExpression *string) *queryBuilder {
	q.dynamoQuery.query.FilterExpression = filterExpression
	return q
}

func (q *queryBuilder) TableName(tableName *string) *queryBuilder {
	q.dynamoQuery.query.TableName = tableName
	return q
}
