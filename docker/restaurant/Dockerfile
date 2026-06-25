# ==============================================================================
# STAGE 1: Source Code Downloader & Requirement Export
# ==============================================================================
FROM alpine:3.19 AS git-cloner
RUN apk add --no-cache git
WORKDIR /src
RUN git clone --depth 1 https://github.com/theowni/Damn-Vulnerable-RESTaurant-API-Game.git .

FROM python:3.10-bookworm as builder
RUN pip install poetry==1.4.2
WORKDIR /app
COPY --from=git-cloner /src/pyproject.toml /src/poetry.lock ./
RUN poetry export -f requirements.txt --output requirements.txt --without-hashes

# ==============================================================================
# STAGE 2: Runtime
# ==============================================================================
FROM python:3.10-slim-bookworm

RUN apt-get update && apt-get -y install --no-install-recommends \
    libpq-dev \
    gcc \
    vim \
    sudo \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=git-cloner /src /app
COPY --from=builder /app/requirements.txt /app/requirements.txt

RUN pip install --no-cache-dir -r requirements.txt

# Allow app user to run find with root permissions via sudo for the challenge
RUN echo 'ALL ALL=(ALL) NOPASSWD: /usr/bin/find' | sudo tee /etc/sudoers.d/find_nopasswd > /dev/null

# Create a dedicated non-root user
RUN useradd -m app && chown -R app:app /app
USER app

EXPOSE 8091

CMD ["bash", "-c", "alembic upgrade head && uvicorn app.main:app --host 0.0.0.0 --port 8091 --workers 1"]
