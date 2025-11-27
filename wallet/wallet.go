package wallet

import (
	"fmt"

	"github.com/pkg/errors"
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

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return errors.New("deposit is less than zero")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= w.balance {
		w.balance -= amount
		return nil
	} else {
		return errors.New("insufficient funds")
	}
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
