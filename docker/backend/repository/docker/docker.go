package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/hackThacker/OWASP-AttackForge/backend/domain"
)

type dockerRepository struct {
	cli *client.Client
}

func NewDockerRepository() (domain.ToolRepository, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &dockerRepository{cli: cli}, nil
}

var staticMetadata = map[string]domain.Tool{
	"mutillidae": {
		Name:        "Mutillidae II",
		Subdomain:   "mutillidae",
		Icon:        "fas fa-bug",
		Description: "OWASP Mutillidae II is a free, open-source, deliberately vulnerable web application providing a target for web-security training.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "admin", Password: "adminpass"},
	},
	"dvwa": {
		Name:        "DVWA",
		Subdomain:   "dvwa",
		Icon:        "fas fa-shield-alt",
		Description: "Damn Vulnerable Web Application (DVWA) is a PHP/MySQL web application that is damn vulnerable, aiding security professionals.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "admin", Password: "password"},
	},
	"bwapp": {
		Name:        "bWAPP",
		Subdomain:   "bwapp",
		Icon:        "fas fa-cookie-bite",
		Description: "A buggy web application that is free and open-source, containing over 100 web vulnerabilities.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "bee", Password: "bug"},
	},
	"xvwa": {
		Name:        "XVWA",
		Subdomain:   "xvwa",
		Icon:        "fas fa-code",
		Description: "Xtreme Vulnerable Web Application (XVWA) is a badly coded PHP/MySQL web application to learn application security.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "admin", Password: "admin"},
	},
	"vwa": {
		Name:        "VWA",
		Subdomain:   "vwa",
		Icon:        "fas fa-spider",
		Description: "Vulnerable Web Application is a security training application with various web vulnerabilities.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "admin", Password: "password"},
	},
	"juiceshop": {
		Name:        "Juice Shop",
		Subdomain:   "juiceshop",
		Icon:        "fas fa-wine-glass-alt",
		Description: "Probably the most modern and sophisticated insecure web application, written in Angular, Express, and Node.js.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "admin@juice-sh.op", Password: "admin123"},
	},
	"webgoat": {
		Name:        "WebGoat",
		Subdomain:   "webgoat",
		Icon:        "fas fa-mountain",
		Description: "WebGoat is a deliberately insecure application that allows developers to test vulnerabilities commonly found in Java-based applications.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/WebGoat/",
		Category:    "Java Application",
		Credentials: domain.Credentials{Username: "guest", Password: "guest"},
	},
	"webwolf": {
		Name:        "WebWolf",
		Subdomain:   "webwolf",
		Icon:        "fas fa-wolf-pack-battalion",
		Description: "A companion application for WebGoat, used to simulate attacker actions like receiving emails or hosting files.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/login",
		Category:    "Java Application",
		Credentials: domain.Credentials{Username: "guest", Password: "guest"},
	},
	"tomcat": {
		Name:        "Tomcat Manager",
		Subdomain:   "tomcat",
		Icon:        "fas fa-server",
		Description: "Vulnerable Apache Tomcat server instance with manager console access enabled for credential brute-forcing.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Infrastructure",
		Credentials: domain.Credentials{Username: "admin", Password: "hackthacker"},
	},
	"wrongsecrets": {
		Name:        "WrongSecrets",
		Subdomain:   "wrongsecrets",
		Icon:        "fas fa-user-secret",
		Description: "A vulnerable application designed to teach how to find and prevent secrets exposure in cloud and container setups.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Cloud Security",
		Credentials: domain.Credentials{Username: "None", Password: "N/A"},
	},
	"securityshepherd": {
		Name:        "Security Shepherd",
		Subdomain:   "securityshepherd",
		Icon:        "fas fa-user-shield",
		Description: "A web and mobile application security training platform designed to foster and improve security awareness.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "admin", Password: "password"},
	},
	"vulnerableapp": {
		Name:        "VulnerableApp",
		Subdomain:   "vulnerableapp",
		Icon:        "fas fa-cubes",
		Description: "A facade application that consolidates various vulnerable applications written in different languages (Java, PHP, JSP).",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "None", Password: "N/A"},
	},
	"crapi": {
		Name:        "crAPI",
		Subdomain:   "crapi",
		Icon:        "fas fa-exchange-alt",
		Description: "Completely Ridiculous API (crAPI) is designed to teach API security vulnerabilities in modern microservice architectures.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "API Security",
		Credentials: domain.Credentials{Username: "user@example.com", Password: "password"},
	},
	"brokencrystals": {
		Name:        "BrokenCrystals",
		Subdomain:   "brokencrystals",
		Icon:        "fas fa-gem",
		Description: "A benchmark application containing modern security vulnerabilities, with integration for Ollama LLM and ChromaDB vector store.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "API Security",
		Credentials: domain.Credentials{Username: "admin", Password: "admin"},
	},
	"dvws": {
		Name:        "DVWS Node",
		Subdomain:   "dvws",
		Icon:        "fas fa-network-wired",
		Description: "Damn Vulnerable Web Services is a vulnerable web services application for learning API and web service security.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "API Security",
		Credentials: domain.Credentials{Username: "admin", Password: "admin"},
	},
	"zerohealth": {
		Name:        "Zero Health",
		Subdomain:   "zerohealth",
		Icon:        "fas fa-heartbeat",
		Description: "A vulnerable medical application stack with a separate client frontend and a server backend API.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "Web Application",
		Credentials: domain.Credentials{Username: "admin", Password: "admin"},
	},
	"restaurant": {
		Name:        "Damn Vulnerable RESTaurant",
		Subdomain:   "restaurant",
		Icon:        "fas fa-utensils",
		Description: "A REST API game designed to teach developers and security practitioners how to secure REST APIs.",
		Protocols:   []string{"https"},
		Port:        "8443",
		URI:         "/",
		Category:    "API Security",
		Credentials: domain.Credentials{Username: "admin", Password: "admin"},
	},
}

