package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		fmt.Printf("Balance %s", wallet.Balance())
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(90)

		fmt.Printf("Balance %s", wallet.Balance())
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw insufficent funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}
