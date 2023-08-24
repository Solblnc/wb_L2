package main

import "fmt"

func main() {
	bank := Bank{
		Name:  "SBER",
		Cards: []Card{},
	}

	card1 := Card{
		Name:    "card1",
		Balance: 400,
		Bank:    &bank,
	}

	card2 := Card{
		Name:    "card2",
		Balance: 100,
		Bank:    &bank,
	}

	user1 := User{
		Name: "user1",
		Card: &card1,
	}

	user2 := User{
		Name: "user2",
		Card: &card2,
	}

	prod := Product{
		Name:  "Phone",
		Price: 300,
	}

	shop := Shop{
		Name:     "TestShop",
		Products: []Product{prod},
	}

	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Println("USER 1")
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("USER 2")
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
