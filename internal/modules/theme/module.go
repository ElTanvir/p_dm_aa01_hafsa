package theme

import (
	"context"
	"p_dm_aa01_hafsa/internal/modules/theme/repository"
	"p_dm_aa01_hafsa/internal/modules/theme/usecase"
	"p_dm_aa01_hafsa/internal/server"
)

func Init(server *server.Server) {
	useCase := getUseCase(server)
	useCase.InvalidateCache(context.Background())
}

func getUseCase(s *server.Server) usecase.ThemeUseCase {
	repo := repository.NewThemeRepository(s.GetDB())
	useCase := usecase.NewThemeUseCase(repo)
	return useCase
}

func GetHandler(s *server.Server) usecase.ThemeUseCase {
	useCase := getUseCase(s)
	return useCase
}
