package main

import (
	"ejercicio/commons/services"
	"ejercicio/ejercicio_2/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Creating Handler...")
	exec := handler.Execution{MessageInterceptor: &services.GenericInterceptor{Calculator: &services.GenericCalculator{}}}

	if os.Getenv("local") == "true"{
		input := "{\"satellites\": [" +
			"{\"name\": \"kenobi\", " +
			"\"distance\": 10.0," +
			"\"message\": [\"\", \"\", \"\", \"\", \"\"]}," +
			"{\"name\": \"skywalker\", " +
			"\"distance\": 1105.5," +
			"\"message\": [\"\", \"\", \"\", \"\", \"\"]}," +
			"{\"name\": \"sato\", " +
			"\"distance\": 1402.7," +
			"\"message\": [\"\", \"\", \"\", \"\", \"\"]} ]}"

		request := events.APIGatewayProxyRequest{Body: input, HTTPMethod: "POST"}
		exec.HandlerRequest(nil, request)
	}else{
		lambda.Start(exec.HandlerRequest)
	}
}