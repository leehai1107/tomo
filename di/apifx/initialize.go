package apifx

import (
	"github.com/leehai1107/tomo/service/tomo/delivery/http"
	"github.com/leehai1107/tomo/service/tomo/repository"
	"github.com/leehai1107/tomo/service/tomo/usecase"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Provide(
	provideRouter,
	provideHandler,
	provideRepo,
	provideUserUsecase,
	provideAdminUsecase,
)

func provideRouter(handler http.IHandler) http.Router {
	return http.NewRouter(handler)
}

func provideHandler(userUsecase usecase.IUserUsecase, adminUsecase usecase.IAdminUsecase) http.IHandler {
	handler := http.NewHandler(userUsecase, adminUsecase)
	return handler
}

func provideRepo(db *gorm.DB) repository.IUserRepo {
	return repository.NewUserRepo(db)
}

func provideUserUsecase(repo repository.IUserRepo) usecase.IUserUsecase {
	return usecase.NewUserUsecase(repo)
}

func provideAdminUsecase(repo repository.IUserRepo) usecase.IAdminUsecase {
	return usecase.NewAdminUsecase(repo)
}
