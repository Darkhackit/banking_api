package service

import (
	"github/Darkhackit/banking_api/domain"
	"github/Darkhackit/banking_api/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomerById(string) (*domain.Customer, *errs.AppErrors)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}
func (s DefaultCustomerService) GetCustomerById(customerId string) (*domain.Customer, *errs.AppErrors) {
	return s.repo.ById(customerId)
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
