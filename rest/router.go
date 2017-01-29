package rest

import (
	"net/http"
	"github.com/gorilla/mux"
	. "github.com/jdbellamy/go-mailer/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"ShowMsgs", "GET", "/msgs", ListMessages},
	Route{"SendMsg", "POST", "/msgs", SendMessage},
	Route{"Metrics", "GET", "/metrics", promhttp.Handler().ServeHTTP},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
