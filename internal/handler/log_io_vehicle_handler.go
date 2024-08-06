package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/common"
	"github.com/j0n4t45d3v/parking_management/internal/service"
)

func GetAllLogsIOVehicle(res http.ResponseWriter, req *http.Request) {
	result, status, err := service.GetAllLogs()

	if err != nil {
		errorResponse := common.ToJsonError(int(status), err.Error())
		fmt.Fprint(res, errorResponse)
		return
	}

	response := common.ToJsonSucess(int(status), result)
	fmt.Fprint(res, response)
}

func GetOneByIdLogIO(res http.ResponseWriter, req *http.Request) {
	idLog := mux.Vars(req)["id"]
	result, status, err := service.GetByIdLog(idLog)

	if err != nil {
		errorResponse := common.ToJsonError(int(status), err.Error())
		fmt.Fprint(res, errorResponse)
		return
	}

	response := common.ToJsonSucess(int(status), result)
	fmt.Fprint(res, response)
}
