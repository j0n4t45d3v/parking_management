package service

import (
	"github.com/j0n4t45d3v/parking_management/internal/domain"
	"github.com/j0n4t45d3v/parking_management/internal/repository"
)

func GetAllLogs() ([]domain.LogIOVehicle, int16, error) {

	result, err := repository.FindAllLogsIoVehicles()

	if err != nil {
		return []domain.LogIOVehicle{}, 500, err
	}

	return result, 200, nil
}

func GetByIdLog(id string) (domain.LogIOVehicle, int16, error) {
	result, err := repository.FindByIdLogIoVehicle(id)

	if err != nil {
		return domain.LogIOVehicle{}, 500, err
	}

	return result, 200, nil
}
