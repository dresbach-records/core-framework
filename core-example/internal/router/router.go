package router

import (
	"net/http"
	"core-example/internal/handlers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/services", handlers.Services)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/dashboard", handlers.Dashboard)

	return mux
}
