package http

import (
	"jubobe/internal/model"
	"jubobe/internal/service"
	"jubobe/pkg/errors"
	"net/http"
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
	Data []listPatientsRespData `json:"Data"`
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
	Data createOrderRespData `json:"Data"`
}

type createOrderRespData struct {
	ID int `json:"Id"`
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

type updateOrderReq struct {
	Message string `json:"Message" validate:"gte=1,lte=255"`
}

type updateOrderParam struct {
	ID int `param:"id" validate:"gte=1"`
}

// @Title  UpdateOrder
// @Description update a order
// @Param id path int true "order id"
// @Param reqBody body updateOrderReq true "update order fields"
// @Success 200
// @Failure 400 object errors.HTTPError
// @Failure 404 object errors.HTTPError
// @Router /orders/{id} [put]
func (h *handler) UpdateOrder(c echo.Context) error {
	ctx := c.Request().Context()

	var req updateOrderReq
	err := c.Bind(&req)
	if err != nil {
		return errors.Wrap(errors.ErrInvalidInput, err.Error())
	}

	req.Message = strings.TrimSpace(req.Message)
	if err := c.Validate(&req); err != nil {
		return errors.Wrap(errors.ErrInvalidInput, err.Error())
	}

	var pathParam updateOrderParam
	err = (&echo.DefaultBinder{}).BindPathParams(c, &pathParam)
	if err != nil {
		return errors.Wrap(errors.ErrInvalidInput, err.Error())
	}
	if err := c.Validate(&pathParam); err != nil {
		return errors.Wrap(errors.ErrInvalidInput, err.Error())
	}

	err = h.svc.UpdateOrder(ctx, &model.OrderOption{
		Filter: model.OrderFilter{
			ID: pathParam.ID,
		}},
		model.UpdateOrderInput{
			Message: req.Message,
		},
	)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
