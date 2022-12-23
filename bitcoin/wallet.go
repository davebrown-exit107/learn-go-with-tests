// Work with bitcoin and wallets
package bitcoin

import "fmt"

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
func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

// Withdraw bitcoin fron the wallet
func (w *Wallet) Withdraw(amt Bitcoin) {
	w.balance -= amt
}

// Return the balance of the bitcoin in the wallet
func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}
