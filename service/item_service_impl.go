package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/oklog/ulid/v2"
	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
	portRepository "github.com/raafly/inventory-management/repository/port"
	portService "github.com/raafly/inventory-management/service/port"
	"golang.org/x/exp/rand"
)

type ItemServiceImpl struct {
	ItemRepository		portRepository.ItemRepository
	DB 					*sql.DB
	Validate			*validator.Validate
}

func NewItemService(itemRepository portRepository.ItemRepository, DB *sql.DB, validate *validator.Validate) portService.ItemService {
	return &ItemServiceImpl{
		ItemRepository: itemRepository,
		DB: DB,
		Validate: validate,
	}
}

func (s *ItemServiceImpl) Create(ctx context.Context, request model.ItemCreate) model.ItemResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	uniqueId := id.String()

	item := entity.Item {
		Id: uniqueId,
		Name: request.Name,
		Category: request.Category,
		Quantity: request.Quantity,
	}

	item = s.ItemRepository.Create(ctx, tx, item)
	return helper.ToItemResponse(item)
}

func (s *ItemServiceImpl) Update(ctx context.Context, request model.ItemUpdate) model.ItemResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item := entity.Item {
		Name: request.Name,
		Quantity: request.Quantity,
	}

	item = s.ItemRepository.Update(ctx, tx, item)
	return helper.ToItemResponse(item)
}

func (s *ItemServiceImpl) Delete(ctx context.Context, itemName string) {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item, err := s.ItemRepository.FindById(ctx, tx, itemName)
	helper.PanicIfError(err)

	s.ItemRepository.Delete(ctx, tx, item.Name)
}

func (s *ItemServiceImpl) FindById(ctx context.Context, itemName string ) model.ItemResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item, err := s.ItemRepository.FindById(ctx, tx, itemName)
	helper.PanicIfError(err)
	
	return helper.ToItemResponse(item)
}

func (s *ItemServiceImpl) FindAll(ctx context.Context) []model.ItemResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	items := s.ItemRepository.FindAll(ctx, tx)
	return helper.ToItemResponses(items)
}

