# ==============================================================================
# STAGE 1: Build / Source Downloader
# ==============================================================================
FROM alpine:3.19 AS builder
RUN apk add --no-cache git

ARG APP_SOURCE_URL
ARG APP_NAME

WORKDIR /src

# Clone the repository dynamically with shallow depth to keep cache small
RUN git clone --depth 1 ${APP_SOURCE_URL} .

# Run post-clone patches
RUN if [ "${APP_NAME}" = "dvwa" ]; then \
      cp config/config.inc.php.dist config/config.inc.php; \
    fi && \
    if [ "${APP_NAME}" = "xvwa" ]; then \
      # Fix xvwa symlink recursion requirement
      ln -s . xvwa; \
    fi && \
    if [ "${APP_NAME}" = "mutillidae" ]; then \
      # Remove restrictive .htaccess files to bypass localhost restrictions
      find . -name ".htaccess" -delete; \
    fi

# ==============================================================================
# STAGE 2: Runtime Environment
# ==============================================================================
FROM php:8.3-apache-bookworm

# Prevent apt warnings
ENV DEBIAN_FRONTEND=noninteractive

# Install core systems dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    libzip-dev \
    libcurl4-openssl-dev \
    libxml2-dev \
    libonig-dev \
    curl \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install -j$(nproc) mysqli pdo_mysql gd zip curl mbstring xml \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

# Enable Apache rewrite engine
RUN a2enmod rewrite

# Harden container: Run Apache as non-root user 'www-data' on port 8080
# We modify Apache configs and change ownership so www-data can listen on 8080 and rewrite its own config templates.
RUN sed -i 's/Listen 80/Listen 8080/' /etc/apache2/ports.conf \
    && sed -i 's/<VirtualHost \*:80>/<VirtualHost \*:8080>/' /etc/apache2/sites-available/000-default.conf \
    && mkdir -p /var/run/apache2 /var/lock/apache2 /var/log/apache2 \
    && chown -R www-data:www-data /var/www /var/run/apache2 /var/lock/apache2 /var/log/apache2 /etc/apache2 /var/lib/apache2

# Copy vulnerable PHP configuration rules
COPY owasp-lab.ini /usr/local/etc/php/conf.d/owasp-lab.ini
COPY disable_strict_mysqli.php /usr/local/etc/php/disable_strict_mysqli.php

# Copy files from builder stage
COPY --from=builder --chown=www-data:www-data /src /var/www/html

# Copy and configure the startup entrypoint script
COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

# Run as non-root
USER www-data
WORKDIR /var/www/html

# Expose non-privileged HTTP port
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["apache2-foreground"]
