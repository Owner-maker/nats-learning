package http

import (
	"github.com/gin-gonic/gin"
	"nats-learning/internal/models"
	"nats-learning/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

type getAllOrdersResponse struct {
	Data []models.Order `json:"data"`
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/order:uid", h.GetOrderById)
	router.GET("/orders", h.GetAllOrders)

	return router
}

func (h *Handler) GetOrderById(c *gin.Context) {
	uid := c.Param("uid")
	order, err := h.service.GetCachedOrder(uid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllCachedOrders()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllOrdersResponse{
		Data: orders,
	})
}
