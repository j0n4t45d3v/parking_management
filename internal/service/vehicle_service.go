package service

import (
	"github.com/j0n4t45d3v/parking_management/internal/domain"
	"github.com/j0n4t45d3v/parking_management/internal/repository"
)

func GetAllVehicles() ([]domain.Vehicle, int16, error) {

	vehicles, err := repository.FindAllVehicles()

	if err != nil {
		return []domain.Vehicle{}, 500, err
	}

	return vehicles, 200, nil
}

func GetOneById(id string) (domain.Vehicle, int16, error) {
	vehicle, err := repository.FindByIdVehicle(id)

	if err != nil {
		return domain.Vehicle{}, 500, err
	}

	return vehicle, 200, nil
}

func RegisterVehicle(vehicle domain.Vehicle) (string, int16, error) {
	idVehicle, err := repository.SaveVehicle(vehicle)

	if err != nil {
		return "", 500, err
	}

	return idVehicle, 201, nil
}

func DeleteVehicle(id string) (int16, error) {
	err := repository.DeleteVehicle(id)

	if err != nil {
		return 500, err
	}

	return 200, nil
}

func UpdateVehicle(vehicle domain.Vehicle, id string) (domain.Vehicle, int16, error) {
  result, err := repository.UpdateVehicle(id, vehicle)

	if err != nil {
		return domain.Vehicle{}, 500, err
	}

	return result,200, nil
}
