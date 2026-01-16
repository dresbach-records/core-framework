package main

import (
	"log"
	"net/http"

	"core-example/internal/router"
)

func main() {
	// router.New() retorna o mux
	mux := router.New()

	// Configura o servidor de arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("CORE example running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
