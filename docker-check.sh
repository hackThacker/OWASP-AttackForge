#!/bin/bash
# ==============================================================================
# OWASP ATTACKFORGE - DOCKER TELEMETRY & HEALTH CHECKER
# ==============================================================================

set -e

# Terminal Colors
WHT='\033[1;37m'
RED='\033[1;31m'
GRN='\033[1;32m'
YLW='\033[1;33m'
CYN='\033[1;36m'
DIM='\033[2;37m'
RST='\033[0m'

# Load configurations
if [ -f .env ]; then
  source .env
fi
LAB_DOMAIN="${DOMAIN:-hackthacker.lab}"

clear
echo -e "${CYN}======================================================================${RST}"
echo -e "${WHT}              OWASP ATTACKFORGE: TELEMETRY & STATUS                   ${RST}"
echo -e "${CYN}======================================================================${RST}\n"

# Helper function to check container health states
check_container() {
  local container="$1"
  local display_name="$2"
  
  if ! docker ps -a --format '{{.Names}}' | grep -q "^${container}$"; then
    echo -e "  $(printf '%-20s' "$display_name") : ${RED}[ NOT DEPLOYED ]${RST}"
    return
  fi
  
  local status=$(docker inspect --format='{{.State.Status}}' "$container" 2>/dev/null || echo "unknown")
  
  if [ "$status" != "running" ]; then
    echo -e "  $(printf '%-20s' "$display_name") : ${RED}[ STOPPED ($status) ]${RST}"
    return
  fi
  
  local has_health=$(docker inspect --format='{{if .State.Health}}yes{{else}}no{{end}}' "$container" 2>/dev/null || echo "no")
  
  if [ "$has_health" = "yes" ]; then
    local health=$(docker inspect --format='{{.State.Health.Status}}' "$container" 2>/dev/null || echo "unknown")
    if [ "$health" = "healthy" ]; then
      echo -e "  $(printf '%-20s' "$display_name") : ${GRN}[ HEALTHY ]${RST}"
    elif [ "$health" = "starting" ]; then
      echo -e "  $(printf '%-20s' "$display_name") : ${YLW}[ STARTING ]${RST}"
    else
      echo -e "  $(printf '%-20s' "$display_name") : ${RED}[ UNHEALTHY ]${RST}"
    fi
  else
    echo -e "  $(printf '%-20s' "$display_name") : ${GRN}[ RUNNING ]${RST} ${DIM}(no healthcheck)${RST}"
  fi
}

echo -e "${CYN}--- Container Services status ----------------------------------------${RST}"
check_container "hackthacker-labs-nginx"       "Nginx Proxy Router"
check_container "hackthacker-labs-db"          "MariaDB Database"
check_container "hackthacker-labs-mutillidae"  "Mutillidae II (PHP)"
check_container "hackthacker-labs-dvwa"        "DVWA (PHP)"
check_container "hackthacker-labs-bwapp"       "bWAPP (PHP)"
check_container "hackthacker-labs-xvwa"        "XVWA (PHP)"
check_container "hackthacker-labs-vwa"         "VWA (PHP)"
check_container "hackthacker-labs-juiceshop"   "Juice Shop (Node)"
check_container "hackthacker-labs-webgoat"     "WebGoat (Java)"
check_container "hackthacker-labs-webwolf"     "WebWolf (Java)"
check_container "hackthacker-labs-tomcat"      "Apache Tomcat (Java)"
check_container "hackthacker-labs-wrongsecrets" "OWASP WrongSecrets"
check_container "hackthacker-labs-securityshepherd-db" "SecShepherd DB"
check_container "hackthacker-labs-securityshepherd-mongo" "SecShepherd NoSQL"
check_container "hackthacker-labs-securityshepherd" "OWASP SecShepherd"
check_container "hackthacker-labs-vulnerableapp-facade" "VulnerableApp Facade"
check_container "hackthacker-labs-crapi-web"   "crAPI Frontend"
check_container "hackthacker-labs-brokencrystals-app" "Broken Crystals Web"
check_container "hackthacker-labs-dvws-node"   "DVWS Node App"
check_container "hackthacker-labs-zerohealth-client" "ZeroHealth Frontend"
check_container "hackthacker-labs-restaurant-app" "RESTaurant App"
echo ""

