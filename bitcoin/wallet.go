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
func (w *Wallet) Deposit(amt Bitcoin) error {
	if amt < Bitcoin(0) {
		return fmt.Errorf("Cannot deposit negative bitcoin in the amt of %v", amt)
	}
	w.balance += amt
	return nil
}

// Withdraw bitcoin fron the wallet
func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return fmt.Errorf("Wallet overdrawn: Cannot withdraw %v from wallet with %v balance", amt, w.balance)
	}
	w.balance -= amt
	return nil
}

// Return the balance of the bitcoin in the wallet
func (w *Wallet) Balance() (balance Bitcoin) {
	return w.balance
}
