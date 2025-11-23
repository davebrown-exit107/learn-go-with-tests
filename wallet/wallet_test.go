package wallet_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/wallet"
)

func TestWallet(t *testing.T) {
	test_wallet := wallet.Wallet{}

	test_wallet.Deposit(wallet.Bitcoin(10))

	got := test_wallet.Balance()
	want := wallet.Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
