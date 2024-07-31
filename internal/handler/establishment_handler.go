package handler

import (
	"fmt"
	"net/http"

	"github.com/j0n4t45d3v/parking_management/internal/common"
	"github.com/j0n4t45d3v/parking_management/internal/domain"
	"github.com/j0n4t45d3v/parking_management/internal/service"
)

var colection_establishment = []domain.Establishment{
	{
		Name:           "Teste1",
		Phone:          "(81) 9 9999-9999",
		QtdCars:        12,
		QtdMotorcycles: 24,
		Document:       "123.456.789/0123-4",
		Addrees: domain.Addrees{
			City:         "Recife",
			State:        "PE",
			Street:       "Rua cambuca",
			Number:       "142",
			Neighborhood: "Boa Viajem",
		},
	},
	{
		Name:           "Teste1",
		Phone:          "(81) 9 9999-9999",
		QtdCars:        12,
		QtdMotorcycles: 24,
		Document:       "123.456.789/0123-4",
		Addrees: domain.Addrees{
			City:         "Recife",
			State:        "PE",
			Street:       "Rua cambuca",
			Number:       "142",
			Neighborhood: "Boa Viajem",
		},
	},
}

func RegisterEstablisment(res http.ResponseWriter, req *http.Request) {
}

func ListEstablisments(res http.ResponseWriter, req *http.Request) {
	establishments, err := service.GetAllEstablishments()
	if err != nil {
		errorResponse := common.ToJsonError(500, err.Error())
		fmt.Fprint(res, errorResponse)
    return
	}
	response := common.ToJsonSucess(200, establishments)
	fmt.Fprint(res, response)
}
