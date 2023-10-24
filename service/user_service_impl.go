package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"github.com/raafly/inventory-management/config"
	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
	portRepository "github.com/raafly/inventory-management/repository/port"
	portService "github.com/raafly/inventory-management/service/port"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserServiceImpl struct {
	UserRepository 	portRepository.UserRepository
	DB 				*sql.DB
	Validate 		*validator.Validate
}

func NewUserService(userRepository 	portRepository.UserRepository, DB *sql.DB, validate *validator.Validate) portService.UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
	}
}


func (s *UserServiceImpl) SignUp(ctx context.Context, request model.UserSignUp) model.UserResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	uniqueId := id.String()

	user := entity.User{
		Id: uniqueId,
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
	}

	user = s.UserRepository.SignUp(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (s *UserServiceImpl) SignIn(ctx context.Context, request model.UserSignIn) (model.UserResponse, string) {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	data := entity.User{
		Email: request.Email,
		Password: request.Password,
	}

	user, err := s.UserRepository.SignIn(ctx, tx, data)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	helper.PanicIfError(err)

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: user.Username,
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user), token
}

func (s *UserServiceImpl) Update(ctx context.Context, request model.UserUpdate) {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	user := entity.User {
		Email: request.Email,
		Password: request.Password,
	}

	s.UserRepository.Update(ctx, tx, user)
}

func (s *UserServiceImpl) Delete(ctx context.Context, userName string) error {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	err = s.UserRepository.Delete(ctx, tx, userName)
	if err != nil {
		return errors.New("username not found")
	} else {
		return nil
	}
}

func (s *UserServiceImpl) FindById(ctx context.Context, userName string) model.UserResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	user, err := s.UserRepository.FindById(ctx, tx, userName)
	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}

func (s *UserServiceImpl) FindAll(ctx context.Context) []model.UserResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	users := s.UserRepository.FindAll(ctx, tx)
	return helper.ToUserResponses(users)
}