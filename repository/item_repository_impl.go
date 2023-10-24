package repository

import (
	"context"
	"database/sql"

	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/repository/port"
)

type ItemRepositoryImpl struct {
}

func NewItemRepository() port.ItemRepository {
	return &ItemRepositoryImpl{}
}

func (r *ItemRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, item entity.Item) entity.Item {
	SQL := "INSERT INTO items(id, name, category,quantity) VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, item.Id, item.Name, item.Category, item.Quantity)
	helper.PanicIfError(err)

	return item
}

func (r *ItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, item entity.Item) entity.Item {
	SQL := "UPDATE items SET quantity = $1 WHERE name = $2"
	_, err := tx.ExecContext(ctx, SQL, item.Quantity, item.Name)
	helper.PanicIfError(err)

	return item
}

func (r *ItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, itemName string) {
	SQL := "DELETE FROM items WHERE name = $1"
	_, err := tx.ExecContext(ctx, SQL, itemName)
	helper.PanicIfError(err)
}

func (r *ItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, itemName string) (entity.Item, error) {
	SQL := "SELECT id, name, description, category, quantity, int, created_at FROM items WHERE name = $1"
	rows, err := tx.QueryContext(ctx, SQL, itemName)
	helper.PanicIfError(err)
	defer rows.Close()

	item := entity.Item{}
	if rows.Next() {
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Category, &item.Quantity, &item.In, &item.Created_at)
		helper.PanicIfError(err)
		return item, nil
	} else {
		return item, nil
	}
}

func (r *ItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Item {
	SQL := "SELECT id, name, description, category, quantity, int, created_at FROM items"
	rows, _ := tx.QueryContext(ctx, SQL)
	defer rows.Close()

	var item []entity.Item
	for rows.Next() {
		items := entity.Item{}
		err := rows.Scan(&items.Id, &items.Name, &items.Description, &items.Category, &items.Quantity, &items.In, &items.Created_at)
		helper.PanicIfError(err)
		item = append(item, items)
	}

	return item
}