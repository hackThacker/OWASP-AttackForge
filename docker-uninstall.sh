#!/bin/bash
# ==============================================================================
# OWASP LAB TOOLKIT - DOCKER UNINSTALLER
# ==============================================================================

set -e

# Terminal Colors
WHT='\033[1;37m'
RED='\033[1;31m'
GRN='\033[1;32m'
YLW='\033[1;33m'
CYN='\033[1;36m'
RST='\033[0m'

echo -e "${RED}======================================================================${RST}"
echo -e "${WHT}                 OWASP ATTACKFORGE: DOCKER DESTRUCTION                ${RST}"
echo -e "${RED}======================================================================${RST}"
echo -e "${YLW}[!] Warning: This will stop all containers and destroy ALL persistent volumes!${RST}\n"

read -p "Are you sure you want to completely uninstall the dockerized lab? (y/N): " confirm
if [[ "$confirm" != "y" && "$confirm" != "Y" ]]; then
  echo -e "${GRN}[*] Uninstallation aborted. Stack is safe.${RST}"
  exit 0
fi

# Load variables
if [ -f .env ]; then
  source .env
fi
LAB_DOMAIN="${DOMAIN:-hackthacker.lab}"

# Determine Compose Command
if docker compose version >/dev/null 2>&1; then
  COMPOSE_CMD="docker compose"
elif command -v docker-compose >/dev/null; then
  COMPOSE_CMD="docker-compose"
else
  echo -e "${RED}[x] Error: Docker / Docker Compose command not found.${RST}"
  exit 1
fi

echo -e "\n${CYN}[*] Tearing down docker orchestration (including volumes)...${RST}"
$COMPOSE_CMD down -v

# Remove local host mappings if on Linux/macOS
if [[ "$OSTYPE" == "linux-gnu"* || "$OSTYPE" == "darwin"* ]]; then
  if [ "$EUID" -eq 0 ] || sudo -n true 2>/dev/null; then
    echo -e "${CYN}[*] Cleaning up /etc/hosts entries...${RST}"
    sudo sed -i "/\.${LAB_DOMAIN}/d" /etc/hosts || true
    echo -e "${GRN}[✓] Mappings removed from /etc/hosts.${RST}"
  else
    echo -e "${YLW}[!] Sudo permission is needed to clean up /etc/hosts automatically.${RST}"
    echo -e "    Please manually remove the lines referencing '.${LAB_DOMAIN}' from /etc/hosts.${RST}"
  fi
fi

echo -e "\n${GRN}[✓] UNINSTALLATION COMPLETE! All containers and volumes destroyed.${RST}"
echo -e "${RED}======================================================================${RST}"
