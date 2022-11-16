package value

import (
	"errors"
	"fmt"
)

const (
	// token expiration time
	TokenExpirationSeconds = 86400
	// max failed login count
	MaxFailedLoginCount = 3
	// max failed login count expiration time
	FailedLoginCountKeyExpirationSeconds = 300
	uidToTokenKeyFmt                     = "uid_to_token_%d"
	tokenToUidKeyFmt                     = "token_to_uid_%s"
	failedLoginCountKeyFmt               = "failed_login_count_%s"
)

var (
	ErrMaxFailedLoginCountReached = errors.New("max failed login count reached")
	ErrUsernameOrPasswordIsWrong  = errors.New("username or password is wrong")
)

func GetUidToTokenKey(id uint) string {
	return fmt.Sprintf(uidToTokenKeyFmt, id)
}

func GetTokenToUidKey(token string) string {
	return fmt.Sprintf(tokenToUidKeyFmt, token)
}

func GetFailedLoginCountKey(username string) string {
	return fmt.Sprintf(failedLoginCountKeyFmt, username)
}

// login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l LoginRequest) Validate() error {
	if l.Username == "" || l.Password == "" {
		return errors.New("username and password are required")
	}
	return nil
}

type LoginResponse struct {
	BaseResponse `json:",inline"`
	Token        string `json:"token,omitempty"`
}

func NewLoginResponse(code int, msg string, token string) *LoginResponse {
	return &LoginResponse{
		BaseResponse: BaseResponse{
			Code:    code,
			Message: msg,
		},
		Token: token,
	}
}
