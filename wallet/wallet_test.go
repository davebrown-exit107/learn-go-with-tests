package wallet_test

import (
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/wallet"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet wallet.Wallet, want wallet.Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		test_wallet := wallet.Wallet{}

		test_wallet.Deposit(wallet.Bitcoin(10))

		want := wallet.Bitcoin(10)

		assertBalance(t, test_wallet, want)
	})

	t.Run("withdrawl", func(t *testing.T) {
		test_wallet := wallet.Wallet{}
		test_wallet.Deposit(wallet.Bitcoin(10))

		test_wallet.Withdraw(wallet.Bitcoin(5))

		want := wallet.Bitcoin(5)

		assertBalance(t, test_wallet, want)
	})

}
