#!/bin/bash
set -e

# ==============================================================================
# Dynamic Database Configuration Injection
# ==============================================================================

# DVWA Config Injection
if [ "$APP_NAME" = "dvwa" ]; then
  FILE="/var/www/html/config/config.inc.php"
  if [ -f "$FILE" ]; then
    echo "Injecting DVWA database configuration..."
    sed -i "s/\$_DVWA\[ 'db_user' \].*/\$_DVWA\[ 'db_user' \] = '${DB_USER}';/g" "$FILE"
    sed -i "s/\$_DVWA\[ 'db_password' \].*/\$_DVWA\[ 'db_password' \] = '${DB_PASS}';/g" "$FILE"
    sed -i "s/\$_DVWA\[ 'db_server' \].*/\$_DVWA\[ 'db_server' \] = '${DB_HOST}';/g" "$FILE"
  fi
fi

# bWAPP Config Injection
if [ "$APP_NAME" = "bwapp" ]; then
  FILE="/var/www/html/bWAPP/admin/settings.php"
  if [ -f "$FILE" ]; then
    echo "Injecting bWAPP database configuration..."
    sed -i "s/\$db_username =.*/\$db_username = \"${DB_USER}\";/g" "$FILE"
    sed -i "s/\$db_password =.*/\$db_password = \"${DB_PASS}\";/g" "$FILE"
    sed -i "s/\$db_server =.*/\$db_server = \"${DB_HOST}\";/g" "$FILE"
  fi
fi

# XVWA Config Injection
if [ "$APP_NAME" = "xvwa" ]; then
  FILE="/var/www/html/config.php"
  if [ -f "$FILE" ]; then
    echo "Injecting XVWA database configuration..."
    sed -i "s/\$user =.*/\$user = '${DB_USER}';/g" "$FILE"
    sed -i "s/\$pass =.*/\$pass = '${DB_PASS}';/g" "$FILE"
    sed -i "s/\$host =.*/\$host = '${DB_HOST}';/g" "$FILE"
  fi
fi

# Mutillidae Config Injection
if [ "$APP_NAME" = "mutillidae" ]; then
  FILE="/var/www/html/src/includes/database-config.inc"
  if [ -f "$FILE" ]; then
    echo "Injecting Mutillidae database configuration..."
    sed -i "s/define('DB_USERNAME'.*/define('DB_USERNAME', '${DB_USER}');/g" "$FILE"
    sed -i "s/define('DB_PASSWORD'.*/define('DB_PASSWORD', '${DB_PASS}');/g" "$FILE"
    sed -i "s/define('DB_HOST'.*/define('DB_HOST', '${DB_HOST}');/g" "$FILE"
  fi
fi

# VWA Config Injection (Vulnerable Web Application)
if [ "$APP_NAME" = "vwa" ]; then
  echo "Injecting VWA database configuration..."
  find /var/www/html -type f -name "*.php" -exec sed -i "s/\$dbuser = 'root';/\$dbuser = '${DB_USER}';/g" {} +
  find /var/www/html -type f -name "*.php" -exec sed -i "s/\$dbpass = '';/\$dbpass = '${DB_PASS}';/g" {} +
  find /var/www/html -type f -name "*.php" -exec sed -i "s/\$username = \"root\";/\$username = \"${DB_USER}\";/g" {} +
  find /var/www/html -type f -name "*.php" -exec sed -i "s/\$password = \"\";/\$password = \"${DB_PASS}\";/g" {} +
  
  # Ensure VWA connects to the containerised database hostname instead of localhost
  find /var/www/html -type f -name "*.php" -exec sed -i "s/mysqli_connect('localhost'/mysqli_connect('${DB_HOST}'/g" {} +
  find /var/www/html -type f -name "*.php" -exec sed -i "s/mysqli_connect(\"localhost\"/mysqli_connect(\"${DB_HOST}\"/g" {} +
  find /var/www/html -type f -name "*.php" -exec sed -i "s/new mysqli('localhost'/new mysqli('${DB_HOST}'/g" {} +
  find /var/www/html -type f -name "*.php" -exec sed -i "s/new mysqli(\"localhost\"/new mysqli(\"${DB_HOST}\"/g" {} +
fi

# ==============================================================================
# Apache DocumentRoot Adjustment
# ==============================================================================
if [ -n "$APACHE_DOCROOT" ]; then
  echo "Configuring Apache DocumentRoot to ${APACHE_DOCROOT}..."
  sed -i "s|DocumentRoot /var/www/html|DocumentRoot ${APACHE_DOCROOT}|g" /etc/apache2/sites-available/000-default.conf
fi

# Execute CMD
exec "$@"
