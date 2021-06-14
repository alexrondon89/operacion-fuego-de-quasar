package tests

import (
	"ejercicio/commons/dto"
	"ejercicio/commons/models"
	"ejercicio/commons/services"
	"ejercicio/commons/tests"
	"ejercicio/ejercicio_3/actions"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"math"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

// Test happy path, donde la salida del action sea un byte con la informacion
//calculada por la interfaz CalculatorInterface. Se mockea el llamado a dynamo en el meotodo ScanItems
func Test_HappyPathGetAction(t *testing.T) {
	mockDBService := new(tests.MockDynamoDBService)
	getHand := actions.GetAction{DbService: mockDBService, MessageInterceptor: &services.GenericInterceptor{Calculator: &services.GenericCalculator{}}}
	request := events.APIGatewayProxyRequest{HTTPMethod: "GET"}

	satelliteModel := &models.Satellite{Name: "kenobi", Distance: 1000.0, Message: []string{"", "este", "es", "un", "mensaje"}}
	expectedModel := dto.NewEnemyInformationResponseObject()
	expectedModel.Message =" este es un mensaje"
	expectedModel.Position.X = -135.8899
	expectedModel.Position.Y = 871.7798

	listModel := []models.Satellite{*satelliteModel}
	mockDBService.On("ScanItems").Return(&listModel, nil)

	byteExpected, _ := json.Marshal(expectedModel)
	response, err := getHand.ExecuteAction(request)
	assert.Equal(t, byteExpected, response)
	assert.Nil(t, err)
}

//Test negative path, donde se verifica que se reciba un error al no encontrarse niguna nave
// registrada en base de datos. Se mockea el llamado a dynamo para el metodo ScanItems
func TestNegativePathScanItemsEmpty(t *testing.T) {
	mockDBService := new(tests.MockDynamoDBService)
	getHand := actions.GetAction{DbService: mockDBService, MessageInterceptor: &services.GenericInterceptor{Calculator: &services.GenericCalculator{}}}

	request := events.APIGatewayProxyRequest{HTTPMethod: "GET"}
	expected := "there are no satellites with enemy information sent "
	mockDBService.On("ScanItems").Return(nil, errors.New(expected))

	_, err := getHand.ExecuteAction(request)
	assert.Equal(t, expected, err.Error())
	assert.NotNil(t, err)
}

//Test negative path, donde se verifica que se reciba un error al no poderse interceptar el mensaje enemigo.
//Se mockea el llamado la interfaz genericInterceptor para los metodos GetLocation y GetMessage
func TestNegativePathEnemyInformationCanNotBeIntercepted(t *testing.T) {
	mockDBService := new(tests.MockDynamoDBService)
	mockInterceptorService := new(tests.MockInterceptorService)
	getHand := actions.GetAction{DbService: mockDBService, MessageInterceptor: mockInterceptorService}

	queryInput := map[string]string{"satellite_name": "kenobi"}
	request := events.APIGatewayProxyRequest{PathParameters: queryInput, HTTPMethod: "GET"}
	satelliteModel := &models.Satellite{Name: "kenobi", Distance: 1000.0, Message: []string{"", " ", " ", " ", " "}}
	listModel := []models.Satellite{*satelliteModel}

	expected := "It's not possible to get enemy information "
	mockDBService.On("GetItem", "kenobi").Return(satelliteModel, nil)
	mockDBService.On("ScanItems").Return(&listModel, nil)
	mockInterceptorService.On("GetLocation", mock.Anything).Return(float32(math.NaN()), float32(math.NaN()))
	mockInterceptorService.On("GetMessage", mock.Anything).Return("")

	_, err := getHand.ExecuteAction(request)
	assert.Equal(t, expected, err.Error())
	assert.NotNil(t, err)
}
