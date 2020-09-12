package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Env struct {
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes
var env Env

func InitRouter() {
	routes = Routes{
		Route{
			"index",
			"GET",
			"/traffics",
			env.index,
		},
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func StartWebServer(port string) {
	log.Println("Starting HTTP service at " + port)
	r := NewRouter()
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

func (env *Env) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode("key:test")
}
