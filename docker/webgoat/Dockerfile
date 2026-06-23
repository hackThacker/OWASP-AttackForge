# ==============================================================================
# STAGE 1: Download Jar
# ==============================================================================
FROM alpine:3.19 AS downloader
RUN apk add --no-cache wget

WORKDIR /app
RUN wget -q https://github.com/WebGoat/WebGoat/releases/download/v8.2.2/webgoat-server-8.2.2.jar -O webgoat.jar

# ==============================================================================
# STAGE 2: Runtime
# ==============================================================================
FROM eclipse-temurin:21-jre-alpine

# Set up secure non-root running context
RUN addgroup -S webgoat && adduser -S webgoat -G webgoat
WORKDIR /home/webgoat

COPY --from=downloader --chown=webgoat:webgoat /app/webgoat.jar /home/webgoat/webgoat.jar

USER webgoat
EXPOSE 8080
EXPOSE 9001

ENTRYPOINT ["java", "-jar", "webgoat.jar", "--server.port=8080", "--server.address=0.0.0.0"]
