package main

import (
	"errors"
	"fmt"
)

// errors
var (
	ErrInvalidAmount     = errors.New("Invalid amount")
	ErrInsufficientFunds = errors.New("Insufficient funds")
)

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type BankAccount struct {
	id      string
	owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	b.balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("account %s: %w", b.id, ErrInvalidAmount)
	}
	if b.balance < amount {
		return fmt.Errorf("account %s: %w", b.id, ErrInsufficientFunds)
	}
	b.balance -= amount
	return nil
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

type PremiumAccount struct {
	BankAccount
	cashbackRate float64
}

func (p *PremiumAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("account %s: %w", p.id, ErrInvalidAmount)
	}
	p.balance += amount
	return nil
}

func (p *PremiumAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("account %s: %w", p.id, ErrInvalidAmount)
	}
	if p.balance < amount {
		return fmt.Errorf("account %s: %w", p.id, ErrInsufficientFunds)
	}
	p.balance -= amount

	cashback := amount * p.cashbackRate
	p.balance += cashback
	return nil
}

func (p *PremiumAccount) Cashback() float64 {
	return p.balance * p.cashbackRate
}

func main() {
	var account Account

	account = &BankAccount{
		id:      "ACC-001",
		owner:   "Rustam",
		balance: 500,
	}

	account.Deposit(1000)

	fmt.Println(account.Balance())

	account = &PremiumAccount{
		BankAccount: BankAccount{
			id:      "ACC-002",
			owner:   "Ali",
			balance: 500,
		},
		cashbackRate: 0.02,
	}

	account.Deposit(1000)
	fmt.Println(account.Balance())
}
