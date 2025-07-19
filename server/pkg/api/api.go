package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	Rout *mux.Router
}

func (api *API) Endpoints() {
	api.Rout.Use()

	api.Rout.HandleFunc("/api/v1/register", api.register).Methods(http.MethodPost, http.MethodOptions)
	api.Rout.HandleFunc("/api/v1/login", api.login).Methods(http.MethodPost, http.MethodOptions)
	api.Rout.HandleFunc("/api/v1/profile", api.profile).Methods(http.MethodGet, http.MethodOptions)
}
