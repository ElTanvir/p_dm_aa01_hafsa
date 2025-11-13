package handler

import (
	"context"
	db "p_dm_aa01_hafsa/internal/db/sqlc"
	"p_dm_aa01_hafsa/internal/modules/theme/domain"
	"p_dm_aa01_hafsa/internal/modules/theme/usecase"
)

type ThemeHandler interface {
	GetVariablesByType(ctx context.Context, varType string) ([]db.CssVariable, error)
	UpdateVariable(ctx context.Context, req domain.CreateOrUpdateVariableRequest) error
}

type themeHandler struct {
	useCase usecase.ThemeUseCase
}

func NewThemeHandler(useCase usecase.ThemeUseCase) ThemeHandler {
	return &themeHandler{
		useCase: useCase,
	}
}
