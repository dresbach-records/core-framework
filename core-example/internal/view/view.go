package view

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func Render(w http.ResponseWriter, page, layout string, data any) {
	// Define o padrão para encontrar os arquivos de template
	pattern := filepath.Join("ui", "layouts", "*.layout.core")
	files, err := filepath.Glob(pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Adiciona a página específica e os componentes
	pagePath := filepath.Join("ui", "pages", page+".page.core")
	componentsPattern := filepath.Join("ui", "components", "*.comp.core")
	components, _ := filepath.Glob(componentsPattern)

	files = append(files, pagePath)
	files = append(files, components...)

	// Parse e renderiza
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, layout+".layout.core", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
