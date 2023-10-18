package port

import (
	"context"
	"database/sql"

	"github.com/raafly/inventory-management/entity"
)

type UserPort interface {
	SignUp(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	SignIn(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Update(ctx context.Context, tx *sql.Tx, user entity.User)
	Delete(ctx context.Context, tx *sql.Tx, userId int)
	FindById(ctx context.Context, tx *sql.Tx, userId int) (entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.User
}