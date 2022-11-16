package secure

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// %s: uuid, %d: current timestamp in nano seconds, %d: user id
	plainTokenFmt = "devlab_%s_%d_%d"
)

func GenerateUserToken(uid uint) string {
	u := uuid.New()
	s := fmt.Sprintf(plainTokenFmt, u.String(), time.Now().Nanosecond(), uid)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
