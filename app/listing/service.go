package listing

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"github.com/raafly/inventory-management/config"
	"github.com/raafly/inventory-management/helper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserService interface{
	SignUp(ctx context.Context, request UserSignUp) UserResponse
	SignIn(ctx context.Context, request UserSignIn) (UserResponse, string)
	Update(ctx context.Context, request UserUpdate)
	Delete(ctx context.Context, userName string) error
	FindById(ctx context.Context, userName string) UserResponse
	FindAll(ctx context.Context) []UserResponse
} 

type UserServiceImpl struct {
	UserRepository 	UserRepository
	DB 				*sql.DB
	Validate 		*validator.Validate
}

func NewUserService(userRepository 	UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
	}
}


func (s *UserServiceImpl) SignUp(ctx context.Context, request UserSignUp) UserResponse {
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

	user := User{
		Id: uniqueId,
		Username: request.Username,
		Email: request.Email,
		Password: string(hashedPassword),
	}

	user = s.UserRepository.SignUp(ctx, tx, user)
	return ToUserResponse(user)
}

func (s *UserServiceImpl) SignIn(ctx context.Context, request UserSignIn) (UserResponse, string) {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	data := User{
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

	return ToUserResponse(user), token
}

func (s *UserServiceImpl) Update(ctx context.Context, request UserUpdate) {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	user := User {
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

func (s *UserServiceImpl) FindById(ctx context.Context, userName string) UserResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	user, err := s.UserRepository.FindById(ctx, tx, userName)
	helper.PanicIfError(err)

	return ToUserResponse(user)
}

func (s *UserServiceImpl) FindAll(ctx context.Context) []UserResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	users := s.UserRepository.FindAll(ctx, tx)
	return ToUserResponses(users)
}

// item

type ItemService interface {
	Create(ctx context.Context, request ItemCreate) ItemResponse
	Update(ctx context.Context, requst ItemUpdate) ItemResponse
	Delete(ctx context.Context, itemName string)
	FindById(ctx context.Context, itemName string ) ItemResponse
	FindAll(ctx context.Context) []ItemResponse
}

type ItemServiceImpl struct {
	ItemRepository		ItemRepository
	DB 					*sql.DB
	Validate			*validator.Validate
}

func NewItemService(itemRepository ItemRepository, DB *sql.DB, validate *validator.Validate) ItemService {
	return &ItemServiceImpl{
		ItemRepository: itemRepository,
		DB: DB,
		Validate: validate,
	}
}

func (s *ItemServiceImpl) Create(ctx context.Context, request ItemCreate) ItemResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	uniqueId := id.String()

	item := Item {
		Id: uniqueId,
		Name: request.Name,
		Category: request.Category,
		Quantity: request.Quantity,
	}

	item = s.ItemRepository.Create(ctx, tx, item)
	return ToItemResponse(item)
}

func (s *ItemServiceImpl) Update(ctx context.Context, request ItemUpdate) ItemResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item := Item {
		Name: request.Name,
		Quantity: request.Quantity,
	}

	item = s.ItemRepository.Update(ctx, tx, item)
	return ToItemResponse(item)
}

func (s *ItemServiceImpl) Delete(ctx context.Context, itemName string) {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item, err := s.ItemRepository.FindById(ctx, tx, itemName)
	helper.PanicIfError(err)

	s.ItemRepository.Delete(ctx, tx, item.Name)
}

func (s *ItemServiceImpl) FindById(ctx context.Context, itemName string ) ItemResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item, err := s.ItemRepository.FindById(ctx, tx, itemName)
	helper.PanicIfError(err)
	
	return ToItemResponse(item)
}

func (s *ItemServiceImpl) FindAll(ctx context.Context) []ItemResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	items := s.ItemRepository.FindAll(ctx, tx)
	return ToItemResponses(items)
}

// category

type CategoryService interface {
	CreateCtg(ctx context.Context, request CategoryCreate) error
}

type CategoryServiceImpl struct {
	CategoryRepository 	CategoryRepository
	DB 					*sql.DB
	Validate 			*validator.Validate
}

func NewCategoryService(categoryRepository CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl {
		CategoryRepository: categoryRepository,
		DB: DB,
		Validate: validate,
	}
}

func (s *CategoryServiceImpl) CreateCtg(ctx context.Context, request CategoryCreate) error {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	data := Category {
		Id: request.Id,
		Name: request.Name,
		Description: request.Description,
	}

	s.CategoryRepository.Create(ctx, tx, data)

	return nil
}
