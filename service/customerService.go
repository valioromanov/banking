package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	GetCustomerByStatus(int) ([]domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {

	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func (s DefaultCustomerService) GetCustomerByStatus(status int) ([]domain.Customer, *errs.AppError) {

	return s.repo.ByStatus(status)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {

	return DefaultCustomerService{repository}
}
