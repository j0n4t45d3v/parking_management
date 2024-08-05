package route

import (
	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/handler"
)

func SetRouteV1LogIo(r *mux.Router) {
  logIoVehicleRouter := r.PathPrefix("/v1/log_io_vehicle").Subrouter()
  logIoVehicleRouter.HandleFunc("", handler.GetAllLogsIOVehicle).Methods("GET")
}
