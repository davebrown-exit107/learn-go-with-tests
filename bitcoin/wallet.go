package bitcoin

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	if amount >= 0 {
		w.balance += amount
	}
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	if amount >= 0 {
		w.balance -= amount
	}
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
