package route

import (
	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/handler"
)

func SetRouteV1Vehicle(r *mux.Router) {
	vehicleRoute := r.PathPrefix("/v1/vehicles").Subrouter()
	vehicleRoute.HandleFunc("", handler.GetVehicles).Methods("GET")
	vehicleRoute.HandleFunc("/{id}", handler.GetOneVehicle).Methods("GET")
  vehicleRoute.HandleFunc("", handler.SaveVehicle).Methods("POST")
  vehicleRoute.HandleFunc("{id}", handler.EditVehicle).Methods("PUT")
  vehicleRoute.HandleFunc("{id}", handler.DeleteVehicle).Methods("DELETE")
}
