package handler

import (
	"context"
	"ejercicio/commons/interfaces"
	"ejercicio/commons/services"
	"ejercicio/commons/utils"
	"ejercicio/ejercicio_3/actions"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Execution struct {
}

func (exec *Execution) getAction(input events.APIGatewayProxyRequest) interfaces.ActionInterface{
	if strings.EqualFold(input.HTTPMethod, "POST"){
		logrus.Info("Post action recoverd")
		return &actions.PostAction{DbService: &services.DynamoService{}}
	}

	logrus.Info("Get action recoverd")
	return &actions.GetAction{
		DbService: &services.DynamoService{},
		MessageInterceptor: &services.GenericInterceptor{
			Calculator: &services.GenericCalculator{},
		},
	}
}

func (exec *Execution) HandlerRequest(_ context.Context, input events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error){
	logrus.Info("Input: ", input)

	action := exec.getAction(input)
	response, err := action.ExecuteAction(input)
	if utils.CheckError(err){
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf(`{"message":"%s"}`, err), StatusCode: http.StatusBadRequest}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusOK}, nil
}
