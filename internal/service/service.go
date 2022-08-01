package service

import (
	"app-hyperledger/internal/bcapi"
	"app-hyperledger/pkg/models"

	"github.com/goccy/go-json"
	"github.com/pkg/errors"
)

type App interface {
	CreateServices(body *models.ReqBody) (string, error)
}

type Service struct {
	App
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) СreateService(body *models.ReqBody) error {

	// Получение инстанса коннекта к блокчейну
	sys, err := bcapi.CreateInstance(&body.ConfigBC)
	if err != nil {
		return errors.Wrap(err, "cannot create system user instance")
	}
	defer sys.Conn.Close()

	servicesByte, err := json.Marshal(body.Services)
	if err != nil {
		return errors.Wrap(err, "marshal error")
	}

	_, err = sys.Contract.SubmitTransaction("СreateService", string(servicesByte))
	if err != nil {
		return bcapi.ErrorHandling(err)
	}

	return nil
}

func (s *Service) GetServiceByUUID(req *models.ReqBody) (*models.Services, error) {

	// Получение инстанса коннекта к блокчейну
	sys, err := bcapi.CreateInstance(&req.ConfigBC)
	if err != nil {
		return nil, errors.Wrap(err, "cannot get service user instance")
	}
	defer sys.Conn.Close()

	resByte, err := sys.Contract.SubmitTransaction("GetServiceByUUID", req.InputType, req.InputPhone)
	if err != nil {
		return nil, bcapi.ErrorHandling(err)
	}

	service := &models.Services{}
	err = json.Unmarshal(resByte, service)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal error")
	}

	return service, nil
}

func (s *Service) SetServiceStatus(req *models.ReqBody) error {
	// Получение инстанса коннекта к блокчейну
	sys, err := bcapi.CreateInstance(&req.ConfigBC)
	if err != nil {
		return errors.Wrap(err, "cannot get service user instance")
	}
	defer sys.Conn.Close()

	_, err = sys.Contract.SubmitTransaction("SetServiceStatus", req.InputType, req.InputPhone, req.InputStatus)
	if err != nil {
		return bcapi.ErrorHandling(err)
	}

	return nil
}

func (s *Service) WithDrawService(req *models.ReqBody) error {
	// Получение инстанса коннекта к блокчейну
	sys, err := bcapi.CreateInstance(&req.ConfigBC)
	if err != nil {
		return errors.Wrap(err, "cannot get service user instance")
	}
	defer sys.Conn.Close()

	_, err = sys.Contract.SubmitTransaction("WithDrawService", req.InputType, req.InputPhone)
	if err != nil {
		return bcapi.ErrorHandling(err)
	}

	return nil
}

func (s *Service) DeleteService(req *models.ReqBody) error {
	// Получение инстанса коннекта к блокчейну
	sys, err := bcapi.CreateInstance(&req.ConfigBC)
	if err != nil {
		return errors.Wrap(err, "cannot get service user instance")
	}
	defer sys.Conn.Close()

	_, err = sys.Contract.SubmitTransaction("DeleteService", req.InputType, req.InputPhone)
	if err != nil {
		return bcapi.ErrorHandling(err)
	}

	return nil
}
