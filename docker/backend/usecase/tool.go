package usecase

import (
	"context"

	"github.com/hackThacker/OWASP-AttackForge/backend/domain"
)

type toolUsecase struct {
	repo domain.ToolRepository
}

func NewToolUsecase(repo domain.ToolRepository) domain.ToolUsecase {
	return &toolUsecase{repo: repo}
}

func (u *toolUsecase) GetTools(ctx context.Context) ([]*domain.Tool, error) {
	return u.repo.List(ctx)
}

func (u *toolUsecase) StartTool(ctx context.Context, subdomain string) error {
	return u.repo.Start(ctx, subdomain)
}

func (u *toolUsecase) StopTool(ctx context.Context, subdomain string) error {
	return u.repo.Stop(ctx, subdomain)
}

func (u *toolUsecase) RestartTool(ctx context.Context, subdomain string) error {
	return u.repo.Restart(ctx, subdomain)
}
