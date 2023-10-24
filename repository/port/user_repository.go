package port

import (
	"context"
	"database/sql"

	"github.com/raafly/inventory-management/entity"
)

type UserRepository interface {
	SignUp(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	SignIn(ctx context.Context, tx *sql.Tx, user entity.User) (entity.User, error)
	Update(ctx context.Context, tx *sql.Tx, user entity.User)
	Delete(ctx context.Context, tx *sql.Tx, userName string) error
	FindById(ctx context.Context, tx *sql.Tx, userName string) (entity.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.User
}