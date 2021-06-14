package tests

import (
	"ejercicio/commons/services"
	"ejercicio/ejercicio_2/handler"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)
// Test happy path donde se evalua que sea posible calcular la distancia y mensaje del enemigo en base a informacion de entrada.
// El evento de respuesta sera 200 con la ubicacion del enemigo y el mensaje cifrado
func Test_HappyPathHandler(t *testing.T) {
	execution := handler.Execution{MessageInterceptor: &services.GenericInterceptor{Calculator: &services.GenericCalculator{}}}

	input := "{\"satellites\": [" +
		"{\"name\": \"kenobi\", " +
		"\"distance\": 1000.0," +
		"\"message\": [\"este\", \"\", \"\", \"mensaje\", \"\"]}," +
		"{\"name\": \"skywalker\", " +
		"\"distance\": 1105.5," +
		"\"message\": [\"\", \"es\", \"\", \"\", \"secreto\"]}," +
		"{\"name\": \"sato\", " +
		"\"distance\": 1402.7," +
		"\"message\": [\"este\", \"\", \"un\", \"\", \"\"]} ]}"

	request := events.APIGatewayProxyRequest{Body: input, HTTPMethod: "POST"}
	output, _ := handler.PrepareResponse(-213.11671, 1026.2334, "este es un mensaje secreto")
	expected := events.APIGatewayProxyResponse{Body: string(output), StatusCode: http.StatusOK}

	response, err := execution.HandlerRequest(nil, request)
	assert.Equal(t, expected, response)
	assert.Nil(t, err)
}

// Test negative path donde se verifica que se devuelva un mensaje de error con un statusCode 400 cuando el json de entrada
// sea invalido
func Test_NegativePathDueToInvalidJsonInput(t *testing.T) {
	execution := handler.Execution{MessageInterceptor: &services.GenericInterceptor{Calculator: &services.GenericCalculator{}}}

	input := "{aaa}"

	request := events.APIGatewayProxyRequest{Body: input, HTTPMethod: "POST"}
	message := "invalid character 'a' looking for beginning of object key string"
	expected := events.APIGatewayProxyResponse{Body: fmt.Sprintf("message: %s",message), StatusCode: http.StatusBadRequest}

	response, err := execution.HandlerRequest(nil, request)
	assert.Equal(t, expected, response)
	assert.Nil(t, err)
}

// Test negative path donde se verifica que se devuelva un mensaje de error con un statusCode 400 cuando el json de entrada
// sea valido pero este vacio
func Test_NegativePathDueToEmptyJsonInput(t *testing.T) {
	execution := handler.Execution{MessageInterceptor: &services.GenericInterceptor{Calculator: &services.GenericCalculator{}}}

	input := "{}"

	request := events.APIGatewayProxyRequest{Body: input, HTTPMethod: "POST"}
	message := "at least one satellite must be added "
	expected := events.APIGatewayProxyResponse{Body: fmt.Sprintf("message: %s",message), StatusCode: http.StatusBadRequest}

	response, err := execution.HandlerRequest(nil, request)
	assert.Equal(t, expected, response)
	assert.Nil(t, err)
}

// Test negative path donde se verifica que se devuelva un mensaje de error con un statusCode 400 cuando la
// informacion del enemigo no puede ser descifrada
func Test_NegativePathDueToErrorCatchingEnemyInformation(t *testing.T) {
	execution := handler.Execution{MessageInterceptor: &services.GenericInterceptor{Calculator: &services.GenericCalculator{}}}

	input := "{\"satellites\": [" +
		"{\"name\": \"kenobi\", " +
		"\"distance\": 1.0," +
		"\"message\": [\"este\", \"\", \"\", \"mensaje\", \"\"]}," +
		"{\"name\": \"skywalker\", " +
		"\"distance\": 1.5," +
		"\"message\": [\"\", \"es\", \"\", \"\", \"secreto\"]}," +
		"{\"name\": \"sato\", " +
		"\"distance\": 1.7," +
		"\"message\": [\"este\", \"\", \"un\", \"\", \"\"]} ]}"

	request := events.APIGatewayProxyRequest{Body: input, HTTPMethod: "POST"}
	message := "It's not possible to get enemy information "
	expected := events.APIGatewayProxyResponse{Body: fmt.Sprintf("message: %s",message), StatusCode: http.StatusBadRequest}

	response, err := execution.HandlerRequest(nil, request)
	assert.Equal(t, expected, response)
	assert.Nil(t, err)
}