package application

import "github.com/fabruun/go-customers/domain"

type CustomerServiceInterface interface {
	List() ([]*domain.Customer, error)
}

type CustomerService struct{}

func (cs *CustomerService) List([]*domain.Customer, error) {
}
