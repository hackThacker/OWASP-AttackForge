<div align="center">

# 🛡️ OWASP AttackForge

### A Containerized Multi App Vulnerability Range for Offensive Security Training

[![License](https://img.shields.io/badge/license-Apache%202.0-green.svg)](LICENSE)
[![Stars](https://img.shields.io/github/stars/hackThacker/OWASP-AttackForge?color=yellow)](https://github.com/hackThacker/OWASP-AttackForge/stargazers)
[![Forks](https://img.shields.io/github/forks/hackThacker/OWASP-AttackForge?color=blue)](https://github.com/hackThacker/OWASP-AttackForge/network/members)
[![Issues](https://img.shields.io/github/issues/hackThacker/OWASP-AttackForge?color=red)](https://github.com/hackThacker/OWASP-AttackForge/issues)

</div>

<p align="center">
  <a href="#-what-is-this">What is This</a> ·
  <a href="#-why-this-repo">Why This Repo</a> ·
  <a href="#-topics-covered">Topics Covered</a> ·
  <a href="#-quick-start">Quick Start</a> ·
  <a href="#-lab-applications-and-credentials">Lab Apps</a> ·
  <a href="#-who-is-this-for">Who Is This For</a> ·
  <a href="#-license">License</a> ·
  <a href="#-support">Support</a>
</p>

---

## 🔍 What is This

**OWASP AttackForge** is a self contained cyber range that boots nine industry standard vulnerable web applications behind a single Nginx reverse proxy. One `docker compose up` spins up DVWA, Mutillidae II, bWAPP, XVWA, VWA, Juice Shop, WebGoat, WebWolf, and a Tomcat manager console, each in its own hardened, non root container, with database provisioning handled automatically.

```text
Browser
   │  HTTPS (443) / HTTP (80)
   ▼
┌─────────────────────────────┐
│         Nginx Proxy         │
└──┬───┬───┬───┬───┬───┬───┬──┘
   │   │   │   │   │   │   │
 DVWA bWAPP XVWA VWA Mutillidae JuiceShop WebGoat/WebWolf  Tomcat
   │   │   │   │   │
   └───┴───┴───┴───┴──► MariaDB (private, port 3306)
```

> [!CAUTION]
> **Deliberately vulnerable environment.** Every application here ships with real, unpatched RCE, SQL injection, and path traversal flaws by design. Run it only on an isolated host. Never expose ports 80/443 to the public internet or deploy on a shared or production network.

---

## 🤔 Why This Repo

| Without OWASP AttackForge | With OWASP AttackForge |
| :--- | :--- |
| Spin up and patch 9 separate VMs or local installs | One script and one `docker compose up` brings up the whole range |
| Conflicting PHP, Java, and Node runtime versions on one host | Each app runs isolated in its own container |
| Manually create databases and schemas per app | An initializer service provisions the database on first boot |
| Apps exposed directly on their own host ports | Nginx is the only service that touches the host network |
| Containers commonly run as root | Every container runs as a dedicated non root service account |
| No resource limits, one runaway app can choke the host | CPU and memory caps are set per container |

---

## 📚 Topics Covered

* OWASP Top 10 vulnerability classes: SQL injection, XSS, CSRF, IDOR, broken authentication, SSRF, path traversal, command injection, insecure deserialization, and security misconfiguration
* Multi container orchestration and service health checks with Docker Compose
* Nginx as a reverse proxy with auto generated wildcard SSL certificates
* Database isolation and automatic schema provisioning with MariaDB
* Non root container hardening across PHP, Java, and Node.js runtimes
* Private bridge networking that exposes only one entry point to the host
* Resource constrained deployment for low spec training hardware

---

## 🚀 Quick Start

```bash
git clone https://github.com/hackThacker/OWASP-AttackForge.git
cd OWASP-AttackForge
sudo ./docker-install.sh
```

The installer copies `.env.example` to `.env`, builds and starts all 12 containers, and maps the lab domains into your hosts file on Linux automatically.

To run it manually instead:

```bash
cp .env.example .env
docker compose up --build -d
docker ps
```

**Requirements:** Docker Engine 20.10+, Docker Compose V2, 8 GB RAM minimum (12 GB recommended), 5 GB free disk space. Windows users need WSL2 with Docker Desktop and must map the domains by hand in `C:\Windows\System32\drivers\etc\hosts`. On Linux or macOS, edit `/etc/hosts` (the install script does this for you on Linux).

```text
# Add to your hosts file if not using docker-install.sh
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

## 🧪 Lab Applications and Credentials

| Application | URL | Username | Password |
| :--- | :--- | :--- | :--- |
| Mutillidae II | `https://mutillidae.hackthacker.lab` | `admin` | `adminpass` |
| DVWA | `https://dvwa.hackthacker.lab/login.php` | `admin` | `password` |
| bWAPP | `https://bwapp.hackthacker.lab/login.php` | `bee` | `bug` |
| XVWA | `https://xvwa.hackthacker.lab` | `admin` | `admin` |
| VWA | `https://vwa.hackthacker.lab` | `admin` | `password` |
| Juice Shop | `https://juiceshop.hackthacker.lab` | register in UI | register in UI |
| WebGoat | `https://webgoat.hackthacker.lab/WebGoat/login` | create in UI | create in UI |
| WebWolf | `https://webwolf.hackthacker.lab/login` | WebGoat account | WebGoat account |
| Tomcat Manager | `https://tomcat.hackthacker.lab` | `hackthacker` | `hackthacker` |

---

## 📁 Repo Files

```text
OWASP-AttackForge/
├── docker/
│   ├── db/
│   │   └── init.sql
│   ├── initializer/
│   │   ├── Dockerfile
│   │   └── init.sh
│   ├── juiceshop/
│   │   └── Dockerfile
│   ├── nginx/
│   │   ├── Dockerfile
│   │   ├── default.conf.template
│   │   └── entrypoint.sh
│   ├── php/
│   │   ├── Dockerfile.php
│   │   ├── disable_strict_mysqli.php
│   │   ├── entrypoint.sh
│   │   └── owasp-lab.ini
│   ├── tomcat/
│   │   ├── Dockerfile
│   │   ├── entrypoint.sh
│   │   └── tomcat-users.xml
│   ├── webgoat/
│   │   └── Dockerfile
│   └── webwolf/
│       └── Dockerfile
├── .dockerignore
├── .env.example
├── docker-check.sh
├── docker-compose.dev.yml
├── docker-compose.yml
├── docker-install.sh
├── docker-uninstall.sh
└── README.md
```

---

## 🎯 Who is This For

```text
Security trainers running classroom or workshop labs
Students practicing OWASP Top 10 exploitation hands on
Red team members rehearsing web attack chains before an engagement
Blue team analysts generating live exploitation traffic for detection testing
Bug bounty hunters sharpening manual testing skills before live targets
Anyone who wants 9 vulnerable apps running locally without the DevOps pain
```

---

## 🔧 Troubleshooting

* **Containers stuck restarting:** check `docker logs hackthacker-labs-nginx` and confirm no other service (IIS, Apache, Skype) already holds ports 80 or 443.
* **Domains will not resolve:** run `ping juiceshop.hackthacker.lab`, it should resolve to `127.0.0.1`. If not, recheck your hosts file entries.
* **WebGoat or WebWolf crash silently:** they need up to 1.5 GB combined. On Windows, raise the WSL2 memory cap in `%USERPROFILE%\.wslconfig` with a `memory=4GB` entry under `[wsl2]`.

---

## 📊 Repo Stats

![Top Language](https://img.shields.io/github/languages/top/hackThacker/OWASP-AttackForge)
![Languages](https://img.shields.io/github/languages/count/hackThacker/OWASP-AttackForge)
![Last Commit](https://img.shields.io/github/last-commit/hackThacker/OWASP-AttackForge)
![Repo Size](https://img.shields.io/github/repo-size/hackThacker/OWASP-AttackForge)

---

## ⚖️ License

Distributed under the Apache License 2.0. See `LICENSE` for details.

---

## 💬 Support

🐛 Found a bug or have a question? Open an issue on the [Issues page](https://github.com/hackThacker/OWASP-AttackForge/issues).

---

<div align="center">

Made with ❤️ by [hackthacker](https://github.com/hackThacker)

</div>
