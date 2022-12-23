// Work with bitcoin and wallets
package bitcoin

import "fmt"

// A wallet type to allow tracking of the balance
type Wallet struct {
	balance int // The balance of the wallet
}

// Deposit bitcoin in the wallet
func (w *Wallet) Deposit(amt int) error {
	if amt >= 0 {
		w.balance += amt
		return nil
	} else {
		return fmt.Errorf("Cannot deposit a negative amount %d", amt)
	}
}

// Return the balance of the bitcoin in the wallet
func (w *Wallet) Balance() (balance int) {
	return w.balance
}
