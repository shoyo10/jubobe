package http

import (
	"jubobe/internal/model"
	"jubobe/internal/service"
	"jubobe/pkg/errors"
	"strings"

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

// @Title  ListPatients
// @Description list all patients
// @Success 200 {object} listPatientsResp
// @Router /patients [get]
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

type createOrderReq struct {
	PatientID int    `json:"PatientId" validate:"gte=1"`
	Message   string `json:"Message" validate:"gte=1,lte=255"`
}

type createOrderResp struct {
	Data createOrderRespData `json:"data"`
}

type createOrderRespData struct {
	ID int `json:"id"`
}

// @Title  CreateOrder
// @Description create a order
// @Param reqBody body createOrderReq true "order fields"
// @Success 200 {object} createOrderResp "order id"
// @Failure 400 object errors.HTTPError
// @Router /orders [post]
func (h *handler) CreateOrder(c echo.Context) error {
	ctx := c.Request().Context()

	var req createOrderReq
	if err := c.Bind(&req); err != nil {
		return errors.Wrap(errors.ErrInvalidInput, err.Error())
	}

	req.Message = strings.TrimSpace(req.Message)
	if err := c.Validate(&req); err != nil {
		return errors.Wrap(errors.ErrInvalidInput, err.Error())
	}

	newOrder := &model.Order{
		PatientID: req.PatientID,
		Message:   req.Message,
	}
	err := h.svc.CreateOrder(ctx, newOrder)
	if err != nil {
		return err
	}
	return c.JSON(200, createOrderResp{Data: createOrderRespData{ID: newOrder.ID}})
}
