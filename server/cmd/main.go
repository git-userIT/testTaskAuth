package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mod/server/pkg/api"
)

type server struct {
	api    *api.API
	router *mux.Router
}

func main() {
	var portApp = os.Getenv("APP_PORT")
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = &api.API{Rout: srv.router}
	srv.api.Endpoints()

	log.Fatal(http.ListenAndServe(":"+portApp, srv.router))
}
