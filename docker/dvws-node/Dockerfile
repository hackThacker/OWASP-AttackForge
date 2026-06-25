# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/snoopysecurity/dvws-node.git .

# ==============================================================================
# STAGE 2: Runtime
# ==============================================================================
FROM node:22-slim

# Install build dependencies for native node addons
RUN apt-get update && apt-get install -y --no-install-recommends \
    python3 \
    make \
    g++ \
    build-essential \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /home/dvws-node

COPY --from=git-cloner /src /home/dvws-node

# Install package dependencies
RUN npm install

# Add wait script to coordinate database readiness
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

EXPOSE 80 4000 9090

CMD /wait && node startup_script.js && npm start
