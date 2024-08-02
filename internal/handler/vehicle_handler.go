package handler

import (
	"fmt"
	"net/http"

	"github.com/j0n4t45d3v/parking_management/internal/common"
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
