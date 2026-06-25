# ==============================================================================
# OWASP Juice Shop - Rebranded & Hardened Image
# ==============================================================================
# Use official multi-arch Juice Shop image as the base to ensure native ARM64 & AMD64 support.
FROM bkimminich/juice-shop:v17.1.1

# Enforce the exposed port
EXPOSE 3000

# Configure execution environment
ENV PORT=3000 \
    NODE_ENV=production

# Embedded healthcheck using node (exec form is required as distroless has no shell)
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
  CMD ["/nodejs/bin/node", "-e", "require('http').get('http://localhost:3000/', (res) => { if (res.statusCode === 200) process.exit(0); else process.exit(1); }).on('error', (e) => { process.exit(1); })"]
