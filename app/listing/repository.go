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
	Update(ctx context.Context, tx *sql.Tx, user User)
	Delete(ctx context.Context, tx *sql.Tx, userName string) error
	FindById(ctx context.Context, tx *sql.Tx, userName string) (User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []User
}

type UserRepositoryImpl struct{
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) SignUp(ctx context.Context, tx *sql.Tx, user User) User {
	SQL := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) "
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.Email, user.Password) 
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
		return user, errors.New("username and password don't match")
	}
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user User) {
	SQL := "UPDATE FROM users SET password = $1 WHERE email = $2"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Email)
	helper.PanicIfError(err)
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userName string) error {
	SQL := "DELETE FROM users WHERE username = $1"
	_, err := tx.ExecContext(ctx, SQL, userName)

	if err != nil {
		return errors.New("username not found ")
	} else {
		return nil
	}
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userName string) (User, error) {
	SQL := "SELECT username, email FROM users WHERE username = $1"
	rows, err := tx.QueryContext(ctx, SQL, userName)
	helper.PanicIfError(err)
	defer rows.Close()

	user := User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Email)
		helper.PanicIfError(err)	
		return user, nil
	} else {
		return user, errors.New("user not match")
	}
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []User {
	SQL := "SELECT id, username, email FROM users"
	rows, _ := tx.QueryContext(ctx, SQL)
	defer rows.Close()

	var user []User
	for rows.Next() {
		users := User{}
		err := rows.Scan(&users.Id, &users.Username, &users.Email)
		helper.PanicIfError(err)
		user = append(user, users)
	}

	return user
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