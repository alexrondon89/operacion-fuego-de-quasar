package interfaces

import "ejercicio/commons/models"

type DynamoServiceInterface interface {
	GetItem(satelliteName string) (*models.Satellite, error)
	UpdateItem(inputSplitRequest *models.Satellite) (*models.Satellite, error)
	ScanItems() (*[]models.Satellite, error)
}
