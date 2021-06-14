package dto

import (
	"ejercicio/commons/models"
)

type Input *inputRequest

type inputRequest struct {
	Satellites []models.Satellite
}

func NewRequestObject() *inputRequest {
	return &inputRequest{}
}

func NewSplitRequestObject() *models.Satellite{
	return &models.Satellite{}
}
