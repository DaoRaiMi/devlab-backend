package application

import (
	"context"

	"github.com/daoraimi/devlab-backend/domain/repository"
	"github.com/daoraimi/devlab-backend/domain/value"
	"github.com/daoraimi/devlab-backend/infrastructure/secure"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type UserAppIface interface {
	Login(context.Context, value.LoginRequest) (*value.LoginResponse, error)
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

func (u *UserApp) Login(ctx context.Context, req value.LoginRequest) (*value.LoginResponse, error) {
	// get login failed count
	count, err := u.redis.Get(value.GetFailedLoginCountKey(req.Username)).Uint64()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if count > value.MaxFailedLoginCount {
		return nil, value.ErrMaxFailedLoginCountReached
	}

	// get username and hashed password from database
	user, err := u.userRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if user == nil {
		return nil, value.ErrUsernameOrPasswordIsWrong
	}

	failedLoginCountKey := value.GetFailedLoginCountKey(req.Username)
	// compare plain and hashed password
	if secure.ComparePasswords(user.Password, req.Password) {

		// increase the failed login count
		if _, err := u.redis.Incr(failedLoginCountKey).Result(); err != nil {
			return nil, errors.WithStack(err)
		}
		// set the expiration of failed login count key
		if _, err := u.redis.Expire(failedLoginCountKey, value.FailedLoginCountKeyExpirationSeconds).Result(); err != nil {
			return nil, errors.WithStack(err)
		}

		return nil, value.ErrUsernameOrPasswordIsWrong
	}

	// delete the failed login count key
	if _, err = u.redis.Del(failedLoginCountKey).Result(); err != nil {
		return nil, errors.WithStack(err)
	}

	return value.NewLoginResponse(200, "Success", secure.GenerateUserToken(user.ID)), nil
}

func (u *UserApp) Logout(ctx context.Context) error {
	return nil
}
