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
			wh.GET("/available/products/:id", h.getAvailableProducts)
			wh.PUT("/reserve/products", h.reserveProducts)
			wh.PUT("/release/products", h.releaseProducts)
		}
	}
	return router
}
