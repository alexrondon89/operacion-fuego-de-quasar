package handler

import (
	"context"
	"ejercicio/commons/dto"
	"ejercicio/commons/interfaces"
	"ejercicio/commons/models"
	"ejercicio/commons/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
)

type Execution struct {
	MessageInterceptor interfaces.DecodeInterface
}

func (exec *Execution)HandlerRequest(_ context.Context, input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	logrus.Info("Input: ", input.Body)
	request, err := getInput(input.Body)
	if utils.CheckError(err) {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(`{"message":"%s"}`, err), StatusCode: http.StatusBadRequest}, nil
	}

	err = validateInput(request)
	if utils.CheckError(err) {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(`{"message":"%s"}`, err), StatusCode: http.StatusBadRequest}, nil
	}
	distances, messages := utils.GenerateListInformation(request.Satellites)
	coorX, coorY := exec.MessageInterceptor.GetLocation(distances...)
	message := exec.MessageInterceptor.GetMessage(messages...)
	if checkInformation(coorX, coorY, message) {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(`{"message":"%s"}`, utils.InsufficientInformation), StatusCode: http.StatusBadRequest}, nil
	}

	output, err := PrepareResponse(coorX, coorY, message)
	if utils.CheckError(err) {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(`{"message":"%s"}`, err), StatusCode: http.StatusBadRequest}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(output), StatusCode: http.StatusOK}, nil
}

func getInput(body string) (dto.Input, error){
	request := dto.NewRequestObject()
	err := json.Unmarshal([]byte(body), request)

	if err != nil {
		return nil, err
	}
	logrus.Info("Input unmarshaled: ", request)
	return request, nil
}

func validateInput(input dto.Input) error{
	if len(input.Satellites) == 0{
		return errors.New(utils.EmptySatelliteList)
	}
	return nil
}

func checkInformation(x float32, y float32, message string) bool {
	return math.IsNaN(float64(x)) || math.IsNaN(float64(y)) || len(message) ==0
}

func PrepareResponse(coorX float32, coorY float32, message string)( []byte, error){
	response := dto.NewEnemyInformationResponseObject()
	response.Position = models.Coordinates{X: coorX, Y:coorY}
	response.Message = message
	output, err := json.Marshal(response)
	return output, err
}