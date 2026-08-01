#!/bin/sh
set -eu

BUILT_VERSION=$(cat /usr/share/nginx/html/version.txt)

cat > /usr/share/nginx/html/config.js <<EOF
window.config = {
  MP_VERSION: '${MP_VERSION:-$BUILT_VERSION}',
  MP_CLIENT_ID: '${MP_CLIENT_ID:-}',
  API_ADDRESS: '${API_ADDRESS:-http://localhost:8081}',
};
EOF
