package service

import (
	"github.com/j0n4t45d3v/parking_management/internal/domain"
	"github.com/j0n4t45d3v/parking_management/internal/repository"
)

func GetAllEstablishments() ([]domain.Establishment, error) {

	establishments, err := repository.FindAllEstablishment()

	if err != nil {
		return nil, err
	}
   
	return establishments, nil
}
