package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(handler *APIHandler) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	var publicRoutes = Routes{
		{
			"ListMetrics",
			"GET",
			"/metrics",
			handler.ListResourcesWithKind,
		},
	}

	// The public route is always accessible
	for _, route := range publicRoutes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	// Handle static files.
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(handler.frontDir))))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
	})

	return c.Handler(router)
}