func (r *dockerRepository) List(ctx context.Context) ([]*domain.Tool, error) {
	containers, err := r.cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	// Create map of running state & uptime
	states := make(map[string]struct {
		running bool
		uptime  string
	})

	for _, c := range containers {
		// Clean container name (typically starts with /)
		name := strings.TrimPrefix(c.Names[0], "/")
		if !strings.HasPrefix(name, "hackthacker-labs-") {
			continue
		}
		subdomain := strings.TrimPrefix(name, "hackthacker-labs-")

		// Handle specific multi-container target services
		if subdomain == "zerohealth-client" {
			subdomain = "zerohealth"
		} else if subdomain == "brokencrystals-app" {
			subdomain = "brokencrystals"
		} else if subdomain == "vulnerableapp-facade" {
			subdomain = "vulnerableapp"
		} else if subdomain == "crapi-web" {
			subdomain = "crapi"
		} else if subdomain == "dvws-node" {
			subdomain = "dvws"
		} else if subdomain == "restaurant-app" {
			subdomain = "restaurant"
		}

		running := c.State == "running"
		uptime := "Stopped"
		if running {
			// Calculate uptime string
			uptime = getUptimeString(c.Status)
		}

		states[subdomain] = struct {
			running bool
			uptime  string
		}{
			running: running,
			uptime:  uptime,
		}
	}

	var list []*domain.Tool
	for sub, meta := range staticMetadata {
		toolCopy := meta
		state, ok := states[sub]
		if ok {
			toolCopy.Stopped = !state.running
			toolCopy.Uptime = state.uptime
		} else {
			toolCopy.Stopped = true
			toolCopy.Uptime = "Stopped"
		}
		list = append(list, &toolCopy)
	}

	return list, nil
}

func (r *dockerRepository) Start(ctx context.Context, subdomain string) error {
	containerID, err := r.findContainerID(ctx, subdomain)
	if err != nil {
		return err
	}
	return r.cli.ContainerStart(ctx, containerID, container.StartOptions{})
}

func (r *dockerRepository) Stop(ctx context.Context, subdomain string) error {
	containerID, err := r.findContainerID(ctx, subdomain)
	if err != nil {
		return err
	}
	stopTimeout := 10
	return r.cli.ContainerStop(ctx, containerID, container.StopOptions{Timeout: &stopTimeout})
}

func (r *dockerRepository) Restart(ctx context.Context, subdomain string) error {
	containerID, err := r.findContainerID(ctx, subdomain)
	if err != nil {
		return err
	}
	return r.cli.ContainerRestart(ctx, containerID, container.StopOptions{})
}

func (r *dockerRepository) findContainerID(ctx context.Context, subdomain string) (string, error) {
	containers, err := r.cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return "", err
	}

	// Normalize subdomains to match container suffixes
	expectedSuffix := subdomain
	switch subdomain {
	case "zerohealth":
		expectedSuffix = "zerohealth-client"
	case "brokencrystals":
		expectedSuffix = "brokencrystals-app"
	case "vulnerableapp":
		expectedSuffix = "vulnerableapp-facade"
	case "crapi":
		expectedSuffix = "crapi-web"
	case "dvws":
		expectedSuffix = "dvws-node"
	case "restaurant":
		expectedSuffix = "restaurant-app"
	}

	expectedName := fmt.Sprintf("hackthacker-labs-%s", expectedSuffix)

	for _, c := range containers {
		for _, name := range c.Names {
			cleanName := strings.TrimPrefix(name, "/")
			if cleanName == expectedName {
				return c.ID, nil
			}
		}
	}
	return "", fmt.Errorf("container not found for subdomain %s", subdomain)
}

func getUptimeString(status string) string {
	if strings.HasPrefix(status, "Up ") {
		return status
	}
	return "Running"
}
