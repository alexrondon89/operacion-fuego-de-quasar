package tests

import (
	"ejercicio/commons/dto"
	"ejercicio/commons/models"
	"ejercicio/commons/tests"
	"ejercicio/ejercicio_3/actions"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)


//Test happy path, donde la salida del action sea un byte con un mensaje de confirmacion luego de guardarse correctamente
// la informacion interceptada. Se mockea el acceso a dynamo para la interfaz genericDynamo y el metodo UpdateItem
func Test_HappyPathPostAction(t *testing.T) {
	mockDBService := new(tests.MockDynamoDBService)
	postHand := actions.PostAction{DbService: mockDBService}

	bodyInput := "{\"distance\": 1000.0," +
		"\"message\": [\"\", \"este\", \"es\", \"un\", \"mensaje\"]" +
		"}"
	queryInput := map[string]string{"satellite_name": "kenobi"}
	request := events.APIGatewayProxyRequest{Body: bodyInput, PathParameters: queryInput, HTTPMethod: "POST"}

	expectedModel := dto.NewInformationAddedResponseObject()
	expectedModel.Message = "satellite kenobi, your information was saved successfully"
	satelliteModel := &models.Satellite{Name: "kenobi", Distance: 1000.0, Message: []string{"", "este", "es", "un", "mensaje"}}
	mockDBService.On("UpdateItem", satelliteModel).Return(satelliteModel, nil)

	byteExpected, _ := json.Marshal(expectedModel)

	response, err := postHand.ExecuteAction(request)
	assert.Equal(t, byteExpected, response)
	assert.Nil(t, err)
}

//Test negative path, se verifica que el action devuelva un error luego de ocurrir un problema en el metodo
// UpdateItem contra dynamo. Se mockea el acceso a dynamo para el metodo UpdateItem
func Test_HegativePathErrorInMethodUpdateItem(t *testing.T) {
	mockDBService := new(tests.MockDynamoDBService)
	postHand := actions.PostAction{DbService: mockDBService}

	bodyInput := "{\"distance\": 1000.0," +
		"\"message\": [\"\", \"este\", \"es\", \"un\", \"mensaje\"]" +
		"}"
	queryInput := map[string]string{"satellite_name": "kenobi"}
	request := events.APIGatewayProxyRequest{Body: bodyInput, PathParameters: queryInput, HTTPMethod: "POST"}

	message := "One error occurred during update execution!: "
	mockDBService.On("UpdateItem", mock.Anything).Return(nil, errors.New(message))
	_, err := postHand.ExecuteAction(request)
	assert.Equal(t, message, err.Error())
	assert.NotNil(t, err)
}

