package products

import "context"

type Service interface {
	GetAllProducts(ctx context.Context) error
}

type Svc struct {
}

func NewService() Service {
	return &Svc{}
}

func (s *Svc) GetAllProducts(ctx context.Context) error {
	return nil
}
