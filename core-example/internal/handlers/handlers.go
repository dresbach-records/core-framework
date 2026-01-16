package handlers

import (
	"net/http"
	"core-example/internal/state"
	"core-example/internal/view"
)

func Home(w http.ResponseWriter, r *http.Request) {
	s := state.New()
	s["title"] = "Bem-vindo ao CORE"
	s["message"] = "Framework Server-First"

	view.Render(w, "home", "public", s)
}

func Services(w http.ResponseWriter, r *http.Request) {
	s := state.New()
	s["title"] = "Serviços"
	s["services"] = []string{"Websites", "APIs", "Consultoria"}

	view.Render(w, "services", "public", s)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Em um app real, aqui você faria a autenticação
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	view.Render(w, "login", "public", nil)
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	s := state.New()
	s["user"] = "Cliente CORE"
	s["services"] = []string{
		"Hosting Ativo",
		"Backup Diário",
		"SSL Válido",
	}

	view.Render(w, "dashboard", "dashboard", s)
}
