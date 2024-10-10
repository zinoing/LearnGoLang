package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney error = errors.New("there isn't enough money to withdraw")

// NewAccount creates account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount of money on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a *Account) GetBalance() int {
	return a.balance
}

// Withdraw x amount of money on your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change owner of the account
func (a *Account) ChangeOwner(owner string) {
	a.owner = owner
}

// Get owner of the account
func (a *Account) GetOwner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s account.\nHas:", a.balance)
}
