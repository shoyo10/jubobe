package http

import (
	"jubobe/internal/model"
	"jubobe/internal/service"

	"github.com/labstack/echo/v4"
)

type handler struct {
	svc service.Servicer
}

// NewHandler create Handler instance
func NewHandler(svc service.Servicer) Handler {
	return &handler{
		svc: svc,
	}
}

type listPatientsResp struct {
	Data []listPatientsRespData `json:"data"`
}

type listPatientsRespData struct {
	ID      int    `json:"Id"`
	Name    string `json:"Name"`
	OrderID int    `json:"OrderId"`
}

func (h *handler) ListPatients(c echo.Context) error {
	ctx := c.Request().Context()
	opt := &model.PatientOption{
		IsPreloadOrder: true,
	}
	patients, err := h.svc.ListPatients(ctx, opt)
	if err != nil {
		return err
	}
	resp := listPatientsResp{
		Data: make([]listPatientsRespData, len(patients)),
	}
	for i := 0; i < len(patients); i++ {
		p := patients[i]
		resp.Data[i] = listPatientsRespData{
			ID:      p.ID,
			Name:    p.Name,
			OrderID: p.Order.ID,
		}
	}
	return c.JSON(200, resp)
}
