package entity

type User struct {
	name  string
	Money *Money
}

func NewUser(name string, money *Money) *User {
	return &User{name: name, Money: money}
}
