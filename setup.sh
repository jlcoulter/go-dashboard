#!/usr/bin/env bash
# setup.sh — rename module and clean up template markers
# Usage: ./setup.sh mydashboard github.com/you/mydashboard

set -euo pipefail

if [ $# -lt 2 ]; then
  echo "Usage: $0 <dashboard-name> <module-path>"
  echo "Example: $0 healthd github.com/jlcoulter/healthd"
  exit 1
fi

NAME="$1"
MODULE="$2"
OLD_MODULE="github.com/jlcoulter/go-dashboard-template"
OLD_NAME="go-dashboard-template"

# Replace module path in all Go files
find . -name "*.go" -exec sed -i "s|${OLD_MODULE}|${MODULE}|g" {} +

# Replace in go.mod
sed -i "s|${OLD_MODULE}|${MODULE}|g" go.mod

# Replace in Makefile
sed -i "s|${OLD_NAME}|${NAME}|g" Makefile

# Replace in Dockerfile
sed -i "s|${OLD_NAME}|${NAME}|g" Dockerfile

# Replace in README
sed -i "s|${OLD_NAME}|${NAME}|g" README.md

# Replace in .goreleaser.yml
sed -i "s|${OLD_NAME}|${NAME}|g" .goreleaser.yml

# Replace in CI
sed -i "s|${OLD_NAME}|${NAME}|g" .github/workflows/ci.yml

# Remove setup script
rm -- "$0"

echo "Template configured: name=${NAME}, module=${MODULE}"
echo "Next steps:"
echo "  1. Run: go mod tidy"
echo "  2. Run: make test"
echo "  3. Replace static/js/htmx.min.js with the real HTMX library"
echo "  4. Add your dashboard pages and API endpoints"