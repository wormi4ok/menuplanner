#!/bin/sh
set -eu

web="$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)"
repo="$(dirname "$web")"
image="mcr.microsoft.com/playwright:v$(node -p "require('$web/package.json').devDependencies['@playwright/test'].replace('^','')")-noble"

(cd "$repo" && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o web/tests/.bin/api .)

exec docker run --rm --platform linux/amd64 \
  -v "$repo":/src -w /src/web \
  -v menuplanner-node-modules:/src/web/node_modules \
  -v menuplanner-npm-cache:/root/.npm \
  -e TEST_API_BINARY=/src/web/tests/.bin/api \
  "$image" sh -c 'npm ci && npx playwright test "$@"' -- "$@"
