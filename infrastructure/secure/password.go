package secure

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedPwd, plainPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return string(p), nil
}
