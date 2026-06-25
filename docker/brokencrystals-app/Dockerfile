# ==============================================================================
# STAGE 1: Source Code Downloader
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/NeuraLegion/brokencrystals.git .

# ==============================================================================
# STAGE 2: Build NestJS and Client
# ==============================================================================
FROM node:18-alpine AS build

WORKDIR /usr/src/app

COPY --from=git-cloner /src/package*.json ./
COPY --from=git-cloner /src/tsconfig.build.json ./
COPY --from=git-cloner /src/tsconfig.json ./
COPY --from=git-cloner /src/nest-cli.fast.json ./
COPY --from=git-cloner /src/.env ./
COPY --from=git-cloner /src/config ./config
COPY --from=git-cloner /src/keycloak ./keycloak
COPY --from=git-cloner /src/src ./src

ENV NPM_CONFIG_LOGLEVEL=error
RUN npm ci --no-audit
RUN npm run build:fast
RUN npm prune --production

COPY --from=git-cloner /src/client/package*.json ./client/
COPY --from=git-cloner /src/client/src ./client/src
COPY --from=git-cloner /src/client/public ./client/public
COPY --from=git-cloner /src/client/typings ./client/typings
COPY --from=git-cloner /src/client/vcs ./client/vcs
COPY --from=git-cloner /src/client/tsconfig.json ./client/tsconfig.json
COPY --from=git-cloner /src/client/vite.config.ts ./client/vite.config.ts
COPY --from=git-cloner /src/client/index.html ./client/index.html

ENV CYPRESS_INSTALL_BINARY=0
RUN npm ci --prefix=client --no-audit
RUN npm run build --prefix=client

# ==============================================================================
# STAGE 3: Production Runtime
# ==============================================================================
FROM node:18-alpine AS production

WORKDIR /usr/src/app

COPY --from=git-cloner /src/.env ./
COPY --from=git-cloner /src/config ./config
COPY --from=git-cloner /src/keycloak ./keycloak

COPY --from=build /usr/src/app/node_modules ./node_modules
COPY --from=build /usr/src/app/package*.json ./
COPY --from=build /usr/src/app/dist ./dist

COPY --from=build /usr/src/app/client/dist ./client/dist
COPY --from=build /usr/src/app/client/vcs ./client/vcs

USER node
EXPOSE 3000
CMD ["npm", "run", "start:prod"]
