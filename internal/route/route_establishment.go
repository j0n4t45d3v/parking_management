package route

import (
	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/handler"
)

func SetRouteV1Establishment(r *mux.Router) {
  establishmentRoute := r.PathPrefix("/v1/establishment").Subrouter()
  establishmentRoute.HandleFunc("", handler.ListEstablisments).Methods("GET")
  establishmentRoute.HandleFunc("", handler.RegisterEstablisment).Methods("POST")
  establishmentRoute.HandleFunc("/{id}", handler.DeleteEstablisments).Methods("DELETE")
} 
