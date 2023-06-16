package model

import (
	"crypto/sha256"
	"encoding/base64"
)

type User struct {
	ID             string
	Name           string
	HashedPassword string
}

func (u *User) IsSamePassword(password string) bool {
	hash := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(hash[:]) == u.HashedPassword
}
