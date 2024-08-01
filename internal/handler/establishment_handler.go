package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/common"
	"github.com/j0n4t45d3v/parking_management/internal/domain"
	"github.com/j0n4t45d3v/parking_management/internal/service"
)

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

func DeleteEstablisments(res http.ResponseWriter, req *http.Request) {
  idEstablishment := mux.Vars(req)["id"]
	status, err := service.DeleteEstablishment(idEstablishment)
	if err != nil {
		errorResponse := common.ToJsonError(int(status), err.Error())
		fmt.Fprint(res, errorResponse)
    return
	}
	response := common.ToJsonSucessString(int(status), "Establisment Deleted!")
	fmt.Fprint(res, response)
}
