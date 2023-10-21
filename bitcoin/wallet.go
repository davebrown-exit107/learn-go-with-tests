package bitcoin

import (
	"fmt"
)

type errorConst string

const (
	ErrOverdrawn     errorConst = "Account Overdrawn"
	ErrInvalidAmount errorConst = "Invalid amount"
)

func (e errorConst) Error() string {
	return string(e)
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount >= 0 {
		w.balance += amount
		return nil
	} else {
		return ErrInvalidAmount
	}
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrOverdrawn
	}
	if amount >= 0 {
		w.balance -= amount
		return nil
	} else {
		return ErrInvalidAmount
	}
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
