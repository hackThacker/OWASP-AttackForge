# ==============================================================================
# STAGE 1: Download and Extract
# ==============================================================================
FROM alpine:3.19 AS downloader
RUN apk add --no-cache curl tar

ARG TARGETARCH
WORKDIR /app

# Download the specific release matching the target architecture
RUN if [ "$TARGETARCH" = "arm64" ]; then \
        ARCH="arm64"; \
    else \
        ARCH="x64"; \
    fi && \
    curl -L "https://github.com/juice-shop/juice-shop/releases/download/v17.1.1/juice-shop-17.1.1_node20_linux_${ARCH}.tgz" -o juiceshop.tgz \
    && tar -xzf juiceshop.tgz --strip-components=1 \
    && rm juiceshop.tgz

# ==============================================================================
# STAGE 2: Production Runtime
# ==============================================================================
FROM node:26-slim

WORKDIR /app

# Copy application files and ensure ownership is set to the default 'node' user
COPY --from=downloader --chown=node:node /app /app
RUN chown node:node /app

# Configure execution environment
USER node
ENV PORT=3000 \
    NODE_ENV=production

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
  CMD node -e "require('http').get('http://localhost:3000/', (res) => { if (res.statusCode === 200) process.exit(0); else process.exit(1); }).on('error', (e) => { process.exit(1); })"

CMD ["npm", "start"]
