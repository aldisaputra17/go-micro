package helper

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (hash string, err error) {
	cost, err := strconv.Atoi(os.Getenv("AUTH_BCRYPT_COST"))
	if err != nil {
		return
	}

	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}

	crypt, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		err = errors.Wrap(err, "error generate bcrypt hash password")
		return
	}
	hash = string(crypt)

	return
}
