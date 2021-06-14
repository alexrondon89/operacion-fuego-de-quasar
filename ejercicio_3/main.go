package main

import (
	"ejercicio/ejercicio_3/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Creating Handler...")
	exec := handler.Execution{}

	if os.Getenv("local") == "true"{
		input := "{\"distance\": 1000.0," +
			"\"message\": [\"\", \"este\", \"es\", \"un\", \"mensaje\"]" +
			"}"

		queryInput := map[string]string{"satellite_name": "kenobi"}
		request := events.APIGatewayProxyRequest{Body: input, PathParameters: queryInput, HTTPMethod: "POST"}
		exec.HandlerRequest(nil, request)
	}else{
		lambda.Start(exec.HandlerRequest)
	}
}