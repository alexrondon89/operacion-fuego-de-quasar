package services

import (
	"ejercicio/commons/builders"
	"ejercicio/commons/models"
	"ejercicio/commons/utils"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/sirupsen/logrus"
)

type DynamoService struct{
}

func (ds *DynamoService) GetItem(satelliteName string) (*models.Satellite, error){
	logrus.Info("Check if exists message from satellite: ", satelliteName)
	cond1 := expression.Key("Satellite").Equal(expression.Value(satelliteName))
	expr, err := expression.NewBuilder().WithKeyCondition(cond1).Build()
	if utils.CheckError(err) {
		return nil, err
	}

	response, err := builders.NewDynamoQuery().
		KeyConditionExpression(expr.KeyCondition()).
		ExpressionAttributeNames(expr.Names()).
		ExpressionAttributeValues(expr.Values()).
		FilterExpression(expr.Filter()).
		TableName(aws.String(utils.TableName)).
		Execute()

	if err != nil {
		logrus.Info(utils.QueryDynamoError, err)
		return nil, err
	}

	if utils.OutputIsEmpty(response){
		logrus.Info(satelliteName, " has not sent any message before")
		return nil, errors.New(utils.UnknownSatellite)
	}

	model := new(models.Satellite)
	err = utils.GenerateModel(response.Items[0], model)

	return model, err
}

func (ds *DynamoService) UpdateItem(inputSplitRequest *models.Satellite) (*models.Satellite, error) {
	message := inputSplitRequest.Message
	distance := inputSplitRequest.Distance
	name := inputSplitRequest.Name

	logrus.Info("updating satellite information for satellite:", name)
	messageList := generateDynamoAttributeList(message)

	updateExpression := "SET Message = :mess, Distance = :dist"
	valuesExpression := map[string]*dynamodb.AttributeValue{
		":mess": {L: messageList},
		":dist":{N: aws.String(fmt.Sprint(distance))},
	}

	key:= map[string]*dynamodb.AttributeValue{"Satellite": {S: aws.String(name)}}

	response, err := builders.NewDynamoUpdate().
		Key(key).
		TableName(aws.String(utils.TableName)).
		UpdateExpression(aws.String(updateExpression)).
		ExpressionAttributeValues(valuesExpression).
		ReturnValues(aws.String("ALL_NEW")).
		Execute()

	if utils.CheckError(err) {
		logrus.Info(utils.UpdateDynamoError, err)
		return nil, err
	}

	model := new(models.Satellite)
	err = utils.GenerateModel(response.Attributes, model)
	return model, err
}

func (ds *DynamoService) ScanItems() (*[]models.Satellite, error) {

	proj := expression.NamesList(expression.Name("Message"), expression.Name("Distance"), expression.Name("Satellite"))
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if utils.CheckError(err) {
		return nil, err
	}

	response, err := builders.NewDynamoScan().
		TableName(aws.String(utils.TableName)).
		ExpressionAttributeValues(expr.Values()).
		ExpressionAttributeNames(expr.Names()).
		ProjectionExpression(expr.Projection()).
		Execute()

	if utils.CheckError(err) {
		logrus.Info(utils.QueryDynamoError, err)
		return nil, err
	}

	if utils.OutputListIsEmpty(response){
		logrus.Info(utils.EmptyDataBase)
		return nil, errors.New(utils.EmptyDataBase)
	}

	list, err := generateListModel(response)
	return list, err
}

func generateDynamoAttributeList(messages []string) []*dynamodb.AttributeValue{
	var list []*dynamodb.AttributeValue
	for _, v := range messages{
		list = append(list, &dynamodb.AttributeValue{S: aws.String(v)})
	}
	return list
}

func generateListModel(input *dynamodb.ScanOutput)(*[]models.Satellite, error){
	var lista []models.Satellite
	for _, v := range input.Items{
		model := new(models.Satellite)
		err := utils.GenerateModel(v, model)

		if utils.CheckError(err){
			return nil, err
		}

		lista = append(lista, *model)
	}

	return &lista, nil
}
