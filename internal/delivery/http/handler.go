package http

import (
	_ "github.com/Owner-maker/nats-learning/docs"
	"github.com/Owner-maker/nats-learning/internal/models"
	"github.com/Owner-maker/nats-learning/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service.Order
}

type getAllOrdersResponse struct {
	Data []models.Order `json:"data"`
}

func NewHandler(s service.Order) *Handler {
	return &Handler{s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/order/:uid", h.GetOrderById)
		api.GET("/order/db/:uid", h.GetDbOrderById)
		api.GET("/orders", h.GetAllOrders)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
