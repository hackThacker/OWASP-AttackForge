#!/bin/bash
set -e

# Define defaults if environment variables are not set
LAB_DOMAIN="${DOMAIN:-hackthacker.lab}"
ORG_NAME="${SSL_ORG:-OWASP AttackForge}"

# Ensure SSL directory exists
mkdir -p /etc/nginx/ssl

# Check if certificates already exist (e.g. from volume mounting)
if [ ! -f /etc/nginx/ssl/owasp.key ] || [ ! -f /etc/nginx/ssl/owasp.crt ]; then
  echo "Wildcard SSL certificates not found in /etc/nginx/ssl."
  echo "Generating self-signed certificate for CN=*.${LAB_DOMAIN}..."
  
  openssl req -x509 -nodes -days 3650 -newkey rsa:2048 \
    -keyout /etc/nginx/ssl/owasp.key \
    -out /etc/nginx/ssl/owasp.crt \
    -subj "/C=US/ST=State/L=City/O=${ORG_NAME}/CN=*.${LAB_DOMAIN}"
    
  echo "SSL certificates generated successfully."
else
  echo "Using existing SSL certificates found in /etc/nginx/ssl."
fi

# Hand over execution to the standard Nginx Docker entrypoint script
# which interpolates templates from /etc/nginx/templates/ to /etc/nginx/conf.d/
exec /docker-entrypoint.sh "$@"
