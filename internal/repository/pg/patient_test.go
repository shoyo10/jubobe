package pg

import (
	"context"
	"jubobe/internal/model"
	"jubobe/internal/repository"
	"jubobe/pkg/config"
	"jubobe/pkg/postgres"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

func TestPatientDB(t *testing.T) {
	suite.Run(t, new(patientDBSuite))
}

type patientDBSuite struct {
	suite.Suite
	conn *gorm.DB
	repo repository.Repositorier
	ctx  context.Context
}

func (s *patientDBSuite) SetupSuite() {
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
}

func (s *patientDBSuite) SetupTest() {
	if err := s.conn.Delete(&model.Order{}, "1 = 1").Error; err != nil {
		log.Println("SetupTest delete order table failed", err.Error())
	}
	if err := s.conn.Delete(&model.Patient{}, "1 = 1").Error; err != nil {
		log.Println("SetupTest delete patient table failed", err.Error())
	}
}

func (s *patientDBSuite) TearDownSuite() {
	if err := s.conn.Delete(&model.Order{}, "1 = 1").Error; err != nil {
		log.Println("SetupTest delete order table failed", err.Error())
	}
	if err := s.conn.Delete(&model.Patient{}, "1 = 1").Error; err != nil {
		log.Println("TearDownSuite delete patient table failed", err.Error())
	}
}

func (s *patientDBSuite) TestListPatients() {
	// test empty list
	patients, err := s.repo.ListPatients(s.ctx, nil)
	s.Require().NoError(err)
	s.Len(patients, 0)

	// test list with one patient
	p := model.Patient{
		Name: "John Doe",
	}
	err = s.conn.Create(&p).Error
	s.Require().NoError(err)

	patients, err = s.repo.ListPatients(s.ctx, nil)
	s.Require().NoError(err)
	s.Len(patients, 1)
	s.Equal("John Doe", patients[0].Name)

	// test list with one patient and preload order
	patients, err = s.repo.ListPatients(s.ctx, &model.PatientOption{IsPreloadOrder: true})
	s.Require().NoError(err)
	s.Len(patients, 1)
	s.Equal("John Doe", patients[0].Name)
	s.Zero(patients[0].Order.ID)

	// test list with one patient and order
	o := model.Order{
		PatientID: p.ID,
		Message:   "msg 123",
	}
	err = s.conn.Create(&o).Error
	s.Require().NoError(err)

	patients, err = s.repo.ListPatients(s.ctx, &model.PatientOption{IsPreloadOrder: true})
	s.Require().NoError(err)
	s.Len(patients, 1)
	s.Equal("John Doe", patients[0].Name)
	s.Equal("msg 123", patients[0].Order.Message)

	// test list with two patient
	p = model.Patient{
		Name: "Jay",
	}
	err = s.conn.Create(&p).Error
	s.Require().NoError(err)

	patients, err = s.repo.ListPatients(s.ctx, nil)
	s.Require().NoError(err)
	s.Len(patients, 2)
	s.Equal("John Doe", patients[0].Name)
	s.Equal("Jay", patients[1].Name)
}
