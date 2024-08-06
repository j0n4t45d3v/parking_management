package service

import (
	"github.com/j0n4t45d3v/parking_management/internal/domain"
	"github.com/j0n4t45d3v/parking_management/internal/repository"
)

func GetAllEstablishments() ([]domain.Establishment, int16, error) {

	establishments, err := repository.FindAllEstablishment()

	if err != nil {
		return nil, 500, err
	}

	return establishments, 200, nil
}

func SaveEstablishment(establishment domain.Establishment) (int64, int16, error) {
  idAddress, status, err := SaveAddress(establishment.Addrees)

  if err != nil {
    return 0, status, err
  }

  idEstablishment, err := repository.SaveEstablishment(establishment, idAddress)

  if err != nil {
    return 0, status, err
  }
	return idEstablishment, 201, nil
}

func SaveAddress(address domain.Addrees) (int64, int16, error) {
  idAddress, err := repository.SaveAddress(address) 

  if err != nil {
    return 0, 500, err
  }

  return idAddress, 201, nil
}

func DeleteEstablishment(id string) (int16, error) {
  err := repository.DeleteEstablishment(id)
  if err != nil {
    return 500, err
  }
  return 200, nil
}

func DeleteAddress(id string) (int16, error) {
  err := repository.DeleteAddress(id)
  if err != nil {
    return 500, err
  }
  return 200, nil
}
