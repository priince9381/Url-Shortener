package router

import (
	"github.com/gorilla/mux"
	"github.com/priince938/app/internal/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shorten", controller.CreateShortURL).Methods("POST")
	return router
}
