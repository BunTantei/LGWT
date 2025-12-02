// Package pointers is everything to learn about pointers but also errors.
package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds.")

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("The amount deposit is %d \n", amount)
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
	fmt.Printf("The balance after deposit is %d \n", w.balance)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}
	fmt.Printf("amount withdraw is %d \n", amount)
	w.balance -= amount
	fmt.Printf("the balance after withdrawing is %d \n", w.balance)
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
