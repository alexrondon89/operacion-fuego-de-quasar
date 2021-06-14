package dto

import (
	"ejercicio/commons/models"
)

type enemyInformationResponse struct {
	Position models.Coordinates `json:"position"`
	Message  string             `json:"message"`
}

type informationAddedResponse struct {
	Message string	`json:"message"`
}

func NewEnemyInformationResponseObject() *enemyInformationResponse {
	return &enemyInformationResponse{}
}

func NewInformationAddedResponseObject () *informationAddedResponse {
	return &informationAddedResponse{}
}

