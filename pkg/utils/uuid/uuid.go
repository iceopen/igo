package uuid

import (
	"errors"

	"github.com/google/uuid"
)

// GetUUID 获取 uuid
func GetUUID() (string, error) {
	u, e := uuid.NewUUID()
	if e != nil {
		return "", errors.New("new uuid v4 error")
	}
	return u.String(), nil
}
