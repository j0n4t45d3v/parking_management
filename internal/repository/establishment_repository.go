package repository

import (
	"github.com/j0n4t45d3v/parking_management/database"
	"github.com/j0n4t45d3v/parking_management/internal/domain"
)

func FindAllEstablishment() ([]domain.Establishment, error) {

	con, _ := database.GetConnection()

	query := " SELECT " +
		" e.name, " +
		" e.document, " +
		" e.phone, " +
		" e.qtde_motoclycles " +
		" e.qtde_cars " +
		" a.city " +
		" a.street " +
		" a.number " +
		" a.state " +
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
	}

	if err := rows.Err(); err != nil {
		return establishments, err
	}

	return establishments, nil
}
