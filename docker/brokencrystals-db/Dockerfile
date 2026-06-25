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
FROM postgres:17-alpine
COPY --from=builder /src/pg.sql /docker-entrypoint-initdb.d/pg.sql
