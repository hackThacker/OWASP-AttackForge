# ==============================================================================
# STAGE 1: Download Jar
# ==============================================================================
FROM alpine:3.19 AS downloader
RUN apk add --no-cache wget

WORKDIR /app
RUN wget -q https://github.com/WebGoat/WebGoat/releases/download/v8.2.2/webwolf-8.2.2.jar -O webwolf.jar

# ==============================================================================
# STAGE 2: Runtime
# ==============================================================================
FROM eclipse-temurin:21-jre-alpine

# Set up secure non-root running context
RUN addgroup -S webwolf && adduser -S webwolf -G webwolf
WORKDIR /home/webwolf

COPY --from=downloader --chown=webwolf:webwolf /app/webwolf.jar /home/webwolf/webwolf.jar

USER webwolf
EXPOSE 9090

ENTRYPOINT ["java", "-jar", "webwolf.jar", "--server.port=9090", "--server.address=0.0.0.0"]
