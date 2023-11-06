package listing

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raafly/inventory-management/helper"
)

type UserRepository interface {
	SignUp(ctx context.Context, tx *sql.Tx, user User) User
	SignIn(ctx context.Context, tx *sql.Tx, user User) (User, error)
	Update(ctx context.Context, tx *sql.Tx, user UserUpdate) UserUpdate
	Delete(ctx context.Context, tx *sql.Tx, userName string)
	FindById(ctx context.Context, tx *sql.Tx, userName string) (*User, error)
}

type UserRepositoryImpl struct{
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) SignUp(ctx context.Context, tx *sql.Tx, user User) User {
	SQL := "INSERT INTO users (id, username, email, password, cpassword) VALUES ($1, $2, $3, $4, $5) "
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.Email, user.Password, user.Cpassword) 
	helper.PanicIfError(err)
	return user
} 

func (r *UserRepositoryImpl) SignIn(ctx context.Context, tx *sql.Tx, user User) (User, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE email = $1"
	rows, err := tx.QueryContext(ctx, SQL, user.Email)
	helper.PanicIfError(err)
	defer rows.Close()

	user = User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("account not found")
	}
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user UserUpdate) UserUpdate {
	SQL := "UPDATE FROM users SET password = $1 WHERE email = $2"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Email)
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userName string) {
	SQL := "DELETE FROM users WHERE username = $1"
	tx.ExecContext(ctx, SQL, userName)
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userName string) (*User, error) {
	var user User
	SQL := "SELECT usermame, email, password FROM users WHERE username = $1"
	rows, err := tx.QueryContext(ctx, SQL, userName)
	if err != nil {
		return nil, errors.New("account not found")
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Email, &user.Password)
		helper.PanicIfError(err)
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
	SQL := "INSERT INTO items(id, name, category,quantity) VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, item.Id, item.Name, item.Category, item.Quantity)
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
		err := rows.Scan(&item.Id, &item.Name, &item.Description, &item.Category, &item.Quantity, &item.In, &item.Created_at)
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
		err := rows.Scan(&items.Id, &items.Name, &items.Description, &items.Category, &items.Quantity, &items.In, &items.Created_at)
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
