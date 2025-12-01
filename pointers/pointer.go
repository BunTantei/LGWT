// Package pointers is everything to learn about pointers but also errors.
package pointers

type Wallet struct {
	balance int
}

func (w Wallet) Deposit(amount int) {

}

func (w Wallet) Balance() int {
	return w.balance
}
