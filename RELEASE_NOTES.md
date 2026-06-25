# Release v1.1.0 (Stable Release)

## 📝 Executive Summary
This release marks the transition of the cyber range to **OWASP AttackForge v1.1.0**. It expands the self-contained laboratory stack by integrating **seven new vulnerable applications and API sets**, bringing the total count to **18 distinct cyber labs**. It also resolves critical CI/CD runner compilation errors, introduces smart fallback paths for repositories without external container registry keys, and fixes minor startup routing issues across the application microservices.

---

## 🚀 New Features

We have integrated seven additional vulnerable applications into the training range:

| Application | Domain & URL | Username | Password / Mode | Primary Vulnerability Context |
| :--- | :--- | :--- | :--- | :--- |
| **SasanLabs VulnerableApp** | `https://vulnerableapp.hackthacker.lab` | (Gateway UI) | (Unified UI) | Cross-site Scripting (XSS), SQL Injection, Command Injection, JWT bypasses |
| **OWASP crAPI** | `https://crapi.hackthacker.lab` | Register in UI | Register in UI | API vulnerabilities, BOLA/IDOR, broken authentication, rate-limiting bypass |
| **crAPI Mailhog** | `https://crapi-mailhog.hackthacker.lab` | - | Inbox Viewer | Mock email interface to intercept password reset tokens for API labs |
| **Broken Crystals** | `https://brokencrystals.hackthacker.lab` | Register in UI | Register in UI | Next-gen JS vulnerabilities, gRPC injection, Keycloak OAuth bypass, Ollama AI safety leaks |
| **BC Mailcatcher** | `https://brokencrystals-mailcatcher.hackthacker.lab` | - | Inbox Viewer | Mock mail interface to capture verification links and session tokens |
| **DVWS Node** | `https://dvws.hackthacker.lab` | (API Tasks) | (API Tasks) | Damn Vulnerable Web Services in Node.js (REST, SOAP, GraphQL, and WebSockets flaws) |
| **Zero-Health Web** | `https://zerohealth.hackthacker.lab` | (Portal UI) | (Portal UI) | Vulnerable patient portal demonstrating CORS flaws, IDOR, and SSRF in Node/React stacks |
| **Zero-Health API** | `https://zerohealth-api.hackthacker.lab` | - | API Health | Backend REST API server for Zero-Health portal (JWT vulnerabilities) |
| **RESTaurant API** | `https://restaurant.hackthacker.lab/docs` | (Swagger UI) | Swagger UI | FastAPI-based challenge containing IDOR, SQLi, Mass Assignment, and RCE vectors |

---

## 🏗️ Infrastructure & CI/CD Improvements

*   **Conditional Docker Hub Registry Fallback**: Overhauled `build.yml` to automatically skip Docker Hub authentication and skips pushing `docker.io/` tags if repository secrets (`DOCKERHUB_USERNAME`) are not configured. This enables seamless container builds and GitHub Container Registry (GHCR) publishing in forks or third-party repositories without needing external hub keys.
*   **Reusable Workflows Concurrency Fix**: Removed top-level `concurrency` specifications from `build.yml` to prevent GitHub Actions compiler from aborting with workflow startup validation failures. Concurrency settings are now correctly managed by GHA caller workflows.
*   **Decoupled Reusable Job Permissions**: Declared explicit job-level GHA permissions for caller jobs in `release.yml` (`contents: read/write`, `packages: write`, `id-token: write`, and `attestations: write`) to ensure all signing and image publication tasks execute with sufficient scopes.
*   **Supply Chain Attestation Expansion**: Updated `supply-chain-attest.yml` SBOM matrix to automatically sign and attest the seven newly introduced application container images.

---

## 🔒 Security Hardening & Isolation

*   **Container Non-Root Contexts**: Hardened the Dockerfiles of all seven new applications to run processes under standard user contexts (`node`, `python`, `keycloak`) instead of root namespaces.
*   **Isolated Database Bridging**: Excluded public port mapping for backend databases (Broken Crystals, Zero-Health, crAPI, RESTaurant, DVWS DBs) and isolated them completely on the internal `hackthacker` bridge network, routing only web traffic through the proxy.
*   **Wildcard Nginx Router**: Provisioned secure SSL routing blocks in the main Nginx templates to handle the new custom local domain paths without breaking access to original environments.

---

## 🔧 Changed & Improved Components

*   **README and Documentation Sync**: Synced container counts (45 containers total) and lab lists across the repository documentation files (`README.md`, `docs/cicd/README.md`).
*   **Workspace Clean Up**: Added untracked shell scripts (`check-owasp-lab.sh`, `install-owasp-lab.sh`, etc.) to `.gitignore` to prevent committing installer and cache assets during manual push operations.

---

## 🐛 Bug Fixes

*   **Keycloak Boot Race Solved**: Patched the Keycloak database waiting scripts in Broken Crystals to prevent the auth server from starting before its relational database socket is ready.
*   **RESTaurant DB Compilation Fixed**: Resolved poetry schema database migration conflicts in RESTaurant API initialization loops.
*   **Git-Cliff Context Rendering Fixed**: Replaced template environment variables in `cliff.toml` with the explicit repository path to bypass template rendering context errors.
*   **Detached HEAD Push Resolved**: Modified the changelog push step in `changelog.yml` to use `git push origin HEAD:main`, allowing successful pushes from detached HEAD states on tag commits.

---

## 💻 Access & Startup Instructions

To start the range and resolve local DNS paths:

```bash
# Clone the range repository
git clone https://github.com/hackThacker/OWASP-AttackForge.git
cd OWASP-AttackForge

# Run the automated installer
sudo ./docker-install.sh
```

---

## 🌟 Why This Release Matters (Impact Statement)
This update expands **OWASP AttackForge** into a fully featured DevSecOps training facility. The new integrations bring modern REST/GraphQL, OAuth, gRPC, and AI model vulnerabilities into play. Along with key infrastructure fixes, this allows security trainers to deploy a secure, highly hardened local cyber range out-of-the-box on both ARM64 and AMD64 environments.
