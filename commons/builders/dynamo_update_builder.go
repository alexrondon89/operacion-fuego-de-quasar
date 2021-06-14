package builders

import (
	"ejercicio/commons/infrastructure"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoUpdate struct {
	 update	*dynamodb.UpdateItemInput
}

type updateBuilder struct {
	dynamoUpdate	*dynamoUpdate
}

func NewDynamoUpdate() *updateBuilder{
	return &updateBuilder{dynamoUpdate: &dynamoUpdate{update: &dynamodb.UpdateItemInput{}}}
}

func (u *updateBuilder) Execute() (*dynamodb.UpdateItemOutput, error){
	return infrastructure.NewClientAws().GetDynamoClient().UpdateItem(u.dynamoUpdate.update)
}

//builder

func (u *updateBuilder) Key (key map[string]*dynamodb.AttributeValue) *updateBuilder{
	u.dynamoUpdate.update.Key = key
	return u
}

func (u *updateBuilder) TableName (tableName *string) *updateBuilder{
	u.dynamoUpdate.update.TableName = tableName
	return u
}

func (u *updateBuilder) UpdateExpression (updateExpression *string) *updateBuilder{
	u.dynamoUpdate.update.UpdateExpression = updateExpression
	return u
}

func (u *updateBuilder) ConditionExpression (conditionExpression *string) *updateBuilder{
	u.dynamoUpdate.update.ConditionExpression = conditionExpression
	return u
}

func (u *updateBuilder) ExpressionAttributeValues (expressionAttributeValues map[string]*dynamodb.AttributeValue) *updateBuilder{
	u.dynamoUpdate.update.ExpressionAttributeValues = expressionAttributeValues
	return u
}

func (u *updateBuilder) ExpressionAttributeNames (expressionAttributeNames map[string]*string) *updateBuilder{
	u.dynamoUpdate.update.ExpressionAttributeNames = expressionAttributeNames
	return u
}

func (u *updateBuilder) ReturnValues (returnValues *string) *updateBuilder{
	u.dynamoUpdate.update.ReturnValues = returnValues
	return u
}
