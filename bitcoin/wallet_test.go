package bitcoin_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/davebrown-exit107/learn-go-with-tests/bitcoin"
)

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Error("wanted nil but got an error")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted error but got nil")
	}
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertBalance(t testing.TB, wallet bitcoin.Wallet, want bitcoin.Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if want != got {
		t.Errorf("got: %v want: %v", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := bitcoin.Wallet{}
		err := wallet.Deposit(bitcoin.Bitcoin(10))
		want := bitcoin.Bitcoin(10)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("deposit negative", func(t *testing.T) {
		wallet := bitcoin.Wallet{}
		wallet.SetBalance(bitcoin.Bitcoin(10))
		err := wallet.Deposit(bitcoin.Bitcoin(-10))

		assertError(t, err, bitcoin.ErrNegativeDeposit)
		assertBalance(t, wallet, bitcoin.Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := bitcoin.Wallet{}
		wallet.SetBalance(bitcoin.Bitcoin(30))
		err := wallet.Withdraw(bitcoin.Bitcoin(10))
		want := bitcoin.Bitcoin(20)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})
	t.Run("overdraw", func(t *testing.T) {
		wallet := bitcoin.Wallet{}
		wallet.SetBalance(bitcoin.Bitcoin(10))
		err := wallet.Withdraw(bitcoin.Bitcoin(20))

		assertError(t, err, bitcoin.ErrInsufficientFunds)
		assertBalance(t, wallet, bitcoin.Bitcoin(10))
	})

}

// An example of how to deposit and withdraw bitcoin from your wallet
// as well as how to output the balance
func ExampleWallet() {
	// Setup the wallet with 10 BTC
	wallet := bitcoin.Wallet{}
	wallet.SetBalance(bitcoin.Bitcoin(10))

	// Deposit 25 BTC
	err := wallet.Deposit(bitcoin.Bitcoin(25))
	if err != nil {
		fmt.Println("Error depositing Bitcoin. Exiting")
		os.Exit(1)
	}

	// Withdraw 15 BTC
	err = wallet.Withdraw(bitcoin.Bitcoin(15))
	if err != nil {
		fmt.Println("Error withdrawing Bitcoin. Exiting")
		os.Exit(1)
	}

	// Grab the balance and print it
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 20 BTC
}

func ExampleWallet_Deposit() {
	// Setup the wallet with 10 BTC
	wallet := bitcoin.Wallet{}
	wallet.SetBalance(bitcoin.Bitcoin(10))

	// Deposit 25 BTC
	err := wallet.Deposit(bitcoin.Bitcoin(25))
	if err != nil {
		fmt.Println("Error depositing Bitcoin. Exiting")
		os.Exit(1)
	}

	// Grab the balance and print it
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 35 BTC
}

func ExampleWallet_Withdraw() {
	// Setup the wallet with 25 BTC
	wallet := bitcoin.Wallet{}
	wallet.SetBalance(bitcoin.Bitcoin(25))

	// Withdraw 10 BTC
	err := wallet.Withdraw(bitcoin.Bitcoin(10))
	if err != nil {
		fmt.Println("Error withdrawing Bitcoin. Exiting")
		os.Exit(1)
	}

	// Grab the balance and print it
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 15 BTC
}
