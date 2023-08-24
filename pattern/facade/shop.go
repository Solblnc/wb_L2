package main

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	Name  string
	Price int
}

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) Sell(user User, product string) error {
	fmt.Println("\n[Shop] Request for user for card balance...")
	time.Sleep(1 * time.Second)
	err := user.Card.GetBalance()
	if err != nil {
		return err
	}
	fmt.Printf("\n[Shop] Checking if user %s can buy a product?....", user.Name)
	time.Sleep(1 * time.Second)
	for _, prod := range s.Products {
		if prod.Name != product {
			continue
		}

		if prod.Price > user.GetBalance() {
			return errors.New("[Shop] Insufficient funds for buy this product(((")
		}
		fmt.Printf("\n[Shop] Product %s is bought!!!\n", prod.Name)
	}
	return nil
}
