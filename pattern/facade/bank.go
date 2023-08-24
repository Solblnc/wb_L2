package main

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

type Card struct {
	Name    string
	Balance int
	Bank    *Bank
}

func (b *Bank) CheckBalance(cardNumber string) error {
	fmt.Printf("[Bank] Get balance of the card: %s", cardNumber)
	time.Sleep(1 * time.Second)
	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		}

		if card.Balance <= 0 {
			return errors.New("\n[Bank] Insufficient funds")
		}
	}
	fmt.Println("\nsufficient funds!")
	return nil
}

func (c *Card) GetBalance() error {
	fmt.Println("[Card] Request bank for balance on the card...")
	time.Sleep(1 * time.Second)
	return c.Bank.CheckBalance(c.Name)
}
