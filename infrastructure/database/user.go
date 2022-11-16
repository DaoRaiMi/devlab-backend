package database

import (
	"context"

	"github.com/daoraimi/devlab-backend/domain/entity"
	"github.com/daoraimi/devlab-backend/domain/repository"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepoIface {
	return &UserRepo{
		db: db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepo) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return nil, nil
}

func (u *UserRepo) ListUser(ctx context.Context, page, pageSize uint) (entity.UserList, error) {
	return nil, nil
}
