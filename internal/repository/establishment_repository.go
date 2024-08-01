package repository

import (
	"fmt"

	"github.com/j0n4t45d3v/parking_management/database"
	"github.com/j0n4t45d3v/parking_management/internal/domain"
)

func FindAllEstablishment() ([]domain.Establishment, error) {

	con, _ := database.GetConnection()

	query := " SELECT " +
		" e.name, " +
		" e.document, " +
		" e.phone, " +
		" e.qtde_motoclycles, " +
		" e.qtde_cars, " +
		" a.city, " +
		" a.street, " +
		" a.number, " +
		" a.state, " +
		" a.neighborhood " +
		" FROM establishment e " +
		" JOIN addresses a ON e.address_id = a.id"

	rows, err := con.Query(query)
	establishments := []domain.Establishment{}
	if err != nil {
		return establishments, err
	}

	defer rows.Close()
	for rows.Next() {
		var establishment domain.Establishment

		rows.Scan(
			&establishment.Name,
			&establishment.Document,
			&establishment.Phone,
			&establishment.QtdMotorcycles,
			&establishment.QtdCars,
			&establishment.Addrees.City,
			&establishment.Addrees.Street,
			&establishment.Addrees.Number,
			&establishment.Addrees.State,
			&establishment.Addrees.Neighborhood,
		)

		establishments = append(establishments, establishment)
	}

	if err := rows.Err(); err != nil {
		return establishments, err
	}

	return establishments, nil
}

func SaveAddress(address domain.Addrees) (int64, error) {
	query := "INSERT INTO addresses (city, street, number, state, neighborhood) " +
		" VALUES ($1, $2, $3, $4, $5) RETURNING id"

  var idAddress int64 
	con, _ := database.GetConnection()
	err := con.QueryRow(
		query,
		address.City,
		address.Street,
		address.Number,
		address.State,
		address.Neighborhood,
	).Scan(&idAddress)

	if err != nil {
    fmt.Println("Error address", err.Error())
    fmt.Println(query)
		return 0, err
	}

	return idAddress, nil
}

func SaveEstablishment(establishment domain.Establishment, idAddress int64) (int64, error) {

	query := "INSERT INTO establishment (name, document, phone, qtde_motoclycles, qtde_cars, address_id) " +
		" VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"

	con, _ := database.GetConnection()
  var idEstablishment int64
	err := con.QueryRow(
		query,
		establishment.Name,
		establishment.Document,
		establishment.Phone,
		establishment.QtdMotorcycles,
		establishment.QtdCars,
    idAddress,
	).Scan(&idEstablishment)

	if err != nil {
    fmt.Println("Error Establish", err.Error())
		return 0, err
	}

	return idEstablishment, nil
}
