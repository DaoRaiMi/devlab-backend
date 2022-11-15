package value

import "errors"

const (
	TokenExpirationSeconds = 86400
	UserTokenKeyFmt        = ""
	TokenToUserFmt         = ""
)

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
