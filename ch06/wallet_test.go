package bitcoin_test

import (
	"fmt"
	"testing"

	bitcoin "github.com/davebrown-exit107/learn-go-with-tests/ch06"
)

func TestWallet(t *testing.T) {
	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Errorf("wanted error %q, got nil", want)
		}
		if got != want {
			t.Errorf("wanted error %q, got %q", want, got)
		}
	}

	assertBalance := func(t testing.TB, wallet bitcoin.Wallet, want bitcoin.Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit - normal", func(t *testing.T) {

		wallet := bitcoin.Wallet{}

		_ = wallet.Deposit(bitcoin.Bitcoin(10))

		want := bitcoin.Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("deposit - invalid amount", func(t *testing.T) {

		wallet := bitcoin.Wallet{}

		got := wallet.Deposit(bitcoin.Bitcoin(-10))

		assertError(t, got, bitcoin.ErrInvalidAmount)
	})

	t.Run("withdraw - normal", func(t *testing.T) {

		wallet := bitcoin.Wallet{}
		_ = wallet.Deposit(20)
		_ = wallet.Withdraw(10)

		want := bitcoin.Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw - invalid amount", func(t *testing.T) {

		wallet := bitcoin.Wallet{}
		_ = wallet.Deposit(20)
		got := wallet.Withdraw(-10)

		assertError(t, got, bitcoin.ErrInvalidAmount)
	})

	t.Run("withdraw - overdrawn", func(t *testing.T) {

		wallet := bitcoin.Wallet{}
		_ = wallet.Deposit(20)
		got := wallet.Withdraw(30)

		assertError(t, got, bitcoin.ErrOverdrawn)
	})
	t.Run("stringer", func(t *testing.T) {
		wallet := bitcoin.Wallet{}
		_ = wallet.Deposit(20)

		got := fmt.Sprintf("%s", wallet.Balance())
		want := "20 BTC"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
