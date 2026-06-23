# 🛡️ OWASP AttackForge

[![Docker](https://img.shields.io/badge/docker-compatible-blue.svg?logo=docker&logoColor=white)](https://www.docker.com)
[![License](https://img.shields.io/badge/license-Apache%202.0-green.svg)](LICENSE)
[![Security](https://img.shields.io/badge/security-deliberately%20vulnerable-red.svg)](https://owasp.org)
[![Platform](https://img.shields.io/badge/platform-linux%20%7C%20windows-lightgrey.svg)](https://github.com)

A high-performance, security-hardened, and containerized multi-application cyber range and vulnerability lab environment developed for educational training, penetration testing practice, and security research.

---

## 📋 Table of Contents

* [Overview](#overview)
* [Features](#features)
* [Architecture](#architecture)
* [Requirements](#requirements)
* [Supported Platforms](#supported-platforms)
* [Installation](#installation)
  * [Windows Setup](#windows-setup)
  * [Linux Setup](#linux-setup)
  * [Docker Deployment](#docker-deployment)
* [Configuration](#configuration)
* [Lab Environment](#lab-environment)
* [Network Architecture](#network-architecture)
* [Directory Structure](#directory-structure)
* [Security Considerations](#security-considerations)
* [Troubleshooting](#troubleshooting)
* [Development](#development)
* [Contributing](#contributing)
* [Roadmap](#roadmap)
* [License](#license)
* [Acknowledgements](#acknowledgements)

---

## 🔍 Overview

**OWASP AttackForge** is an enterprise-grade local simulation range that consolidates nine industry-standard vulnerable web applications into a single unified orchestration stack. Behind a secure Nginx reverse proxy, the environment routes custom domain requests to isolated Node.js, Java, Tomcat, and PHP application runtimes.

This workspace enables cybersecurity students, developers, and red/blue teams to simulate advanced security scenarios (like web exploitation, network pivots, database injection, and log auditing) in a secure, self-contained environment without exposing the host to external vulnerabilities.

---

## ⚡ Features

*   **Unified Orchestration:** Launch 9 distinct vulnerable web applications with a single command.
*   **Production-Grade Reverse Proxy:** Nginx routes traffic securely with auto-generated wildcard SSL certificates.
*   **Hardened Non-Root Security:** Containers operate as non-root service accounts (`nginx`, `node`, `tomcat`, `webgoat`, `www-data`, and `nobody`) to prevent host compromise.
*   **Automatic Database Provisioning:** Database tables, schemas, and credentials are set up dynamically on startup by an asynchronous initializer service.
*   **CPU & Memory Capping:** Resource limit constraints prevent denial-of-service loops from exhausting host memory or processor cycles.
*   **Custom Local Domains:** Access applications cleanly via subdomains of `*.hackthacker.lab`.

---

## 🏛️ Architecture

The architecture isolates the backend applications within a private bridge network while exposing only Nginx to the host.

### Conceptual Routing Diagram

```text
User Browser / CLI
       │
       ▼ (HTTPS: Port 443 / HTTP: Port 80)
┌────────────────────────────────────────────────────────┐
│                      Nginx Proxy                       │
└───┬──────────────┬──────────────┬──────────────────┬───┘
    │              │              │                  │
    ▼              ▼              ▼                  ▼
┌───────┐      ┌───────┐      ┌───────┐          ┌───────┐
│ DVWA  │      │ bWAPP │      │ XVWA  │          │ VWA   │
│ (PHP) │      │ (PHP) │      │ (PHP) │          │ (PHP) │
└───┬───┘      └───┬───┘      └───┬───┘          └───┬───┘
    │              │              │                  │
    └──────────────┼──────────────┼──────────────────┘
                   ▼
             ┌───────────┐
             │  MariaDB  │ (Private Port 3306)
             └───────────┘
```

### Full Architecture Map

```text
Host Machine Interface (Localhost)
       │
       ├─► [Port 80/443] ──────► Nginx Router (Proxy Server)
       │                            │
       │  ┌─────────────────────────┼────────────────────────┐
       │  │ (hackthacker Network)   │                        │
       │  │                         ├─► mutillidae:8080      │
       │  │                         ├─► dvwa:8080 ─────────┐ │
       │  │                         ├─► bwapp:8080 ────────┼─┤
       │  │                         ├─► xvwa:8080 ─────────┼─┤
       │  │                         ├─► vwa:8080 ──────────┼─┤
       │  │                         ├─► juiceshop:3000     │ │
       │  │                         ├─► webgoat:8080 ◄───┐ │ │
       │  │                         ├─► webwolf:9090 ────┘ │ │
       │  │                         └─► tomcat:8080        │ │
       │  │                                                ▼ │
       │  │                                              [db]│
       │  └──────────────────────────────────────────────────┘
```

---

## ⚙️ Requirements

### Hardware Requirements
*   **CPU:** Dual-core processor or higher (Intel VT-x / AMD-V virtualization support enabled).
*   **RAM:** 8 GB minimum (12 GB or higher recommended).
*   **Storage:** 5 GB of free SSD space.

### Software Requirements
*   **Docker Engine:** Version 20.10.0 or higher.
*   **Docker Compose:** Compose V2 support (CLI plugin).
*   **Git Client:** Any standard command-line Git.

### Supported Operating Systems
*   Windows 10 / 11 (with WSL2 backend configured)
*   Windows Server 2022 or higher
*   Ubuntu Linux (20.04 LTS / 22.04 LTS / 24.04 LTS)
*   Debian Linux (11 / 12)
*   Kali Linux
*   Rocky Linux / RHEL / Fedora
*   Arch Linux

---

# 🚀 Installation

## Windows Setup

### Prerequisites
1. Ensure **WSL2** (Windows Subsystem for Linux) is installed. Run `wsl --install` in PowerShell.
2. Enable **Virtual Machine Platform** and **Hyper-V** features in Windows Features.

### Install Docker Desktop
1. Download and install [Docker Desktop for Windows](https://www.docker.com/products/docker-desktop/).
2. In Settings -> General, ensure **"Use the WSL 2 based engine"** is checked.
3. In Settings -> Resources -> WSL Integration, enable integration for your default distro.

### Clone Repository
Open PowerShell and clone the repository:
```powershell
git clone https://github.com/HackThackerLabs/OWASP-AttackForge.git
cd OWASP-AttackForge
```

### Start Environment
```powershell
./docker-install.sh
```
*Note: If `./docker-install.sh` fails due to execution policy or script restrictions, start the container stack manually using:*
```powershell
docker compose up -d
```

### Verify Services
```powershell
docker ps
```
Ensure all containers show `Up (healthy)`.

### Access Applications
To resolve the custom domains on Windows, refer to the [Hosts File Configuration](#hosts-file-configuration) section.

---

## Linux Setup

### Install Dependencies

**Ubuntu / Debian / Kali Linux:**
```bash
sudo apt update
sudo apt install docker.io docker-compose-plugin git curl -y
sudo systemctl enable --now docker
```

**Fedora / Rocky Linux / RHEL:**
```bash
sudo dnf install docker docker-compose git curl -y
sudo systemctl enable --now docker
```

**Arch Linux:**
```bash
sudo pacman -S docker docker-compose git curl --noconfirm
sudo systemctl enable --now docker
```

### Clone Repository
```bash
git clone https://github.com/HackThackerLabs/OWASP-AttackForge.git
cd OWASP-AttackForge
```

### Start Environment
Deploy with the installer script (will automatically configure host DNS mapping):
```bash
sudo ./docker-install.sh
```
Or start manually via raw compose:
```bash
docker compose up -d
```

### Verify Deployment
```bash
docker ps
```

---

## Docker Deployment

The AttackForge stack maps private ports to restrict direct network exposure. The Nginx reverse proxy serves as the sole gateway.

**Orchestrate and Manage Commands:**
```bash
# Pull the latest base images
docker compose pull

# Compile custom configurations and run in background
docker compose up --build -d

# Trace container startup logs
docker compose logs -f

# Verify service health and metrics
./docker-check.sh

# Tear down the stack and delete database persistent storage volumes
./docker-uninstall.sh
```

---

## 🔧 Configuration

### Environment Variables
Adjust parameters in the `.env` configuration file (generated from `.env.example` at installation):

| Variable | Default Value | Description |
| :--- | :--- | :--- |
| `DOMAIN` | `hackthacker.lab` | Base subdomain suffix |
| `SSL_ORG` | `OWASP AttackForge` | Subject organization in self-signed SSL certs |
| `DB_USER` | `hackthacker` | Default database application username |
| `DB_PASS` | `hackthacker` | Default database application password |
| `MARIADB_ROOT_PASSWORD`| `hackthacker` | Root administrator password for MariaDB |

### Hosts File Configuration
To access the applications by their domains, append the following mappings to your hosts configuration file:

**Windows System Path:** `C:\Windows\System32\drivers\etc\hosts` (Open Notepad as Administrator)  
**Linux / macOS Path:** `/etc/hosts` (Edit using `sudo nano /etc/hosts`)

```text
# HackThackerLabs AttackForge Cyber Range
127.0.0.1 mutillidae.hackthacker.lab
127.0.0.1 dvwa.hackthacker.lab
127.0.0.1 bwapp.hackthacker.lab
127.0.0.1 xvwa.hackthacker.lab
127.0.0.1 vwa.hackthacker.lab
127.0.0.1 juiceshop.hackthacker.lab
127.0.0.1 webgoat.hackthacker.lab
127.0.0.1 webwolf.hackthacker.lab
127.0.0.1 tomcat.hackthacker.lab
```

---

## 🎯 Lab Environment

| Application | Domain Access URL | Default User | Default Password | Port |
| :--- | :--- | :--- | :--- | :--- |
| **Mutillidae II** | [https://mutillidae.hackthacker.lab](https://mutillidae.hackthacker.lab) | `admin` | `adminpass` | `443` |
| **DVWA** | [https://dvwa.hackthacker.lab/login.php](https://dvwa.hackthacker.lab/login.php) | `admin` | `password` | `443` |
| **bWAPP** | [https://bwapp.hackthacker.lab/login.php](https://bwapp.hackthacker.lab/login.php) | `bee` | `bug` | `443` |
| **XVWA** | [https://xvwa.hackthacker.lab](https://xvwa.hackthacker.lab) | `admin` | `admin` | `443` |
| **VWA** | [https://vwa.hackthacker.lab](https://vwa.hackthacker.lab) | `admin` | `password` | `443` |
| **Juice Shop** | [https://juiceshop.hackthacker.lab](https://juiceshop.hackthacker.lab) | *(Register in UI)* | *(Register in UI)* | `443` |
| **WebGoat** | [https://webgoat.hackthacker.lab/WebGoat/login](https://webgoat.hackthacker.lab/WebGoat/login) | *(Create in UI)* | *(Create in UI)* | `443` |
| **WebWolf** | [https://webwolf.hackthacker.lab/login](https://webwolf.hackthacker.lab/login) | *(WebGoat account)*| *(WebGoat account)*| `443` |
| **Tomcat Console**| [https://tomcat.hackthacker.lab](https://tomcat.hackthacker.lab) | `hackthacker` | `hackthacker` | `443` |

---

## 📁 Directory Structure

```text
OWASP-AttackForge/
├── docker/
│   ├── db/                 # SQL database setup and permissions
│   ├── initializer/        # HTTP API migration auto-initializer scripts
│   ├── juiceshop/          # NodeJS non-root context configs
│   ├── nginx/              # Nginx template engine configurations
│   ├── php/                # Hardened bookworm apache PHP environment
│   ├── tomcat/             # Debian temurin JRE tomcat files
│   ├── webgoat/            # WebGoat Alpine JRE security limits
│   └── webwolf/            # WebWolf database links configuration
├── docker-compose.yml      # Core production deployment configuration
├── docker-compose.dev.yml  # Development host port mappings overrides
├── docker-check.sh         # Active endpoint curl status validator
├── docker-install.sh       # Deploy stack and hosts mapper script
├── docker-uninstall.sh     # Cleanup stack and volume wiper script
├── README.md               # Main repository documentation
└── .env.example            # Environment configuration template
```

---

## ⚠️ Security Considerations

> [!CAUTION]
> **DELIBERATELY VULNERABLE ENVIRONMENT**
> *   This project contains software with critical remote code execution (RCE), SQL injection, and path traversal vulnerabilities.
> *   **DO NOT** deploy this stack on public, shared, production, or untrusted networks.
> *   **DO NOT** expose Nginx ports 80/443 directly to the public internet.
> *   Ensure the docker daemon host is secured and isolated within host-only loopbacks.

---

## 🔧 Troubleshooting

### Containers Not Starting / Restarting Loops
1. Check Nginx logs to confirm if any host is unreachable:
   ```bash
   docker logs attackforge-nginx
   ```
2. Verify that there are no port binding conflicts on the host (like local IIS or Apache instances listening on ports 80 or 443).

### DNS Resolution Issues
* If you cannot connect to the applications using browser domains, check that the host mappings are loaded correctly:
  ```bash
  ping juiceshop.hackthacker.lab
  ```
  It should resolve successfully to `127.0.0.1`. If not, review the hosts mappings editing steps.

### Port Conflicts
* In Windows, port `80` is often locked by standard OS services (e.g. Web Deployment Agent, IIS, or Skype). Run `netstat -ano | findstr :80` in command prompt to identify the process ID causing conflicts, stop the respective service, and try again.

### Resource Constraints
* WebGoat and WebWolf run on JVMs and require up to `1.5 GB` of memory combined. If they crash silently, check your Docker RAM settings and increase WSL2 resource limits by creating a `%USERPROFILE%\.wslconfig` file:
  ```text
  [wsl2]
  memory=4GB
  ```

---

## 🛠️ Development

To override the production network boundaries and expose all microservice ports directly for isolated debugging, run the dev overrides command:
```bash
docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

---

## 🤝 Contributing

Contributions to OWASP AttackForge are welcome!
1. Fork the Project Repository.
2. Create a Feature Branch (`git checkout -b feature/AmazingFeature`).
3. Commit changes (`git commit -m 'Add some AmazingFeature'`).
4. Push to the Branch (`git push origin feature/AmazingFeature`).
5. Open a Pull Request.

---

## 🗺️ Roadmap

*   [ ] Add Kali Linux pentest client image container within the orchestration.
*   [ ] Configure local logging pipeline using Grafana/Prometheus logs collector.
*   [ ] Introduce OWASP security audit automation scripts.
*   [ ] Expand documentation on specific vulnerabilities covered in the range.

---

## 📄 License

Distributed under the Apache License 2.0. See `LICENSE` for details.

---

## 💖 Acknowledgements

*   Upstream vulnerability project creators: [DVWA](https://github.com/digininja/DVWA), [Mutillidae II](https://github.com/webpwnized/mutillidae), [bWAPP](https://github.com/iMoon07/bWAPPs), [XVWA](https://github.com/s4n7h0/xvwa), [VWA](https://github.com/hummingbirdscyber/Vulnerable-Web-Application), [Juice Shop](https://github.com/juice-shop/juice-shop), [WebGoat & WebWolf](https://github.com/WebGoat/WebGoat).
*   Inspired by the work of [iMoon](https://github.com/iMoon07) and [Taro Lay](https://github.com/tarolay) on the original host installer.
# OWASP-AttackForge
# OWASP-AttackForge
