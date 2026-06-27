# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/snoopysecurity/dvws-node.git .

# ==============================================================================
# STAGE 2: Builder
# ==============================================================================
FROM node:22-slim AS builder

WORKDIR /build

# Install build dependencies for native node addons
RUN apt-get update && apt-get install -y --no-install-recommends \
    python3 \
    make \
    g++ \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Copy package files and source code
COPY --from=git-cloner /src /build

# Install dependencies including compilation of native modules
RUN npm ci --no-audit || npm install --no-audit

# Download wait script
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /build/wait
RUN chmod +x /build/wait

# ==============================================================================
# STAGE 3: Runtime
# ==============================================================================
FROM node:22-slim

WORKDIR /home/dvws-node

# Copy built node_modules, wait, and source files from builder stage
COPY --from=builder /build /home/dvws-node

# Set non-root user permissions
RUN chown -R node:node /home/dvws-node

USER node
EXPOSE 80 4000 9090

CMD ["/home/dvws-node/wait", "&&", "node", "startup_script.js", "&&", "npm", "start"]
