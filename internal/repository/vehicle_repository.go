package repository

import (
	"github.com/j0n4t45d3v/parking_management/database"
	"github.com/j0n4t45d3v/parking_management/internal/domain"
)

func FindAllVehicles() ([]domain.Vehicle, error) {
	query := "SELECT " +
		"v.band, " +
		"v.model, " +
		"v.plate, " +
		"v.type, " +
		"v.withdrawn " +
		"FROM vehicles v"

	vehicles := []domain.Vehicle{}

	con, _ := database.GetConnection()
	rows, err := con.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var vehicle domain.Vehicle
		rows.Scan(
			&vehicle.Band,
			&vehicle.Model,
			&vehicle.Plate,
			&vehicle.Type,
			&vehicle.Withdrawn,
		)

		vehicles = append(vehicles, vehicle)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return vehicles, nil
}

func FindByIdVehicle(id string) (domain.Vehicle, error) {
	query := "SELECT " +
		"v.band, " +
		"v.model, " +
		"v.plate, " +
		"v.type, " +
		"FROM vehicles v" +
		"WHERE v.id = $1"

	var vehicle domain.Vehicle

	con, _ := database.GetConnection()
	err := con.QueryRow(query, id).Scan(
		&vehicle.Band,
		&vehicle.Model,
		&vehicle.Plate,
		&vehicle.Type,
	)
  
  if err != nil {
    return domain.Vehicle{}, err
  }

	return vehicle, nil
}

func SaveVehicle(vehicle domain.Vehicle) (string, error) {
	query := "INSERT INTO vehicles (band, model, plate, type) VALUES ($1, $2, $3, $4) RETURNING id"

	con, _ := database.GetConnection()

	var idVehicle string

	err := con.QueryRow(
		query,
		vehicle.Band,
		vehicle.Model,
		vehicle.Plate,
		vehicle.Type,
	).Scan(&idVehicle)

	if err != nil {
		return "", err
	}

	return idVehicle, nil
}

func DeleteVehicle(id string) error {
	query := "DELETE FROM vehicles WHERE id = $1"

	con, _ := database.GetConnection()
	_, err := con.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func UpdateVehicle(id string, vehicleUpdate domain.Vehicle) (domain.Vehicle, error) {
	query := "UPDATE vehicles SET band = $1, model = $2, plate = $3, type = $4 WHERE id = $5"

	con, _ := database.GetConnection()

	_, err := con.Exec(
		query,
		vehicleUpdate.Band,
		vehicleUpdate.Model,
		vehicleUpdate.Plate,
		vehicleUpdate.Type,
		id,
	)

	if err != nil {
		return domain.Vehicle{}, err
	}

	return vehicleUpdate, nil
}
