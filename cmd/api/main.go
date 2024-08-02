package main

import (
	"net/http"
  "log"
	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/route"
)

func main() {
  muxRoute := mux.NewRouter()
  api := muxRoute.PathPrefix("/api").Subrouter()
  route.SetRouteV1Establishment(api)
  log.Println("Running Server...")
  http.ListenAndServe(":8000", muxRoute)
}
