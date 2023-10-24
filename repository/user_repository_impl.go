package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/repository/port"
)

type UserRepositoryImpl struct{
}

func NewUserRepository() port.UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) SignUp(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) "
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Username, user.Email, user.Password) 
	helper.PanicIfError(err)
	return user
} 

func (r *UserRepositoryImpl) SignIn(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error) {
	SQL := "SELECT id, username, email, password FROM users WHERE email = $1"
	rows, err := tx.QueryContext(ctx, SQL, user.Email)
	helper.PanicIfError(err)
	defer rows.Close()

	user = entity.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("username and password don't match")
	}
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) {
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

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userName string) (entity.User, error) {
	SQL := "SELECT username, email FROM users WHERE username = $1"
	rows, err := tx.QueryContext(ctx, SQL, userName)
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