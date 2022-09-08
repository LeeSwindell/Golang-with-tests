package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("not enough dough, mate")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	
	w.balance -= amount
	return nil
}