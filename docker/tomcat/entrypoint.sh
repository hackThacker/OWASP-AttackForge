#!/bin/sh
set -e

# Determine username and password from environment variables
USER_NAME="${TOMCAT_USER:-${DB_USER:-admin}}"
USER_PASS="${TOMCAT_PASS:-${DB_PASS:-adminpass}}"

echo "Injecting Tomcat administrator credentials: ${USER_NAME}"

# Replace placeholders in the Tomcat users configuration
sed -i "s/TOMCAT_USER_PLACEHOLDER/${USER_NAME}/g" /usr/local/tomcat/conf/tomcat-users.xml
sed -i "s/TOMCAT_PASS_PLACEHOLDER/${USER_PASS}/g" /usr/local/tomcat/conf/tomcat-users.xml

# Execute the default container command (usually 'catalina.sh run')
exec "$@"
