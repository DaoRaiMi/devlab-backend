package repository

import (
	"context"

	"github.com/daoraimi/devlab-backend/domain/entity"
)

type UserRepoIface interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	ListUser(ctx context.Context, page, pageSize uint) ([]*entity.User, error)
}
