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

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Error("wanted an error but did not get one")
		}

		if got != want {
			t.Errorf("wanted %q got %q", want, got)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Error("no error expected but got one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		test_wallet := wallet.Wallet{}
		err := test_wallet.Deposit(wallet.Bitcoin(10))
		want := wallet.Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, test_wallet, want)
	})

	t.Run("negative deposit", func(t *testing.T) {
		test_wallet := wallet.Wallet{}
		err := test_wallet.Deposit(wallet.Bitcoin(-10))
		want := wallet.ErrDepositLessThanZero
		assertError(t, err, want)
	})

	t.Run("withdrawl", func(t *testing.T) {
		test_wallet := wallet.Wallet{}
		test_wallet.Deposit(wallet.Bitcoin(10))
		err := test_wallet.Withdraw(wallet.Bitcoin(5))
		want := wallet.Bitcoin(5)
		assertNoError(t, err)
		assertBalance(t, test_wallet, want)
	})

	t.Run("over-withdrawl", func(t *testing.T) {
		test_wallet := wallet.Wallet{}
		test_wallet.Deposit(wallet.Bitcoin(10))
		err := test_wallet.Withdraw(wallet.Bitcoin(15))
		assertError(t, err, wallet.ErrInsufficientFunds)
	})

}
