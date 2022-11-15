package application

import (
	"context"

	"github.com/daoraimi/devlab-backend/domain/repository"
	"github.com/daoraimi/devlab-backend/domain/value"
	"github.com/go-redis/redis"
)

type UserAppIface interface {
	Login(context.Context, value.LoginRequest) (value.LoginResponse, error)
	Logout(context.Context) error
}

type UserApp struct {
	userRepo repository.UserRepoIface
	redis    *redis.Client
}

func NewUserApp(u repository.UserRepoIface, redis *redis.Client) UserAppIface {
	return &UserApp{
		userRepo: u,
		redis:    redis,
	}
}

func (u *UserApp) Login(ctx context.Context, req value.LoginRequest) (value.LoginResponse, error) {
	return *value.NewLoginResponse(200, "Success", "asdfjsdjfii92sdfxcvadgfas"), nil
}

func (u *UserApp) Logout(ctx context.Context) error {
	return nil
}
