// Work with bitcoin and wallets
package bitcoin

import (
	"errors"
	"fmt"
)

var ErrNegativeDeposit = errors.New("cannot deposit negative bitcoin")
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// The currency bitcoin
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// A wallet type to allow tracking of the balance
type Wallet struct {
	balance Bitcoin // The balance of the wallet
}

// Deposit bitcoin in the wallet
func (w *Wallet) Deposit(amt Bitcoin) error {
	if amt < Bitcoin(0) {
		return ErrNegativeDeposit
	}
	w.balance += amt
	return nil
}

// Withdraw bitcoin fron the wallet
func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}

// Return the balance of the bitcoin in the wallet
func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}
