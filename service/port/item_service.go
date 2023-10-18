package port

import (
	"context"

	"github.com/raafly/inventory-management/model"
)

type ItemPort interface {
	Create(ctx context.Context, request model.ItemCreate) model.ItemResponse
	Update(ctx context.Context, requst model.ItemUpdate) model.ItemResponse
	Delete(ctx context.Context, itemId int)
	FindById(ctx context.Context, itemId int ) model.ItemResponse
	FindAll(ctx context.Context) []model.ItemResponse
}