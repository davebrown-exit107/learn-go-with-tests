package bitcoin_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/bitcoin"
)

func TestWallet(t *testing.T) {
	wallet := bitcoin.Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := bitcoin.Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
