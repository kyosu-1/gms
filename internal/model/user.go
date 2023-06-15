package model

type User struct {
	ID             string
	Name           string
	HashedPassword string
}

func (u *User) IsSamePassword(password string) bool {
	return u.HashedPassword == password
}
