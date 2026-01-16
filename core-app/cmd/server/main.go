package main

import (
	"log"
	"net/http"
	"strings"

	"core-app/internal/render"
)

func main() {
	// Inicializa o motor de templates CORE, apontando para o diretório 'ui'.
	renderer, err := render.New("../../ui") // Caminho relativo de cmd/server para ui
	if err != nil {
		log.Fatalf("Falha ao inicializar o motor de renderização CORE: %v", err)
	}

	// Define o handler principal que usa o motor de renderização.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Determina o nome da página a partir da URL.
		pageName := strings.Trim(r.URL.Path, "/")
		if pageName == "" {
			pageName = "home"
		}

		// Usa o motor para renderizar a página solicitada.
		renderer.Execute(w, r, pageName, nil) // O 'nil' é para os dados do estado, que ainda não implementamos.
	})

	// Inicia o servidor.
	log.Println("Servidor CORE escutando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Não foi possível iniciar o servidor: %v", err)
	}
}
