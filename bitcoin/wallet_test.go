package bitcoin

import (
	"fmt"
	"os"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if want != got {
			t.Errorf("got: %v want: %v", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(10))
		if err != nil {
			t.Errorf("encountered error during deposit")
		}

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("deposit negative", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(-10))
		if err == nil {
			t.Errorf("No error triggered for negative balance deposit")
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(30)}
		err := wallet.Withdraw(Bitcoin(10))
		if err != nil {
			t.Errorf("encountered error during withdrawl")
		}

		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})
	t.Run("overdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(20))
		if err == nil {
			t.Errorf("No error triggered for overdrawn wallet")
		}
	})

}

// An example of how to deposit and withdraw bitcoin from your wallet
// as well as how to output the balance
func ExampleWallet() {
	// Setup the wallet with 10 BTC
	wallet := Wallet{Bitcoin(10)}

	// Deposit 25 BTC
	err := wallet.Deposit(Bitcoin(25))
	if err != nil {
		fmt.Println("Error depositing Bitcoin. Exiting")
		os.Exit(1)
	}

	// Withdraw 15 BTC
	err = wallet.Withdraw(Bitcoin(15))
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
	wallet := Wallet{Bitcoin(10)}

	// Deposit 25 BTC
	err := wallet.Deposit(Bitcoin(25))
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
	wallet := Wallet{Bitcoin(25)}

	// Withdraw 10 BTC
	err := wallet.Withdraw(Bitcoin(10))
	if err != nil {
		fmt.Println("Error withdrawing Bitcoin. Exiting")
		os.Exit(1)
	}

	// Grab the balance and print it
	balance := wallet.Balance()
	fmt.Println(balance)
	// Output: 15 BTC
}
