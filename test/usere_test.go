package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/oklog/ulid/v2"
	"github.com/raafly/inventory-management/config"
	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/repository"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

func GenerateId() string {
	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	helper.PanicIfError(err)

	uniqueId := id.String()
	return uniqueId
}

func TestSignUp(t *testing.T) {
	id := GenerateId()

	db := config.NewDB()
	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	ctx := context.Background()
	data := entity.User {
		Id: id,
		Username: "rafly",
		Email: "rafliexecutor375@gmail.com",
		Password: "rafli ganteng 123",
	}

	user := repository.NewUserRepository().SignUp(ctx, tx, data)
	fmt.Println("data -> ", user)
}

func TestFindById(t *testing.T) {
	db := config.NewDB()
	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	ctx := context.Background()

	user, err := repository.NewUserRepository().FindById(ctx, tx, "rafly")
	helper.PanicIfError(err)
	fmt.Println("data -> ", user)
}

func TestSignIn(t *testing.T) {
	db := config.NewDB()
	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	
	data := entity.User {
		Email: "xinter@gmail.com",
		Password: "xinter1247",
	}

	hassPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	data.Password = string(hassPassword)

	ctx := context.Background()
	log.Println(data.Password)
	user, err := repository.NewUserRepository().SignIn(ctx, tx, data)
	helper.PanicIfError(err)
	fmt.Println(user)
}

func TestFindAll(t *testing.T) {
	db := config.NewDB()
	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	ctx := context.Background()

	users := repository.NewUserRepository().FindAll(ctx, tx)
	for _, user := range users {
		fmt.Println(user)
	}	
}

func TestDelete(t *testing.T) {
	db := config.NewDB()
	tx, err := db.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	ctx := context.Background()
	repository.NewUserRepository().Delete(ctx, tx, "rafly")
}