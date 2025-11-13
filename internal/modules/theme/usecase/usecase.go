package usecase

import (
	"context"
	db "p_dm_aa01_hafsa/internal/db/sqlc"
	"p_dm_aa01_hafsa/internal/modules/theme/domain"
	"p_dm_aa01_hafsa/internal/modules/theme/repository"
)

type ThemeUseCase interface {
	GenerateCSS(ctx context.Context) (string, error)
	InvalidateCache(ctx context.Context) error
	GetAllVariables(ctx context.Context) ([]db.CssVariable, error)
	GetVariablesByType(ctx context.Context, varType string) ([]db.CssVariable, error)
	GetCSSVariableByName(ctx context.Context, name string) (db.CssVariable, error)
	CreateVariable(ctx context.Context, req domain.CreateOrUpdateVariableRequest) error
	UpdateVariable(ctx context.Context, req domain.CreateOrUpdateVariableRequest) error
	DeleteVariable(ctx context.Context, id string) error
}

type themeUseCase struct {
	repo repository.ThemeRepository
}

func NewThemeUseCase(repo repository.ThemeRepository) ThemeUseCase {
	return &themeUseCase{
		repo: repo,
	}
}
