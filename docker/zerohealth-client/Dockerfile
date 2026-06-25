# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/aligorithm/Zero-Health.git .

# ==============================================================================
# STAGE 2: Build & Runtime
# ==============================================================================
FROM node:18-alpine

ARG VITE_API_URL=https://zerohealth-api.owasp.hacking
ENV VITE_API_URL=$VITE_API_URL

WORKDIR /app

COPY --from=git-cloner /src/client /app

RUN npm install
RUN npm run build
RUN npm install -g serve

EXPOSE 3000

CMD ["serve", "-s", "dist", "-l", "3000"]
