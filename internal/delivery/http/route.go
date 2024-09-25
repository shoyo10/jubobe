package http

import (
	"github.com/labstack/echo/v4"
)

// SetRoutes set routes
func SetRoutes(e *echo.Echo, h Handler) {
	e.GET("/patients", h.ListPatients)
	// e.POST("/orders", h.CreateOrder)
	// e.PUT("/orders/:id", h.UpdateOrder)
}
