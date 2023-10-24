package port

import (
	"context"

	"github.com/raafly/inventory-management/model"
)

type UserService interface{
	SignUp(ctx context.Context, request model.UserSignUp) model.UserResponse
	SignIn(ctx context.Context, request model.UserSignIn) (model.UserResponse, string)
	Update(ctx context.Context, request model.UserUpdate)
	Delete(ctx context.Context, userName string) error
	FindById(ctx context.Context, userName string) model.UserResponse
	FindAll(ctx context.Context) []model.UserResponse
} 