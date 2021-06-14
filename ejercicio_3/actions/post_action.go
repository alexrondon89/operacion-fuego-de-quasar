package actions

import (
	"ejercicio/commons/dto"
	"ejercicio/commons/interfaces"
	"ejercicio/commons/models"
	"ejercicio/commons/utils"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

type PostAction struct {
	DbService	interfaces.DynamoServiceInterface
}

func (a *PostAction) ExecuteAction(input events.APIGatewayProxyRequest) ([]byte , error) {
	object, err := getPostInput(input)
	if utils.CheckError(err) {
		return nil, err
	}

	if err = validateInput(object); err!=nil{
		return nil, err
	}

	model, err := a.DbService.UpdateItem(object)
	if utils.CheckError(err) {
		return nil, err
	}

	byteResponse, err := preparePostResponse(model)
	logrus.Info("satellite updated in database: ", model)
	return byteResponse, nil
}

func getPostInput(input events.APIGatewayProxyRequest) (*models.Satellite, error){
	request := dto.NewSplitRequestObject()
	err := json.Unmarshal([]byte(input.Body), request)
	request.Name = input.PathParameters["satellite_name"]

	if utils.CheckError(err) {
		return nil, err
	}
	logrus.Info("Input unmarshaled: ", request)
	return request, nil
}

func validateInput(input *models.Satellite) error{
	if input.Distance <= 0.0 || len(input.Message) == 0 {
		return errors.New(utils.InvalidInformation)
	}

	return nil
}

func preparePostResponse(satellite *models.Satellite) ([]byte, error){
	response := dto.NewInformationAddedResponseObject()
	response.Message = "satellite " +satellite.Name + ", your information was saved successfully"
	output, err := json.Marshal(response)
	return output, err
}