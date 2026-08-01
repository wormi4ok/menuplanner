#!/bin/sh
set -eu

cat > /usr/share/nginx/html/config.js <<EOF
window.config = {
  MP_VERSION: '${MP_VERSION:-dev}',
  MP_CLIENT_ID: '${MP_CLIENT_ID:-}',
  API_ADDRESS: '${API_ADDRESS:-http://localhost:8081}',
};
EOF
