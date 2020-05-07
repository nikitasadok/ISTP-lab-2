package routes

import (
	"CSGORest/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	getRoute := Route{Name: "List", Method: "GET", Pattern: "/matches", HandlerFunc: handlers.ListMatchesHandler}
	postRoute := Route{Name: "Add", Method: "POST", Pattern: "/matches", HandlerFunc: handlers.AddMatchHandler}
	deleteRoute := Route{Name: "Delete", Method: "DELETE", Pattern: "/matches/{id:[0-9]+}", HandlerFunc: handlers.DeleteMatchHandler}
	putRoute := Route{Name: "Update", Method: "PUT", Pattern: "/matches/{id:[0-9]+}", HandlerFunc: handlers.UpdateMatchHandler}
	router.Methods(getRoute.Method).Path(getRoute.Pattern).Name(getRoute.Name).Handler(getRoute.HandlerFunc)
	router.Methods(postRoute.Method).Path(postRoute.Pattern).Name(postRoute.Name).Handler(postRoute.HandlerFunc)
	router.Methods(deleteRoute.Method).Path(deleteRoute.Pattern).Name(deleteRoute.Name).Handler(deleteRoute.HandlerFunc)
	router.Methods(putRoute.Method).Path(putRoute.Pattern).Name(putRoute.Name).HandlerFunc(putRoute.HandlerFunc)

	return router
}
