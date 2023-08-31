package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ulnesterova/avito_test/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api")
	{
		api.POST("/segments", h.createSegment)
		api.DELETE("/segments/:slug", h.deleteSegment)
		api.POST("/users/:id/segments", h.addUserSegments)
		api.DELETE("/users/:id/:slug", h.deleteUserSegments)
		api.GET("/users/:id/segments", h.getUserSegments)
	}
	return router
}
