package user

import (
	"http/web"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type password string

func (p password) isValid() error {
	if len(p) < 6 {
		return web.BadRequest("Password must be at least 6 character")
	}
	return nil
}

func (p password) Hash() (string, error) {
	if len(os.Getenv("nohash")) > 0 {
		return string(p), nil
	}
	var s, err = bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		objectUserLog.Error(err)
		return "", web.InternalServerError("generate hash password failed")
	}
	return string(s), nil
}

func (p password) Compare(hash string) error {
	var err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return web.Unauthorized("wrong password")
		}
		objectUserLog.Error(err)
		return web.InternalServerError("compare hash password failed")
	}
	return nil
}

// HashTo check and hash the password
func (p password) HashTo(s *string) error {
	if err := p.isValid(); err != nil {
		return err
	}
	if hash, err := p.Hash(); err != nil {
		return err
	} else {
		*s = hash
		return nil
	}
}

func (u *User) ComparePassword(p string) error {
	return password(p).Compare(u.Password)
}
