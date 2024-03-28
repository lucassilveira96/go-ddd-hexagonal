package service

import (
	"context"
	"go-ddd-hexagonal/domain/model"
	"go-ddd-hexagonal/domain/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *model.Product) error {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *model.Product) error {
	return s.repo.Update(ctx, product)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) FindAllProduct(ctx context.Context) ([]*model.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) FindByIdProduct(ctx context.Context, id int64) (*model.Product, error) {
	return s.repo.FindByID(ctx, id)
}
