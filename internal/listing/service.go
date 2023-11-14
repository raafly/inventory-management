package listing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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

	user, err := s.UserRepository.SignIn(data)
	compareHash(user.Password, data.Password)
	
	expTime := time.Now().Add(time.Hour * 1)
	claims := &config.JWTClaims{
		Username: user.Username,
		Email: user.Email,
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
	Create(request ItemCreate) error
	UpdateStatus(requst ItemUpdate) error
	UpdateQuantity(request ItemUpdate) 
	Delete(itemId int) error
	FindById(itemId int) (*ItemResponse, error)
	FindAll() []ItemResponse
	// UpadteDescription(request ItemUpdate)
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

func (s ItemServiceImpl) Create(request ItemCreate) error {
	if err := s.Validate.Struct(request); err != nil {
		return fmt.Errorf("validate error %v", err.Error())
	}

	item := Item {
		Name: request.Name,
		Category: request.Category,
		Quantity: request.Quantity,
	}

	if err := s.ItemRepository.Create(item); err != nil {
		return err
	}
	return nil
}

func (s ItemServiceImpl) FindById(itemId int) (*ItemResponse, error) {
	var itemRes ItemResponse

	if item, err := s.ItemRepository.FindById(itemId); err != nil {
		return &itemRes, fmt.Errorf("id item not found %v", err.Error()) 
	} else {
		itemRes = ItemResponse{
			Id: item.Id,
			Name: item.Name,
			Description: item.Description,
			Quantity: item.Quantity,
			Status: item.Status,
			Category: item.Category,
			Created_at: item.Created_at,
		}

		return &itemRes, nil
	}
}

func (s ItemServiceImpl) UpdateStatus(request ItemUpdate) error {
	if err := s.Validate.Struct(request); err != nil {
		return fmt.Errorf("validate error %v", err.Error())
	}

	data, err := s.ItemRepository.FindById(request.Id)
	if err  != nil {
		panic(NewNotFoundError(err.Error()))
	}

	data = &Item {
		Id: data.Id,
		Status: request.Status,
	}

	s.ItemRepository.UpdateStatus(data.Id, data.Status)
	return nil
}

func (s ItemServiceImpl) UpdateQuantity(request ItemUpdate) {
	if err := s.Validate.Struct(request); err != nil {
		fmt.Printf("validate error %v", err.Error())
	}

	data, err := s.ItemRepository.FindById(request.Id)
	if err  != nil {
		panic(NewNotFoundError(err.Error()))
	}

	data = &Item {
		Id: data.Id,
		Quantity: request.Quantity,
	}
	log.Print(data)

	s.ItemRepository.UpdateQuantity(data.Id, data.Quantity)
}

func (s ItemServiceImpl) Delete(itemId int) error {
	s.ItemRepository.Delete(itemId)
	return nil
}

func (s ItemServiceImpl) FindAll() []ItemResponse {
	items := s.ItemRepository.FindAll()
	
	var itemResponse []ItemResponse
	for _, item := range items {
		itemResponse = append(itemResponse, ItemResponse(item))
	}
	
	return itemResponse
}

/*


func (s ItemServiceImpl) UpadteDescription(request ItemUpdate) {
	if item, err := s.ItemRepository.FindById(request.Id); err != nil {
		fmt.Errorf("id item not found %v", err.Error())
	} else {
		s.ItemRepository.UpadteDescription(item.Id, item.Description)
	}	
}

*/

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