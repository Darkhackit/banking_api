package service

import (
	"github/Darkhackit/banking_api/domain"
	"github/Darkhackit/banking_api/dto"
	"github/Darkhackit/banking_api/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, error)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppErrors)
}
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, len(customers))
	for i, c := range customers {
		response[i] = dto.CustomerResponse{
			Id:          c.Id,
			Name:        c.Name,
			Status:      c.StatusAsText(),
			DateOfBirth: c.DateOfBirth,
			City:        c.City,
			ZipCode:     c.ZipCode,
		}
	}
	return response, nil
}
func (s DefaultCustomerService) GetCustomerById(customerId string) (*dto.CustomerResponse, *errs.AppErrors) {
	customer, err := s.repo.ById(customerId)
	if err != nil {
		return nil, err
	}
	response := customer.ToDto()
	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
