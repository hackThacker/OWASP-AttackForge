# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/aligorithm/Zero-Health.git .

# ==============================================================================
# STAGE 2: Build client
# ==============================================================================
FROM node:18-alpine AS builder

ARG VITE_API_URL=https://zerohealth-api.owasp.hacking
ENV VITE_API_URL=$VITE_API_URL

WORKDIR /app

COPY --from=git-cloner /src/client /app

RUN npm ci --no-audit || npm install --no-audit
RUN npm run build

# ==============================================================================
# STAGE 3: Runtime serving via Nginx
# ==============================================================================
FROM nginx:1.25-alpine

# Copy built production assets to Nginx html directory
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy our optimized Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Setup permission ownership for non-root execution
RUN mkdir -p /var/cache/nginx /var/run /run && \
    chown -R nginx:nginx /usr/share/nginx/html /etc/nginx /var/cache/nginx /var/run /run /var/log/nginx

USER nginx
EXPOSE 3000

CMD ["nginx", "-g", "daemon off;"]
