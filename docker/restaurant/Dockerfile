# ==============================================================================
# STAGE 1: Source Code Downloader & Requirement Export
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/theowni/Damn-Vulnerable-RESTaurant-API-Game.git .

FROM python:3.10-slim-bookworm AS builder
RUN pip install --no-cache-dir poetry==1.4.2
WORKDIR /app
COPY --from=git-cloner /src/pyproject.toml /src/poetry.lock ./
RUN poetry export -f requirements.txt --output requirements.txt --without-hashes

# Install compiler tools
RUN apt-get update && apt-get -y install --no-install-recommends \
    libpq-dev \
    gcc \
    libc-dev \
    && rm -rf /var/lib/apt/lists/*

# Install packages into system site-packages
RUN pip install --no-cache-dir -r requirements.txt

# ==============================================================================
# STAGE 3: Runtime
# ==============================================================================
FROM python:3.10-slim-bookworm

# Install ONLY runtime dependencies (libpq-dev is needed for postgres runtime link, curl for health checks, sudo/vim for challenges)
RUN apt-get update && apt-get -y install --no-install-recommends \
    libpq-dev \
    curl \
    vim \
    sudo \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy python packages and executable entrypoints from builder
COPY --from=builder /usr/local/lib/python3.10/site-packages /usr/local/lib/python3.10/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin

COPY --from=git-cloner /src /app

# Allow app user to run find with root permissions via sudo for the challenge
RUN mkdir -p /etc/sudoers.d && echo 'ALL ALL=(ALL) NOPASSWD: /usr/bin/find' > /etc/sudoers.d/find_nopasswd

# Create a dedicated non-root user
RUN useradd -m app && chown -R app:app /app

WORKDIR /app/app
USER app

EXPOSE 8091

CMD ["bash", "-c", "alembic upgrade head && uvicorn main:app --host 0.0.0.0 --port 8091 --workers 1"]
