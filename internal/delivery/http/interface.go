package http

import "github.com/labstack/echo/v4"

type Handler interface {
	// ListPatients list patients
	ListPatients(c echo.Context) error
	CreateOrder(c echo.Context) error
	UpdateOrder(c echo.Context) error
	GetOrder(c echo.Context) error
}
