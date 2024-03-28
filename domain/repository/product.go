package repository

import (
	"context"
	"go-ddd-hexagonal/domain/model"
)

type ProductRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Product, error)
	FindAll(ctx context.Context) ([]*model.Product, error)
	Create(ctx context.Context, product *model.Product) error
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, product *model.Product) error
}
