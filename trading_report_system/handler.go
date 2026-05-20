package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetTradingReport(c *gin.Context) {
	var filters ReportFilters

	// Basic query param parsing to filters
	filters.ClientIDs = c.QueryArray("client_ids")
	filters.StrategyNames = c.QueryArray("strategy_names")
	filters.Symbols = c.QueryArray("symbols")
	filters.OrderType = c.Query("order_type")
	filters.DateFrom = c.Query("date_from")
	filters.DateTo = c.Query("date_to")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	filters.Page = page
	filters.Limit = limit

	if c.Query("include_positions") == "true" {
		filters.IncludePositions = true
	}
	if c.Query("include_trades") == "true" {
		filters.IncludeTrades = true
	}

	report, err := h.service.GetReport(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}

func (h *Handler) SetupRoutes(router *gin.Engine) {
	router.GET("/api/reports/trading-summary", h.GetTradingReport)
}
