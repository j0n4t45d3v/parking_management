package route

import (
	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/handler"
)

func SetRouteV1Establishment(r *mux.Router) {
  r.HandleFunc("/v1/establishment", handler.ListEstablisments)
} 
