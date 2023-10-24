package port

import (
	"context"

	"github.com/raafly/inventory-management/model"
)

type ItemService interface {
	Create(ctx context.Context, request model.ItemCreate) model.ItemResponse
	Update(ctx context.Context, requst model.ItemUpdate) model.ItemResponse
	Delete(ctx context.Context, itemName string)
	FindById(ctx context.Context, itemName string ) model.ItemResponse
	FindAll(ctx context.Context) []model.ItemResponse
}