# Go Dashboard Template

A GitHub template for Go web dashboards — server-rendered HTML with HTMX interactivity, embedded templates, and a JSON API side-by-side.

## Why not go-api?

`go-api` returns JSON from every endpoint. Dashboards are different:
- They render HTML pages with `html/template`
- They use HTMX for partial page updates without a JS framework
- They embed static assets (CSS, JS) with `embed.FS`
- They mix server-rendered pages with JSON API endpoints
- They need layout templates, partial templates, and template composition

## Features

- **HTMX** — partial page updates, no SPA framework
- **Server-rendered HTML** with Go `html/template` + layout composition
- **Embedded assets** — CSS/JS compiled into the binary with `embed.FS`
- **JSON API** endpoints alongside HTML pages (same server, `/api/` prefix)
- **Structured logging** with `log/slog`
- **Health endpoints** at `/healthz` and `/readyz`
- **Configuration** via environment variables with viper
- **Graceful shutdown** on SIGINT/SIGTERM
- **Docker** multi-stage build (~10MB)
- **CI** via GitHub Actions
- **Release** via GoReleaser

## Usage

1. Click **"Use this template"** on GitHub to create a new repo
2. Run the setup script:
   ```sh
   ./setup.sh mydashboard github.com/you/mydashboard
   ```
3. Add pages in `internal/handler/` and templates in `templates/`
4. Add API endpoints alongside HTML handlers

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go            # Entry point, router setup, graceful shutdown
├── internal/
│   ├── config/
│   │   └── config.go          # Viper config
│   ├── handler/
│   │   ├── dashboard.go        # HTML dashboard page handlers
│   │   ├── partial.go          # HTMX partial handlers
│   │   └── api.go              # JSON API handlers
│   ├── middleware/
│   │   └── logging.go          # Request logging
│   └── template/
│       └── render.go           # Template loading, layout composition
├── templates/
│   ├── layout.html             # Base layout (head, nav, footer)
│   ├── dashboard.html          # Dashboard page
│   └── partials/
│       └── stats.html          # HTMX partial
├── static/
│   ├── css/
│   │   └── style.css           # Dashboard styles
│   └── js/
│       └── htmx.min.js         # HTMX (embedded)
├── .github/
│   └── workflows/
│       └── ci.yml
├── .goreleaser.yml
├── Dockerfile
├── Makefile
├── go.mod
├── setup.sh
└── README.md
```

## Quick Start

```sh
# Run locally
make run

# Run tests
make test

# Build binary
make build

# Build Docker image
make docker
```

## Adding a New Page

1. Create a handler in `internal/handler/`:

```go
func MyPage(w http.ResponseWriter, r *http.Request) {
    render.Page(w, "mypage.html", render.Data{
        Title: "My Page",
    })
}
```

2. Create a template in `templates/mypage.html` (it extends `layout.html`).

3. Wire the route in `cmd/server/main.go`.

## Adding HTMX Partials

1. Create a partial handler that returns HTML (not JSON):

```go
func MyPartial(w http.ResponseWriter, r *http.Request) {
    render.Partial(w, "partials/mything.html", render.Data{
        "Items": items,
    })
}
```

2. In your template, use HTMX:

```html
<div hx-get="/partials/mything" hx-trigger="load" hx-swap="innerHTML">
  Loading...
</div>
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | HTTP listen port |
| `LOG_LEVEL` | `info` | Log level (debug, info, warn, error) |

## License

MIT