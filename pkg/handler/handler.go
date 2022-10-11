package handler

import (
	"github.com/asadbek21coder/bookshelf/pkg/service"
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

	router.POST("/signup", h.signUp)
	router.GET("/myself", h.getUser)

	books := router.Group("books")
	{
		books.GET("/", h.getAllBooks)
		books.POST("/", h.createBook)
		books.PATCH("/:id", h.updateBook)
		books.DELETE("/:id", h.deleteBook)
	}

	return router
}
