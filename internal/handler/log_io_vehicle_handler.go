package handler

import (
	"fmt"
	"net/http"

	"github.com/j0n4t45d3v/parking_management/internal/common"
	"github.com/j0n4t45d3v/parking_management/internal/service"
)

func GetAllLogsIOVehicle(res http.ResponseWriter, req *http.Request) {
	result, status, err := service.GetAllLogs()

	if err != nil {
		errorResponse := common.ToJsonError(int(status), err.Error())
		fmt.Fprint(res, errorResponse)
	}

	response := common.ToJsonSucess(int(status), result)
	fmt.Fprint(res, response)
}

