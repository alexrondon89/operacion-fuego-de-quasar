package interfaces

import (
	"github.com/aws/aws-lambda-go/events"
)

type ActionInterface interface {
	ExecuteAction(input events.APIGatewayProxyRequest) ([]byte, error)
}