package port

import (
	"context"
	"database/sql"

	"github.com/raafly/inventory-management/entity"
)

type ItemRepository interface {
	Create(ctx context.Context, tx *sql.Tx, item entity.Item) entity.Item
	Update(ctx context.Context, tx *sql.Tx, item entity.Item) entity.Item
	Delete(ctx context.Context, tx *sql.Tx, itemName string)
	FindById(ctx context.Context, tx *sql.Tx, itemName string) (entity.Item, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Item
}