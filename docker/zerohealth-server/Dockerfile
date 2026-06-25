# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/aligorithm/Zero-Health.git .

# ==============================================================================
# STAGE 2: Runtime
# ==============================================================================
FROM node:18-alpine

RUN apk add --no-cache postgresql-client curl

WORKDIR /app

COPY --from=git-cloner /src /app

# Install build dependencies for compiling native node modules (like libxmljs)
RUN apk add --no-cache --virtual .build-deps build-base python3 && \
    npm install && \
    apk del .build-deps

# Make scripts executable
RUN chmod +x scripts/*.sh

EXPOSE 5000

CMD ["node", "server.js"]
