package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.GET(":uid", h.GetOrderById)
			orders.GET("", h.GetAllOrders)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

// GetOrderById
// @Summary GetOrderById
// @Description Allows to get specific order from the app's cache via its uid
// @ID ger-order-by-id
// @Accept json
// @Produce json
// @Param uid path string true "order's uid"
// @Success 200 {object} models.Order
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/order/:uid [get]
func (h *Handler) GetOrderById(c *gin.Context) {
	uid := c.Param("uid")
	order, err := h.service.GetCachedOrder(uid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, order)
}

// GetAllOrders
// @Summary GetAllOrders
// @Description Allows to get all orders from the app's cache
// @ID ger-all-orders
// @Accept json
// @Produce json
// @Success 200 {object} getAllOrdersResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/orders [get]
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
