package route

import (
	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/handler"
)

func SetRouteV1Establishment(r *mux.Router) {
  baseUri := "/v1/establishment"
  r.HandleFunc(baseUri, handler.ListEstablisments).Methods("GET")
  r.HandleFunc(baseUri, handler.RegisterEstablisment).Methods("POST")
  r.HandleFunc(baseUri, handler.DeleteEstablisments).Methods("DELETE")
} 
