FROM alpine:3.24

# Install curl and bash for the setup triggers
RUN apk add --no-cache curl bash

# Run as non-privileged system user
USER nobody

# Define default execution endpoint
ENTRYPOINT ["/bin/bash", "/usr/local/bin/init.sh"]
