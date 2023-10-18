package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/helper"
)

type UserRepositoryImpl struct{
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) SignUp(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := "INSERT INTO users(id, username, email, password) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.Email, user.Password) 
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user
} 

func (r *UserRepositoryImpl) SignIn(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	SQL := "SELECT email, password FROM users WHERE email = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Email, user.Password)
	helper.PanicIfError(err)
	defer rows.Close()

	user = entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("username and password don't match")
	}
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) {
	SQL := "UPDATE FROM users SET password = ? WHERE email = ?"
	_, err := tx.ExecContext(ctx, SQL, user.Password, user.Email)
	helper.PanicIfError(err)
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfError(err)
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error) {
	SQL := "SELECT username, email FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Email)
		helper.PanicIfError(err)	
		return user, nil
	} else {
		return user, errors.New("user not match")
	}
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.User {
	SQL := "SELECT id, username, email FROM users"
	rows, _ := tx.QueryContext(ctx, SQL)
	defer rows.Close()

	var user []entity.User
	for rows.Next() {
		users := entity.User{}
		err := rows.Scan(&users.Id, &users.Username, &users.Email)
		helper.PanicIfError(err)
		user = append(user, users)
	}

	return user
}