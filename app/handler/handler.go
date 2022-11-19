package handler

import (
	"avito-test-task/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserBalance)
			users.PUT("/:id", h.addMoneyToUserBalance)
		}
		reserve := api.Group("/reserve")
		{
			reserve.PUT("/:id", h.reserveMoneyFromUserBalance)
			reserve.DELETE("/:id", h.writeOffRevenue)
		}

	}

	return router
}
