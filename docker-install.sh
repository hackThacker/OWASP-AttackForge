#!/bin/bash
# ==============================================================================
# OWASP LAB TOOLKIT - DOCKER INSTALLER
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

echo -e "${CYN}======================================================================${RST}"
echo -e "${WHT}                 OWASP ATTACKFORGE: DOCKER INSTALLER                  ${RST}"
echo -e "${CYN}======================================================================${RST}\n"

# 1. Environment Config Setup
if [ ! -f .env ]; then
  echo -e "${YLW}[!] Creating .env configuration from template...${RST}"
  cp .env.example .env
  echo -e "${GRN}[✓] Created .env file. Please edit it to customize settings if required.${RST}"
else
  echo -e "${GRN}[✓] Found existing .env file.${RST}"
fi

# Load variables
source .env
LAB_DOMAIN="${DOMAIN:-hackthacker.lab}"

# 2. Prerequisites Check
if ! command -v docker >/dev/null; then
  echo -e "${RED}[x] Error: Docker is not installed. Please install Docker and try again.${RST}"
  exit 1
fi

# Check for compose support
if docker compose version >/dev/null 2>&1; then
  COMPOSE_CMD="docker compose"
elif command -v docker-compose >/dev/null; then
  COMPOSE_CMD="docker-compose"
else
  echo -e "${RED}[x] Error: Docker Compose is not installed. Please install it and try again.${RST}"
  exit 1
fi
echo -e "${GRN}[✓] Prerequisites checked. Using '${COMPOSE_CMD}' for deployment.${RST}\n"

# 3. Booting the stack
echo -e "${CYN}[*] Building and launching OWASP Lab container stack...${RST}"
$COMPOSE_CMD up --build -d

echo -e "\n${GRN}[✓] Container orchestration started successfully!${RST}\n"

# 4. Hosts File Mapping Guidance
# Build the hosts entry string
APPS=("mutillidae" "dvwa" "bwapp" "xvwa" "vwa" "juiceshop" "webgoat" "webwolf" "tomcat")
HOSTS_ENTRY="127.0.0.1"
for app in "${APPS[@]}"; do
  HOSTS_ENTRY="${HOSTS_ENTRY} ${app}.${LAB_DOMAIN}"
done

echo -e "${CYN}----------------------------------------------------------------------${RST}"
echo -e "${WHT}                     LOCAL DOMAIN CONFIGURATION                       ${RST}"
echo -e "${CYN}----------------------------------------------------------------------${RST}"

# Attempt to configure hosts if run under Linux/macOS with sudo permissions
if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
  if [ "$EUID" -eq 0 ] || sudo -n true 2>/dev/null; then
    echo -e "${CYN}[*] Linux/macOS environment detected. Updating /etc/hosts...${RST}"
    # Remove older entries containing domain name to prevent duplicates
    sudo sed -i "/\.${LAB_DOMAIN}/d" /etc/hosts || true
    echo -e "${HOSTS_ENTRY}" | sudo tee -a /etc/hosts >/dev/null
    echo -e "${GRN}[✓] /etc/hosts file updated successfully!${RST}\n"
  else
    echo -e "${YLW}[!] Sudo authentication is needed to automatically map domains in /etc/hosts.${RST}"
    echo -e "    Please run the following command manually:"
    echo -e "    ${WHT}echo \"${HOSTS_ENTRY}\" | sudo tee -a /etc/hosts${RST}\n"
  fi
else
  # Guide Windows hosts file updates
  echo -e "${YLW}[!] If you are running on Windows, please run Notepad as Administrator,${RST}"
  echo -e "    open ${WHT}C:\\Windows\\System32\\drivers\\etc\\hosts${RST} and add the line below:"
  echo -e "    ${GRN}${HOSTS_ENTRY}${RST}\n"
fi

echo -e "${CYN}----------------------------------------------------------------------${RST}"
echo -e "${GRN}[+] DEPLOYMENT COMPLETED!${RST}"
echo -e "    To check status and get application URLs, run:"
echo -e "    ${WHT}./docker-check.sh${RST}"
echo -e "${CYN}----------------------------------------------------------------------${RST}"
