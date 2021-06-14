package tests

import (
	"ejercicio/commons/models"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDBService struct{ mock.Mock }

func (m *MockDynamoDBService) GetItem(satelliteName string) (*models.Satellite, error){
	args := m.Called(satelliteName)
	if len(m.ExpectedCalls) == 0 {
		return nil, fmt.Errorf("Not DEFINED expected calls. ")
	}
	if args.Get(0) == nil {
		return &models.Satellite{}, args.Error(1)
	}
	return args.Get(0).(*models.Satellite), args.Error(1)
}

func (m *MockDynamoDBService) UpdateItem(inputSplitRequest *models.Satellite) (*models.Satellite, error){
	args := m.Called(inputSplitRequest)
	if len(m.ExpectedCalls) == 0 {
		return nil, fmt.Errorf("Not DEFINED expected calls. ")
	}
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Satellite), args.Error(1)
}

func (m *MockDynamoDBService) ScanItems() (*[]models.Satellite, error){
	args := m.Called()
	if len(m.ExpectedCalls) == 0 {
		return nil, fmt.Errorf("Not DEFINED expected calls. ")
	}
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]models.Satellite), args.Error(1)
}
