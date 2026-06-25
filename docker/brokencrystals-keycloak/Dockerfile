# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS builder
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/NeuraLegion/brokencrystals.git .

# ==============================================================================
# STAGE 2: Runtime
# ==============================================================================
FROM quay.io/keycloak/keycloak:26.1.2
COPY --from=builder /src/keycloak/imports/realm-export.json /opt/keycloak/data/import/realm-export.json
