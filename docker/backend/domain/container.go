package domain

import "context"

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Tool struct {
	Name        string      `json:"name"`
	Subdomain   string      `json:"subdomain"`
	Icon        string      `json:"icon"`
	Description string      `json:"description"`
	Protocols   []string    `json:"protocols"`
	Port        string      `json:"port"`
	URI         string      `json:"uri"`
	Category    string      `json:"category"`
	Credentials Credentials `json:"credentials"`
	Uptime      string      `json:"uptime"`
	Stopped     bool        `json:"stopped"`
}

type ToolRepository interface {
	List(ctx context.Context) ([]*Tool, error)
	Start(ctx context.Context, subdomain string) error
	Stop(ctx context.Context, subdomain string) error
	Restart(ctx context.Context, subdomain string) error
}

type ToolUsecase interface {
	GetTools(ctx context.Context) ([]*Tool, error)
	StartTool(ctx context.Context, subdomain string) error
	StopTool(ctx context.Context, subdomain string) error
	RestartTool(ctx context.Context, subdomain string) error
}
