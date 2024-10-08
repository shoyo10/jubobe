package http

import (
	"github.com/labstack/echo/v4"
)

// @title Jubobe API Document
// @version 1.0
// @description This is jubo backend api document.

// @contact.name Shoyo
// @contact.url https://github.com/shoyo10/jubobe

// @host localhost:9090
// @BasePath /

// SetRoutes set routes
func SetRoutes(e *echo.Echo, h Handler) {
	e.GET("/api/patients", h.ListPatients)
	e.POST("/api/orders", h.CreateOrder)
	e.PUT("/api/orders/:id", h.UpdateOrder)
	e.GET("/api/orders/:id", h.GetOrder)
}
