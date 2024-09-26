package pg

import (
	"context"

	"jubobe/internal/model"
	"jubobe/internal/repository"
	"jubobe/pkg/config"
	"jubobe/pkg/errors"
	"jubobe/pkg/postgres"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

func TestOrderDB(t *testing.T) {
	suite.Run(t, new(orderDBSuite))
}

type orderDBSuite struct {
	suite.Suite
	conn *gorm.DB
	repo repository.Repositorier
	ctx  context.Context
}

func (s *orderDBSuite) SetupSuite() {
	dir, err := os.Getwd()
	log.Println("dir:", dir)
	s.Require().NoError(err)

	os.Setenv("CONFIG_DIR", dir+"/../../../configs")
	os.Setenv("CONFIG_NAME", "app-test")

	cfg, err := config.New()
	s.Require().NoError(err)

	conn, err := postgres.New(cfg.Postgres)
	s.Require().NoError(err)

	err = conn.AutoMigrate(&model.Patient{}, &model.Order{})
	s.Require().NoError(err)

	s.conn = conn
	repo, err := New(conn)
	s.Require().NoError(err)
	s.repo = repo
	s.ctx = context.Background()

	if err := s.conn.Delete(&model.Patient{}, "1 = 1").Error; err != nil {
		log.Println("SetupTest delete patient table failed", err.Error())
	}

	err = s.conn.Create(&model.Patient{Name: "John"}).Error
	s.Require().NoError(err)
}

func (s *orderDBSuite) SetupTest() {
	if err := s.conn.Delete(&model.Order{}, "1 = 1").Error; err != nil {
		log.Println("SetupTest delete order table failed", err.Error())
	}
}

func (s *orderDBSuite) TearDownSuite() {
	if err := s.conn.Delete(&model.Order{}, "1 = 1").Error; err != nil {
		log.Println("SetupTest delete order table failed", err.Error())
	}
	if err := s.conn.Delete(&model.Patient{}, "1 = 1").Error; err != nil {
		log.Println("TearDownSuite delete patient table failed", err.Error())
	}
}

func (s *orderDBSuite) TestCreateOrder() {
	var p model.Patient
	err := s.conn.First(&p).Error
	s.Require().NoError(err)

	order := &model.Order{
		PatientID: p.ID,
		Message:   "test 123456",
	}
	err = s.repo.CreateOrder(s.ctx, order)
	s.Require().NoError(err)
	s.NotZero(order.ID)

	var order2 model.Order
	err = s.conn.First(&order2).Error
	s.Require().NoError(err)
	s.Equal(order.ID, order2.ID)
	s.Equal(order.PatientID, order2.PatientID)
	s.Equal(order.Message, order2.Message)

	// test duplicate patient id
	err = s.repo.CreateOrder(s.ctx, order)
	s.Require().True(errors.Is(err, errors.ErrResourceAlreadyExists))
}

func (s *orderDBSuite) TestUpdateOrder() {
	var p model.Patient
	err := s.conn.First(&p).Error
	s.Require().NoError(err)

	order := &model.Order{
		PatientID: p.ID,
		Message:   "test 123456",
	}
	err = s.repo.CreateOrder(s.ctx, order)
	s.Require().NoError(err)

	newMessage := "test 654321"
	orderOpt := &model.OrderOption{
		Filter: model.OrderFilter{ID: order.ID},
	}
	err = s.repo.UpdateOrder(s.ctx, orderOpt, model.UpdateOrderInput{Message: newMessage})
	s.Require().NoError(err)

	var order2 model.Order
	err = s.conn.First(&order2).Error
	s.Require().NoError(err)
	s.Equal(order.ID, order2.ID)
	s.Equal(order.PatientID, order2.PatientID)
	s.Equal(newMessage, order2.Message)

	// update not exist order id
	orderOpt.Filter.ID = 999
	err = s.repo.UpdateOrder(s.ctx, orderOpt, model.UpdateOrderInput{Message: newMessage})
	s.Require().True(errors.Is(err, errors.ErrResourceNotFound))
}

func (s *orderDBSuite) TestGetOrder() {
	var p model.Patient
	err := s.conn.First(&p).Error
	s.Require().NoError(err)

	order := &model.Order{
		PatientID: p.ID,
		Message:   "test 123456",
	}
	err = s.repo.CreateOrder(s.ctx, order)
	s.Require().NoError(err)

	orderOpt := &model.OrderOption{
		Filter: model.OrderFilter{ID: order.ID},
	}
	order2, err := s.repo.GetOrder(s.ctx, orderOpt)
	s.Require().NoError(err)
	s.Equal(order.ID, order2.ID)
	s.Equal(order.PatientID, order2.PatientID)
	s.Equal(order.Message, order2.Message)

	// get not exist order id
	orderOpt.Filter.ID = 999
	order2, err = s.repo.GetOrder(s.ctx, orderOpt)
	s.Require().True(errors.Is(err, errors.ErrResourceNotFound))
	s.Nil(order2)
}
