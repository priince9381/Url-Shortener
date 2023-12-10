package router

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/priince938/app/internal/controller"
)

func Router(ctx context.Context) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", controller.CreateShortURL).Methods("POST")
	router.HandleFunc("/get_url/{short_url}", controller.RedirectURL).Methods("GET")
	return router
}
