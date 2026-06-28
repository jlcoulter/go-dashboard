package handler

import (
	"net/http"

	"github.com/jlcoulter/go-dashboard-template/assets"
	"github.com/jlcoulter/go-dashboard-template/internal/template"
)

// Healthz returns OK for health checks.
func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

// Readyz returns ready for readiness checks.
func Readyz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ready"}`))
}

// Dashboard renders the main dashboard page.
func Dashboard(tmpl *template.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := template.Title("Dashboard")
		data["Stats"] = map[string]interface{}{
			"TotalRequests": 1234,
			"ActiveUsers":   42,
			"Uptime":        "3d 14h",
		}
		if err := tmpl.RenderPage(w, "layout.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// StatsPartial renders the stats partial for HTMX.
func StatsPartial(tmpl *template.Templates) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := template.Data{
			"Stats": map[string]interface{}{
				"TotalRequests": 1234,
				"ActiveUsers":   42,
				"Uptime":        "3d 14h",
			},
		}
		if err := tmpl.RenderPartial(w, "partials/stats.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// StatsAPI returns stats as JSON.
func StatsAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"total_requests":1234,"active_users":42,"uptime":"3d 14h"}`))
}

// StaticRoutes returns the HTTP handler for static files.
func StaticRoutes() http.Handler {
	return assets.StaticHandler()
}