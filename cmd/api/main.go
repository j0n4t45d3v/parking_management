package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/j0n4t45d3v/parking_management/internal/middlerware"
	"github.com/j0n4t45d3v/parking_management/internal/route"
)

func main() {
  muxRoute := mux.NewRouter()
  muxRoute.Use(middlerware.ContentTypeMiddlerware)
  api := muxRoute.PathPrefix("/api").Subrouter()
  route.SetRouteV1Establishment(api)
  log.Println("Running Server...")
  http.ListenAndServe(":8000", muxRoute)
}
