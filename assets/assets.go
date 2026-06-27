package assets

import (
	"embed"
	"net/http"
)

//go:embed all:templates
var Templates embed.FS

//go:embed all:css all:js
var Static embed.FS

// StaticHandler returns an HTTP handler for /static/ files.
func StaticHandler() http.Handler {
	// The embedded FS has paths like "css/style.css" and "js/htmx.min.js"
	// We serve them under /static/css/ and /static/js/
	return http.StripPrefix("/static/", http.FileServer(http.FS(Static)))
}