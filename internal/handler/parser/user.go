package parser

type User struct {
	ID             int
	Name           string
	HashedPassword string
}

func (u *User) IsSamePassword(password string) bool {
	return u.HashedPassword == password
}
