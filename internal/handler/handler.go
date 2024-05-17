package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kisnikita/wh-api/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		wh := api.Group("/warehouse")
		{
			wh.GET(":id", h.getByID)
			wh.POST("/reserve", h.reserveProducts)
			wh.POST("/release", h.releaseProducts)
		}
	}
	return router
}
