package test

import (
	"fmt"
	"testing"

	"github.com/raafly/inventory-management/pkg/config"
)

type User struct {
	Id string
	Username string
	Email	string
	Password string
}

var db = config.NewDB()

func SignUp(user User) error {
	SQL := "INSERT INTO users(id, username, email, password) VALUES ($1, $2, $3, $4)"
	if _, err := db.Exec(SQL, user.Id, user.Username, user.Email, user.Password); err != nil {
		return fmt.Errorf("failed: %v", err)
	}

	return nil
}

func TestInsertDB(t *testing.T) {
	data := User {
		Id: "nerwnewcxnxwf",
		Username: "newcyfegyu",
		Email: "uncfrexnucrqem",
		Password: "huxheumheux",
	}

	SignUp(data)
}