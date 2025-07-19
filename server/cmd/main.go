package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/server/pkg/api"
)

type server struct {
	api    *api.API
	router *mux.Router
}

func main() {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = &api.API{Rout: srv.router}
	srv.api.Endpoints()

	log.Fatal(http.ListenAndServe(":63000", srv.router))
}
