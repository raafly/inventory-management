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
	SignIn(user User) (*User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r UserRepositoryImpl) SignUp(user User) error {
	SQL := "INSERT INTO users(id, username, email, password) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(SQL, user.Id, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("FAILED EXEC QUERY %v", err.Error())
	}

	return nil
}

func (r UserRepositoryImpl) SignIn(user User) (*User, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE email = $1"
	if err := r.db.QueryRow(SQL, user.Email).Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
		return &user, fmt.Errorf("FAILED to exec query %v", err.Error())
	}

	return &user, nil
}

// item

type ItemRepository interface {
	Create(item Item) error
	UpdateStatus(id int, status bool)
	UpdateQuantity(id, quatity int)
	Delete(itemId int)
	FindById(itemId int) (*Item, error)
	// FindAll() []Item
	// UpadteDescription(id int, desc string)
}

type ItemRepositoryImpl struct {
	db *sql.DB
}

func NewItemRepository(DB *sql.DB) ItemRepository {
	return &ItemRepositoryImpl{
		db: DB,
	}
}

func (r ItemRepositoryImpl) Create(item Item) error {
	SQL := "INSERT INTO items(name, category, quantity) VALUES($1, $2, $3)"
	if _, err := r.db.Exec(SQL, item.Name, item.Category, item.Quantity); err != nil {
		return fmt.Errorf("failed exec query: %v", err)
	}
	return nil
}

func (r ItemRepositoryImpl) UpdateStatus(id int, status bool) {
	SQL := "UPDATE items SET status = $1 WHERE id = $2"
	if _, err := r.db.Exec(SQL, status, id); err != nil {
		fmt.Errorf("FAILED to exec query %v", err.Error())
	}
}

func (r ItemRepositoryImpl) UpdateQuantity(id, quatity int) {
	SQL := "UPDATE items SET quantity = quantity - $1 WHERE name = $2"
	if _, err := r.db.Exec(SQL, quatity, id); err != nil {
		fmt.Printf("FAILED to exec query %v", err.Error())
	}
}

func (r ItemRepositoryImpl) UpadteDescription(id int, desc string) {
	SQL := "UPDATE items SET description = $1 WHERE name = $2"
	if _, err := r.db.Exec(SQL, desc, id); err != nil {
		fmt.Errorf("FAILED to exec query %v", err.Error())
	}
}

func (r ItemRepositoryImpl) Delete(itemId int) {
	SQL := "DELETE FROM items WHERE id = $1"
	if _, err := r.db.Exec(SQL, itemId); err != nil {
		fmt.Errorf("FAILED to exec query %v", err.Error())
	}
}

func (r ItemRepositoryImpl) FindById(itemId int) (*Item, error) {
	SQL := "SELECT id, name, description, category, quantity, status, created_at FROM items WHERE id = $1"
	rows, err := r.db.Query(SQL, itemId)
	if err != nil {
		return nil, errors.New("item id not found")
	}
	defer rows.Close()

	var items Item
	if rows.Next() {
		err := rows.Scan(&items.Id, &items.Name, &items.Description, &items.Category, &items.Quantity, &items.Status, &items.Created_at)
		if err != nil {
			return nil, fmt.Errorf("failed exec query: %v", err)
		}
		return &items, nil
	} else {
		return nil, errors.New("item nof found")
	}
}

func (r ItemRepositoryImpl) FindAll() []Item {
	SQL := "SELECT id, name, description, category, quantity, status, created_at FROM items"
	rows, _ := r.db.Query(SQL)

	defer rows.Close()

	var item []Item
	for rows.Next() {
		items := Item{}
		if err := rows.Scan(&items.Id, &items.Name, &items.Description, &items.Category, &items.Quantity, &items.Status, &items.Created_at); err != nil {
			fmt.Printf("failed pharsing %v", err)
		}
		item = append(item, items)
	}
	return item
}

// category

type CategoryRepository interface {
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
