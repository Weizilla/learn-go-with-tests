package pointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("depost", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		assert.Equal(t, expected, actual)

		fmt.Println(expected)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		actual := wallet.Balance()
		expected := Bitcoin(10)

		assert.Nil(t, err)
		assert.Equal(t, expected, actual)

		fmt.Println(expected)
	})

	t.Run("insuff funds", func(t *testing.T) {
		starting := Bitcoin(20)
		wallet := Wallet{starting}
		err := wallet.Withdraw(Bitcoin(200))

		assert.Equal(t, ErrInsufficientFunds, err)
		assert.Equal(t, starting, wallet.Balance())
	})

}
