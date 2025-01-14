package main

import (
	"errors"
	"fmt"
)

type Bitcoin float64

var WithdrawalError = errors.New("Cannot withdraw more than the available balance")
var NegativeDepositError = errors.New("Cannot deposit a negative amount")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return NegativeDepositError
	}

	w.balance += amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return WithdrawalError
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
