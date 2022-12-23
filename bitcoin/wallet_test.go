package bitcoin

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(30)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(20)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

// An example of how to deposit and withdraw bitcoin from your wallet
// as well as how to output the balance
func ExampleWallet() {
	// Setup the wallet with 10 BTC
	wallet := Wallet{Bitcoin(10)}

	// Deposit 25 BTC
	wallet.Deposit(Bitcoin(25))

	// Withdraw 15 BTC
	wallet.Withdraw(Bitcoin(15))

	// Grab the balance and print it
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 20 BTC
}

func ExampleWallet_Deposit() {
	// Setup the wallet with 10 BTC
	wallet := Wallet{Bitcoin(10)}

	// Deposit 25 BTC
	wallet.Deposit(Bitcoin(25))

	// Grab the balance and print it
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 35 BTC
}

func ExampleWallet_Withdraw() {
	// Setup the wallet with 25 BTC
	wallet := Wallet{Bitcoin(25)}

	// Withdraw 10 BTC
	wallet.Withdraw(Bitcoin(10))

	// Grab the balance and print it
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 15 BTC
}
