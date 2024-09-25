package http

import (
	"context"
	"encoding/json"
	"io"
	"jubobe/internal/model"
	"jubobe/internal/service/mocks"
	"jubobe/pkg/echorouter"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"
)

func TestHanlder(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

type handlerSuite struct {
	suite.Suite

	mockSvc *mocks.MockServicer
	handler Handler
	app     *fx.App
	e       *echo.Echo
}

func (s *handlerSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())
	s.mockSvc = mocks.NewMockServicer(ctrl)
	s.handler = NewHandler(s.mockSvc)
	s.app = fx.New(
		fx.Supply(&echorouter.Config{}),
		fx.Provide(echorouter.FxNewEcho),
		fx.Populate(&s.e),
	)
	SetRoutes(s.e, s.handler)
	err := s.app.Start(context.Background())
	s.Require().NoError(err)
}

func (s *handlerSuite) TearDownSuite() {
	s.app.Stop(context.Background())
}

func (s *handlerSuite) TestListPatients() {
	patients := []model.Patient{
		{
			ID:   1,
			Name: "John",
		},
		{
			ID:   2,
			Name: "Doe",
			Order: model.Order{
				ID:        1,
				PatientID: 2,
				Message:   "Hello",
			},
		},
	}
	s.mockSvc.EXPECT().ListPatients(gomock.Any(), gomock.Any()).Return(patients, nil)
	rec := request(http.MethodGet, "/patients", nil, s.e)
	var resp listPatientsResp
	json.Unmarshal(rec.Body.Bytes(), &resp)
	s.Require().Equal(http.StatusOK, rec.Code)
	s.Require().Len(resp.Data, 2)
	s.Require().Equal(1, resp.Data[0].ID)
	s.Require().Equal("John", resp.Data[0].Name)
	s.Require().Equal(0, resp.Data[0].OrderID)
	s.Require().Equal(2, resp.Data[1].ID)
	s.Require().Equal("Doe", resp.Data[1].Name)
	s.Require().Equal(1, resp.Data[1].OrderID)
}

func request(method, path string, body io.Reader, e *echo.Echo) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec
}
