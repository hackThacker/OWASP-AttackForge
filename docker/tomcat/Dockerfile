FROM tomcat:10.1-jre21-temurin-jammy

# Install curl for health checking
RUN apt-get update && apt-get install -y --no-install-recommends curl && rm -rf /var/lib/apt/lists/*

# Copy manager and host-manager apps from webapps.dist if present, and remove localhost restrictions
RUN if [ -d /usr/local/tomcat/webapps.dist/manager ]; then \
        cp -r /usr/local/tomcat/webapps.dist/* /usr/local/tomcat/webapps/; \
    fi && \
    sed -i '/<Valve/,/\/>/d' /usr/local/tomcat/webapps/manager/META-INF/context.xml || true && \
    sed -i '/<Valve/,/\/>/d' /usr/local/tomcat/webapps/host-manager/META-INF/context.xml || true

# Copy custom users definition and entrypoint
COPY docker/tomcat/tomcat-users.xml /usr/local/tomcat/conf/tomcat-users.xml
COPY docker/tomcat/entrypoint.sh /usr/local/bin/entrypoint.sh

# Build secure runtime execution environment
RUN groupadd -r tomcat && useradd -r -g tomcat tomcat && \
    chmod +x /usr/local/bin/entrypoint.sh && \
    chown -R tomcat:tomcat /usr/local/tomcat

# Run as non-root
USER tomcat
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["catalina.sh", "run"]
