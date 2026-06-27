package template

import (
	"html/template"
	"net/http"

	"github.com/jlcoulter/go-dashboard-template/assets"
)

// Data is a convenience type for template data.
type Data map[string]interface{}

// Title is a helper to set the page title in Data.
func Title(title string) Data {
	return Data{"Title": title}
}

// Templates holds parsed Go templates.
type Templates struct {
	templates *template.Template
}

// New loads and parses all templates from the embedded filesystem.
func New() (*Templates, error) {
	tmpl := template.New("")

	// Register custom template functions
	tmpl = tmpl.Funcs(template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	// Parse all .html files from assets/templates/
	tmpl, err := tmpl.ParseFS(assets.Templates, "templates/*.html", "templates/partials/*.html")
	if err != nil {
		return nil, err
	}

	return &Templates{templates: tmpl}, nil
}

// RenderPage renders a full HTML page to w.
func (t *Templates) RenderPage(w http.ResponseWriter, name string, data Data) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return t.templates.ExecuteTemplate(w, name, data)
}

// RenderPartial renders an HTML partial (for HTMX responses).
func (t *Templates) RenderPartial(w http.ResponseWriter, name string, data Data) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return t.templates.ExecuteTemplate(w, name, data)
}