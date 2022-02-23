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

var errNomony = errors.New("can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on ytur account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount fron y
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNomony
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of acount
func (a Account) Owner() string {
	return a.owner
}

// sturct 내장 함수
// 출력값 : &{owner value, balance value}
// 커스텀 가능
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas:", a.balance)
}
