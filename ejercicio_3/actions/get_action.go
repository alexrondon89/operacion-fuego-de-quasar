package actions

import (
	"ejercicio/commons/dto"
	"ejercicio/commons/interfaces"
	"ejercicio/commons/models"
	"ejercicio/commons/utils"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"math"
)

type GetAction struct {
	DbService	interfaces.DynamoServiceInterface
	MessageInterceptor interfaces.DecodeInterface
}

func (a *GetAction) ExecuteAction(input events.APIGatewayProxyRequest)([]byte , error) {
	response, err := a.DbService.ScanItems()
	if utils.CheckError(err){
		return nil, err
	}

	distances, messages := utils.GenerateListInformation(*response)
	coorX, coorY := a.MessageInterceptor.GetLocation(distances ...)
	message := a.MessageInterceptor.GetMessage(messages...)
	if checkInformation(coorX, coorY, message) {
		return nil, errors.New(utils.InsufficientInformation)
	}

	byteResponse, err := prepareGetResponse(coorX, coorY, message)
	return byteResponse, err
}

func checkInformation(x float32, y float32, message string) bool {
	return math.IsNaN(float64(x)) || math.IsNaN(float64(y)) || len(message) ==0
}

func prepareGetResponse(coorX float32, coorY float32, message string) ([]byte, error){
	response := dto.NewEnemyInformationResponseObject()
	response.Position = models.Coordinates{X: coorX, Y:coorY}
	response.Message = message
	output, err := json.Marshal(response)
	return output, err
}