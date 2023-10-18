package bitcoin

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	if amount >= 0 {
		w.balance += amount
	}
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}