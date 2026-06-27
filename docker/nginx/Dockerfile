FROM nginx:1.31-alpine

# Install bash and openssl for SSL certificate generation
RUN apk add --no-cache openssl bash curl

# Create directory structures and setup permission ownership for non-root execution
RUN mkdir -p /etc/nginx/ssl /etc/nginx/templates /var/cache/nginx /var/run /run \
    && chown -R nginx:nginx /etc/nginx /var/cache/nginx /var/run /run /var/log/nginx

# Copy configuration template, entrypoint and static landing page
COPY docker/nginx/default.conf.template /etc/nginx/templates/default.conf.template
COPY docker/nginx/entrypoint.sh /usr/local/bin/entrypoint.sh
COPY docker/frontend/index.html /usr/share/nginx/html/index.html

RUN chmod +x /usr/local/bin/entrypoint.sh

# Run as non-root user
USER nginx
EXPOSE 8080 8443

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
