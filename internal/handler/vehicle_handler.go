package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/common"
	"github.com/j0n4t45d3v/parking_management/internal/domain"
	"github.com/j0n4t45d3v/parking_management/internal/service"
)

func GetVehicles(res http.ResponseWriter, req *http.Request) {
  result, status,err := service.GetAllVehicles()

  if err != nil {
    errorResponse := common.ToJsonError(int(status), err.Error())
    fmt.Fprint(res, errorResponse)
    return
  }

  response := common.ToJsonSucess(int(status), result)
  fmt.Fprint(res, response)
}

func GetOneVehicle(res http.ResponseWriter, req *http.Request) {
  idVehicle := mux.Vars(req)["id"]
  result, status, err := service.GetOneById(idVehicle)

  if err != nil {
    errorResponse := common.ToJsonError(int(status), err.Error())
    fmt.Fprint(res, errorResponse)
    return
  }

  response := common.ToJsonSucess(int(status), result)
  fmt.Fprint(res, response)
}

func SaveVehicle(res http.ResponseWriter, req *http.Request) {
  var body domain.Vehicle

  err := json.NewDecoder(req.Body).Decode(&body)
  if err != nil {
    errorResponse := common.ToJsonError(500, "Error decode body")
    fmt.Fprint(res, errorResponse)
    return
  }

  idVehicle, status, err:= service.RegisterVehicle(body)
  idVehicleToInt,_ := strconv.Atoi(idVehicle) 

  if err != nil {
    errorResponse := common.ToJsonError(int(status), err.Error())
    fmt.Fprint(res, errorResponse)
    return
  }

  location := common.BuildUriLocation(*req, "v1/vehicle", idVehicleToInt)
  res.Header().Add("Location", location)

  response := common.ToJsonSucessString(int(status), "Vehicle register") 
  fmt.Fprint(res, response)
}

func DeleteVehicle(res http.ResponseWriter, req *http.Request) {
  idVehicle := mux.Vars(req)["id"]
  
  status, err := service.DeleteVehicle(idVehicle)

  if err != nil {
    errorResponse := common.ToJsonError(int(status), err.Error())
    fmt.Fprint(res, errorResponse)
    return
  }

  response := common.ToJsonSucessString(int(status), "Vehicle deleted with sucessfull")
  fmt.Fprint(res, response)
}

func EditVehicle(res http.ResponseWriter, req *http.Request) {
  var body domain.Vehicle
  idVehicle := mux.Vars(req)["id"]
  err := json.NewDecoder(req.Body).Decode(&body)

  if err != nil {
    errorDecode := common.ToJsonError(500, "Error in decode json")
    fmt.Fprint(res, errorDecode)
    return
  }

  _, status,err := service.UpdateVehicle(body, idVehicle)

  if err != nil {
    errorResponse := common.ToJsonError(int(status), err.Error())
    fmt.Fprint(res, errorResponse)
    return
  }

  response := common.ToJsonSucessString(int(status), "Vehicle deleted with sucessfull")
  fmt.Fprint(res, response)
}


