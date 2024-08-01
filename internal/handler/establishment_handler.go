package handler

import (
	"encoding/json"
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
  var establishments domain.Establishment

  json.NewDecoder(req.Body).Decode(&establishments)

  idEstablishment, status, err := service.SaveEstablishment(establishments)

	if err != nil {
		errorResponse := common.ToJsonError(int(status), err.Error())
		fmt.Fprint(res, errorResponse)
    return
	}
  
  location := common.BuildUriLocation(*req, "v1/establishment", int(idEstablishment))
  res.Header().Add("Content-Type", "application/json")
  res.Header().Add("Location", location)
 
  response := common.ToJsonSucessString(int(status), "Created establishment")
  fmt.Fprint(res, response)
}

func ListEstablisments(res http.ResponseWriter, req *http.Request) {
	establishments, status, err := service.GetAllEstablishments()
	if err != nil {
		errorResponse := common.ToJsonError(int(status), err.Error())
		fmt.Fprint(res, errorResponse)
    return
	}
	response := common.ToJsonSucess(int(status), establishments)
	fmt.Fprint(res, response)
}
