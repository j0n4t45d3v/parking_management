package repository

import (
	"github.com/j0n4t45d3v/parking_management/database"
	"github.com/j0n4t45d3v/parking_management/internal/domain"
)

func FindAllLogsIoVehicles() ([]domain.LogIOVehicle, error) {
	query := "SELECT " +
		"lv.entry_time, " +
		"lv.departure_time " +
		"FROM logs_io_vehicles lv"

	con, _ := database.GetConnection()
	rows, err := con.Query(query)
	defer rows.Close()

	logsIo := []domain.LogIOVehicle{}

	if err != nil {
		return logsIo, err
	}

	for rows.Next() {
		var logIo domain.LogIOVehicle
		rows.Scan(logIo.EntryTime, logIo.DepartureTime)
		logsIo = append(logsIo, logIo)
	}

	err = rows.Err()

	return logsIo, err
}

func FindByIdLogIoVehicle(id string) (domain.LogIOVehicle, error) {
	query := "SELECT " +
		"lv.entry_time, " +
		"lv.departure_time " +
		"FROM logs_io_vehicles lv" +
		"WHERE lv.id = $1"

	var logIO domain.LogIOVehicle

	con, _ := database.GetConnection()
	err := con.QueryRow(query, id).Scan(&logIO.EntryTime, &logIO.DepartureTime)

	return logIO, err
}

func SaveLogIOVehicle(idVehicle string) (int, error) {
	query := "INSERT INTO logs_io_vehicles (entry_time, id_vehicle)" +
		" VALUES (CURRENT_TIMESTAMP, $1) RETURNING id"

	con, _ := database.GetConnection()
	var idLogIO int
	err := con.QueryRow(query).Scan(idLogIO)

	return idLogIO, err
}

func DeleteLogIOVehicle(id string) error {
	query := "DELETE FROM logs_io_vehicles WHERE id = $1"

	con, _ := database.GetConnection()
	_, err := con.Exec(query, id)
	return err
}

func DepartureVehicle(idVehicle string) error {
	query := "UPDATE logs_io_vehicles SET departure_time = CURRENT_TIMESTAMP WHERE id_vehicle = $1"

	con, _ := database.GetConnection()
	_, err := con.Exec(query, idVehicle)
	return err
}
