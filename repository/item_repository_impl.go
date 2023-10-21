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
	SQL := "INSERT INTO items(id, name, quantity) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, item.Id, item.Name, item.Quantity)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	item.Id = int(id)
	return item
}

func (r *ItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, item entity.Item) entity.Item {
	SQL := "UPDATE FROM items SET quantity = ?, out = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, SQL, item.Quantity, item.Out, item.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	item.Id = int(id)
	return item
}

func (r *ItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, itemId int) {
	SQL := "DELETE FROM items WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, itemId)
	helper.PanicIfError(err)
}

func (r *ItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, itemId int) (entity.Item, error) {
	SQL := "SELECT id, name, description, quantity, in, out, created_at FROM item WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, itemId)
	helper.PanicIfError(err)
	defer rows.Close()

	item := entity.Item{}
	if rows.Next() {
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Quantity, &item.In, &item.Out, &item.Created_at)
		helper.PanicIfError(err)
		return item, nil
	} else {
		return item, nil
	}
}

func (r *ItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Item {
	SQL := "SELECT id, name, description, quantity, in, out, created_at FROM item"
	rows, _ := tx.QueryContext(ctx, SQL)
	defer rows.Close()

	var item []entity.Item
	for rows.Next() {
		items := entity.Item{}
		err := rows.Scan(&items.Id, &items.Name, &items.Description, &items.In, &items.Out, &items.Created_at)
		helper.PanicIfError(err)
		item = append(item, items)
	}

	return item
}