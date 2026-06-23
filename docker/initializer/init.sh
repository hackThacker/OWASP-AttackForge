#!/bin/sh
# ==============================================================================
# OWASP Lab Toolkit - Database & App Initializer
# Wait for containers to be up and runs database setup scripts programmatically
# ==============================================================================

set -e

# Helper function to wait for an endpoint to become reachable
wait_for_url() {
  local url="$1"
  local app_name="$2"
  echo "Waiting for ${app_name} on ${url}..."
  until curl -s -o /dev/null -w "%{http_code}" "$url" --connect-timeout 2 | grep -qE "^(200|301|302|401|403|404)$"; do
    sleep 2
  done
  echo "  [+] ${app_name} is up and running!"
}

echo "=== Starting OWASP AttackForge Databases Provisioner ==="

# Wait for all target applications to respond to basic HTTP requests
wait_for_url "http://mutillidae:8080/index.php" "Mutillidae II"
wait_for_url "http://dvwa:8080/setup.php" "DVWA"
wait_for_url "http://bwapp:8080/login.php" "bWAPP"
wait_for_url "http://xvwa:8080/index.php" "XVWA"
wait_for_url "http://vwa:8080/index.php" "VWA"

echo "=== Initializing Application Databases ==="

# 1. Initialize Mutillidae II
echo "[+] Initializing Mutillidae II database..."
curl -s -o /dev/null "http://mutillidae:8080/set-up-database.php"

# 2. Initialize DVWA
echo "[+] Initializing DVWA database..."
# DVWA uses CSRF token protection on its setup page, so we fetch it first
DVWA_TOKEN=$(curl -s -c /tmp/cookies.txt "http://dvwa:8080/setup.php" | grep "user_token" | awk -F"value='" '{print $2}' | awk -F"'" '{print $1}')
curl -s -X POST -b /tmp/cookies.txt -d "create_db=Create+%2F+Reset+Database&user_token=${DVWA_TOKEN}" "http://dvwa:8080/setup.php" > /dev/null

# 3. Initialize bWAPP
echo "[+] Initializing bWAPP database..."
curl -s -o /dev/null "http://bwapp:8080/install.php?install=yes"

# 4. Initialize XVWA
echo "[+] Initializing XVWA database..."
curl -s -o /dev/null "http://xvwa:8080/setup/?action=do"

# 5. Initialize VWA
echo "[+] Initializing VWA database..."
curl -s -X POST -d "submit=Enter" "http://vwa:8080/index.php" > /dev/null

echo "=== All Application Databases Provisioned Successfully ==="
exit 0
