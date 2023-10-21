package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/inventory-management/entity"
	"github.com/raafly/inventory-management/helper"
	"github.com/raafly/inventory-management/model"
	portRepository "github.com/raafly/inventory-management/repository/port"
	portService "github.com/raafly/inventory-management/service/port"

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

	item := entity.Item {
		Id: request.Id,
		Name: request.Name,
		Quantity: request.Id,
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
		Id: request.Id,
		Quantity: request.Quantity,
	}

	item = s.ItemRepository.Update(ctx, tx, item)
	return helper.ToItemResponse(item)
}

func (s *ItemServiceImpl) Delete(ctx context.Context, itemId int) {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item, err := s.ItemRepository.FindById(ctx, tx, itemId)
	helper.PanicIfError(err)

	s.ItemRepository.Delete(ctx, tx, item.Id)
}

func (s *ItemServiceImpl) FindById(ctx context.Context, itemId int ) model.ItemResponse {
	tx, err := s.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)	

	item, err := s.ItemRepository.FindById(ctx, tx, itemId)
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

