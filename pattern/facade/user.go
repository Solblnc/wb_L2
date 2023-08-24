package main

type User struct {
	Name string
	Card *Card
}

func (u *User) GetBalance() int {
	return u.Card.Balance
}
