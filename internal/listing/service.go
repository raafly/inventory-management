package listing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	"github.com/raafly/inventory-management/pkg/config"
	"github.com/raafly/inventory-management/pkg/helper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserService interface{
	SignUp(request UserSignUp) error
	SignIn(request UserSignIn) (string, error)
} 

type UserServiceImpl struct {
	UserRepository 	UserRepository
	DB 	*sql.DB
	Validate *validator.Validate
}

func NewUserService(userRepository 	UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
	}
}

func generateId() (string, error) {
	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	ms := ulid.Timestamp(time.Now())

	id, err := ulid.New(ms, entropy)
	if err != nil {
		return " ", errors.New("failed generate id")
	}

	uniqueId := id.String()
	return uniqueId, nil
}

func hashedPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("fail hash password")
	}

	return hashedPassword, nil
}

func compareHash(dbPassword, requestPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(requestPassword)); err != nil {
		return errors.New("username and password not match")
	}
	return nil
} 

func (s *UserServiceImpl) SignUp(request UserSignUp) error {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	uniqueId, err := generateId()
	if err != nil {
		return err
	}

	hash, err := hashedPassword(request.Password)
	if err != nil {
		return err
	}

	user := User{
		Id: uniqueId,
		Username: request.Username,
		Email: request.Email,
		Password: string(hash),
	}

	err = s.UserRepository.SignUp(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserServiceImpl) SignIn(request UserSignIn) (string, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return " ", fmt.Errorf("validate missing on %v", err)
	}
	
	data := User{
		Email: request.Email,
		Password: request.Password,
	}

	user, err := s.UserRepository.SignIn(data, s.DB)
	compareHash(user.Password, data.Password)
	
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
	if err != nil {
		return " ", fmt.Errorf("failed create token jwt%v", err.Error())
	}
	
	return token, nil
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

	item := Item {
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
	Save(ctx context.Context, request CategoryCreate) error
	Update(ctx context.Context, request CategoryUpdate) (*Category, error)
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

func (s *CategoryServiceImpl) Save(ctx context.Context, request CategoryCreate) error {
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

func (s *CategoryServiceImpl) Update(ctx context.Context, request CategoryUpdate) (*Category, error) {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	get := Category {
		Id:  request.Id,
		Description: request.Description,
	}

	data, err := s.CategoryRepository.Update(ctx, tx, &get)
	if err != nil {
		return nil, err
	}

	return data, nil
}