# Helper to verify HTTP status code via Nginx proxy (binding to 127.0.0.1:443 locally)
check_endpoint() {
  local app_name="$1"
  local display_name="$2"
  local uri="$3"
  local creds="$4"
  
  local url="https://${app_name}.${LAB_DOMAIN}${uri}"
  
  # Fetch HTTP status code locally by hitting the loopback Nginx mapping with custom Host header
  local code=$(curl -s -k -o /dev/null -w "%{http_code}" -H "Host: ${app_name}.${LAB_DOMAIN}" "https://127.0.0.1:443${uri}" --connect-timeout 2 || echo "000")
  
  local http_stat="${RED}[DOWN]${RST}"
  if [[ "$code" =~ ^(200|301|302|401|403|404)$ ]]; then
    http_stat="${GRN}[ UP (HTTP $code) ]${RST}"
  else
    http_stat="${RED}[ DOWN (HTTP $code) ]${RST}"
  fi
  
  printf "  %-15s | %-38s | %-18s | %b\n" "$display_name" "$url" "$creds" "$http_stat"
}

echo -e "${CYN}--- Application URLs, Credentials & Endpoint Health ------------------${RST}"
printf "  ${WHT}%-15s${RST} | ${CYN}%-38s${RST} | ${YLW}%-18s${RST} | ${WHT}%s${RST}\n" "APP NAME" "ACCESS URL" "CREDENTIALS" "ENDPOINT STATUS"
echo -e "  ----------------|----------------------------------------|--------------------|----------------"
check_endpoint "mutillidae" "Mutillidae II"  "/"                 "admin / adminpass"
check_endpoint "dvwa"       "DVWA"           "/login.php"        "admin / password"
check_endpoint "bwapp"      "bWAPP"          "/login.php"        "bee / bug"
check_endpoint "xvwa"       "XVWA"           "/"                 "admin / admin"
check_endpoint "vwa"        "VWA"            "/"                 "admin / password"
check_endpoint "juiceshop"  "Juice Shop"     "/"                 "(Register in UI)"
check_endpoint "webgoat"    "WebGoat"        "/WebGoat/login"    "(Create in UI)"
check_endpoint "webwolf"    "WebWolf"        "/login"            "(Same as WebGoat)"
check_endpoint "tomcat"     "Tomcat Console" "/"                 "${DB_USER:-hackthacker} / ${DB_PASS:-hackthacker}"
check_endpoint "wrongsecrets" "WrongSecrets"   "/"                 "(Challenges/No Creds)"
check_endpoint "securityshepherd" "SecurityShep"  "/"                 "admin / password"
check_endpoint "vulnerableapp" "VulnerableApp" "/"                 "(Unified Gateway)"
check_endpoint "crapi"         "OWASP crAPI"    "/"                 "(Register in UI)"
check_endpoint "crapi-mailhog" "crAPI Mailhog"  "/"                 "(Mail Inbox UI)"
check_endpoint "brokencrystals" "BrokenCrystals" "/"               "(Register in UI)"
check_endpoint "brokencrystals-mailcatcher" "BC Mailcatcher" "/"    "(Mail Inbox UI)"
check_endpoint "dvws"          "DVWS Node"      "/"                 "(API Challenges)"
check_endpoint "zerohealth"    "ZeroHealth Web" "/"                 "(Health Portal UI)"
check_endpoint "zerohealth-api" "ZeroHealth API" "/api/config"      "(API Backend Conf)"
check_endpoint "restaurant"    "RESTaurant API" "/docs"             "(Swagger API UI)"
echo ""

echo -e "${CYN}--- Hosts entry mapping check ---------------------------------------${RST}"
# Check if current host maps one of the domains locally
if grep -q "dvwa.${LAB_DOMAIN}" /etc/hosts 2>/dev/null; then
  echo -e "  ${GRN}[✓] Hosts entry detected in /etc/hosts${RST}"
else
  echo -e "  ${YLW}[!] Domain resolution check failed.${RST}"
  echo -e "      Ensure your system hosts file contains mappings for all configured domains.${RST}"
fi
echo -e "${CYN}======================================================================${RST}"

