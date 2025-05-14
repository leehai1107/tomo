package http

import (
	"github.com/leehai1107/tomo/service/tomo/usecase"
)

// IHandler defines all handler interfaces
type IHandler interface {
	IUserHandler
	IAdminHandler
	ICoffeeShopHandler
	IChatHandler
}

// Handler implements all handler interfaces
type Handler struct {
	userUsecase  usecase.IUserUsecase
	adminUsecase usecase.IAdminUsecase
	// Add other usecases as needed
}

func NewHandler(userUsecase usecase.IUserUsecase, adminUsecase usecase.IAdminUsecase) IHandler {
	return &Handler{
		userUsecase:  userUsecase,
		adminUsecase: adminUsecase,
	}
}
