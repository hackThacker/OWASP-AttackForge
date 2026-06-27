# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/aligorithm/Zero-Health.git .

# ==============================================================================
# STAGE 2: Builder
# ==============================================================================
FROM node:18-alpine AS builder
WORKDIR /app
COPY --from=git-cloner /src /app
# Install build dependencies for compiling native node modules
RUN apk add --no-cache build-base python3 && \
    npm ci --no-audit || npm install --no-audit

# ==============================================================================
# STAGE 3: Runtime
# ==============================================================================
FROM node:18-alpine
RUN apk add --no-cache postgresql-client curl

WORKDIR /app
COPY --from=git-cloner /src /app
COPY --from=builder /app/node_modules ./node_modules

# Make scripts executable and set ownership
RUN chmod +x scripts/*.sh && \
    chown -R node:node /app

USER node
EXPOSE 5000

CMD ["node", "server.js"]
