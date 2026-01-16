package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Engine é o nosso motor de templates CORE.
// Ele armazena os templates pré-carregados para renderização rápida.
type Engine struct {
	templates map[string]*template.Template
}

// New cria e inicializa um novo Engine de templates.
func New(uiDir string) (*Engine, error) {
	e := &Engine{
		templates: make(map[string]*template.Template),
	}

	layoutPath := filepath.Join(uiDir, "layouts", "main.layout.core")
	pagesPattern := filepath.Join(uiDir, "pages", "*.page.core")

	pages, err := filepath.Glob(pagesPattern)
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := strings.TrimSuffix(filepath.Base(page), ".page.core")

		ts, err := template.New(name).ParseFiles(layoutPath, page)
		if err != nil {
			log.Printf("Erro ao analisar template para página '%s': %v\n", name, err)
			continue
		}

		e.templates[name] = ts
		log.Printf("CORE Template: Página '%s' carregada.", name)
	}

	return e, nil
}

// Execute renderiza uma página específica.
func (e *Engine) Execute(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	ts, ok := e.templates[name]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := ts.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("Erro ao executar template %s: %v", name, err)
		http.Error(w, "Erro interno do servidor CORE", http.StatusInternalServerError)
	}
}
