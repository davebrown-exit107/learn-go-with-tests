package bitcoin_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/bitcoin"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet bitcoin.Wallet, want bitcoin.Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {

		wallet := bitcoin.Wallet{}

		wallet.Deposit(bitcoin.Bitcoin(10))

		want := bitcoin.Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := bitcoin.Wallet{}
		wallet.Deposit(20)
		wallet.Withdraw(10)

		want := bitcoin.Bitcoin(10)

		assertBalance(t, wallet, want)
	})
}
