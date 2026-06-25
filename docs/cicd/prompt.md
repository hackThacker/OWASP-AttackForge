# Enterprise DevSecOps CI/CD Architecture Generator Prompt

You are a Principal DevOps Engineer, DevSecOps Architect, and Release Engineering Specialist.

Design a complete enterprise-grade CI/CD platform for the GitHub repository "OWASP-AttackForge".

Project Overview:

* OWASP AttackForge is a Docker-based offensive security training platform.
* It contains multiple vulnerable web applications.
* Applications use PHP, Java, Node.js, MariaDB, MongoDB, and Nginx.
* Deployment uses Docker Compose.
* Images must support both amd64 and arm64.
* Public releases are published on GitHub.
* Images are distributed through Docker Hub and GitHub Container Registry (GHCR).

Requirements:

## Container Registry Strategy

Implement dual publishing:

Primary:

* Docker Hub

Secondary:

* GitHub Container Registry (GHCR)

Image Naming:

Docker Hub:
docker.io/<org>/attackforge-<service>:<version>

GHCR:
ghcr.io/<org>/attackforge-<service>:<version>

Required Tags:

* latest
* major version
* minor version
* patch version
* git sha

Example:
v1.2.3

Generate:

* latest
* 1
* 1.2
* 1.2.3
* sha-abcdef1

---

## Semantic Version Enforcement

Implement automated semantic version validation.

Accepted:
v1.0.0
v1.2.5
v2.4.18

Reject:
1.0
version1
release-1.0

Fail pipeline when tags do not match:

^v[0-9]+.[0-9]+.[0-9]+$

Generate workflow:
semantic-version-validation.yml

---

## Enterprise Build Pipeline

Create:

build.yml

Requirements:

* Matrix builds
* Parallel execution
* Docker Buildx
* Multi-architecture support
* Build caching
* Dependency caching
* Layer optimization
* Build summaries

Platforms:

* linux/amd64
* linux/arm64

---

## Enterprise Caching

Implement:

GitHub Actions Cache
Docker BuildKit Cache
Registry Cache

Requirements:

cache-from:
type=gha

cache-to:
type=gha,mode=max

Expected Improvement:
50-70% faster rebuilds.

Generate full implementation.

---

## Failure Recovery

Implement retry strategy.

Requirements:

* Retry failed image builds
* Retry failed registry pushes
* Continue unaffected services
* Isolate failed service builds

Generate:

retry strategy
failure reporting
workflow summaries

Provide example YAML.

---

## Security Scanning

Generate security workflows:

security.yml

Requirements:

Trivy
Grype
Docker Scout

Scan:

* Containers
* Dockerfiles
* Dependencies
* Secrets

Fail on:
Critical vulnerabilities

Warn on:
High vulnerabilities

Generate complete workflow.

---

## Dependency Management

Generate:

dependabot.yml

Requirements:

GitHub Actions updates
Docker image updates
Security patch automation

Weekly schedule.

---

## Automated Changelog Generation

Generate:

changelog.yml

Requirements:

Conventional Commits

Examples:

feat:
fix:
docs:
refactor:
security:
build:
ci:

Automatically generate:

CHANGELOG.md

on release creation.

Use:
git-cliff

Generate complete configuration.

---

## Automated Release Notes

Generate:

release.yml

Requirements:

Auto-generate release notes from commits.

Sections:

Features
Fixes
Security
Infrastructure
Breaking Changes
Dependencies

Publish automatically when tag is created.

Output:
Professional GitHub Release page.

---

## Release Promotion Pipeline

Implement:

dev
staging
production

Flow:

Pull Request
→ Build
→ Security Scan
→ Integration Tests
→ Staging
→ Release Approval
→ Production Release

Generate all workflows.

---

## Quality Gates

Required before release:

* Linting
* Dockerfile validation
* YAML validation
* Security scans
* Build success
* Integration tests

Block release on failure.

---

## GitHub Environments

Generate:

Development
Staging
Production

Include:
environment protection rules
manual approval gates
secret isolation

---

## Supply Chain Security

Implement:

SBOM generation
Cosign image signing
SLSA provenance
Artifact attestations

Generate workflows.

---

## Monitoring

Generate:

release metrics
deployment metrics
security metrics

Include workflow summaries and GitHub dashboards.

---

## Documentation Requirements

Generate:

1. Architecture diagram
2. CI/CD overview
3. Release process
4. Versioning policy
5. Registry strategy
6. Security policy
7. Contribution guide

Output format:

* Production-ready YAML
* Enterprise DevSecOps standards
* GitHub Actions compatible
* Fully documented
* Scalable to 100+ services
* Suitable for open-source and enterprise environments
