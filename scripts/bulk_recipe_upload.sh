#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

JSON_FILE=${1:-}
if [[ -z ${JSON_FILE} ]]; then
  echo "usage: $0 [JSON_FILE]"
  exit 1
fi

jq -c '.[]' "${JSON_FILE}" | while read -r RECIPE; do
  curl -s -XPOST http://localhost:8081/recipe -d "${RECIPE}"
done
