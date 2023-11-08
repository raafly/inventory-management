package listing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/raafly/inventory-management/pkg/helper"
)

type UserRepository interface {
	SignUp(user User) error
	SignIn(user User, db*sql.DB) (*User, error)
}

type UserRepositoryImpl struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) SignUp(user User) error {
	SQL := "INSERT INTO users(id, username, email, password) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(SQL, user.Id, user.Username, user.Email, user.Password) 
	if err != nil {
		return fmt.Errorf("FAILED EXEC QUERY %v", err.Error())
	}

	return nil
}

func (r *UserRepositoryImpl) SignIn(user User, db*sql.DB) (*User, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE email = $1"

	if err := r.db.QueryRow(SQL, user.Email).Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
		fmt.Errorf("Failed to exec query %v", err.Error())
		return &user, nil
	}

	return &user, nil
}


// item

type ItemRepository interface {
	Create(ctx context.Context, tx *sql.Tx, item Item) Item
	Update(ctx context.Context, tx *sql.Tx, item Item) Item
	Delete(ctx context.Context, tx *sql.Tx, itemName string)
	FindById(ctx context.Context, tx *sql.Tx, itemName string) (Item, error)
	FindAll(ctx context.Context, tx *sql.Tx) []Item
}

type ItemRepositoryImpl struct {
}

func NewItemRepository() ItemRepository {
	return &ItemRepositoryImpl{}
}

func (r *ItemRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, item Item) Item {
	SQL := "INSERT INTO items(name, category,quantity) VALUES($1, $2, $3)"
	_, err := tx.ExecContext(ctx, SQL, item.Name, item.Category, item.Quantity)
	helper.PanicIfError(err)

	return item
}

func (r *ItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, item Item) Item {
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

func (r *ItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, itemName string) (Item, error) {
	SQL := "SELECT id, name, description, category, quantity, int, created_at FROM items WHERE name = $1"
	rows, err := tx.QueryContext(ctx, SQL, itemName)
	helper.PanicIfError(err)
	defer rows.Close()

	item := Item{}
	if rows.Next() {
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Category, &item.Created_at)
		helper.PanicIfError(err)
		return item, nil
	} else {
		return item, nil
	}
}

func (r *ItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []Item {
	SQL := "SELECT id, name, description, category, quantity, int, created_at FROM items"
	rows, _ := tx.QueryContext(ctx, SQL)
	defer rows.Close()

	var item []Item
	for rows.Next() {
		items := Item{}
		err := rows.Scan(&items.Id, &items.Name, &items.Description, &items.Category, &items.Quantity, &items.Created_at)
		helper.PanicIfError(err)
		item = append(item, items)
	}

	return item
}

// category

type CategoryRepository	interface {
	Create(ctx context.Context, tx *sql.Tx, data Category) Category
	Update(ctx context.Context, tx *sql.Tx, data *Category) (*Category, error)
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repo *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, data Category) Category {
	SQL := "INSERT INTO categories(id, name, description) VALUES($1, $2, $3)"
	_, err := tx.ExecContext(ctx, SQL, data.Id, data.Name, data.Description) 
	helper.PanicIfError(err)

	return data
}

func (repo *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, data *Category) (*Category, error) {
	SQL := "UPDATE categories SET description = $1 WHERE id = $2"
	_, err := tx.ExecContext(ctx, SQL, data.Description, data.Id)
	if err != nil {
		return nil, errors.New("id not found")
	}
	
	return data, nil
}
