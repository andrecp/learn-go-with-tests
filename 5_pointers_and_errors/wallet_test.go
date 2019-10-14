package wallet

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit test", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw test", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(40)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(30))
	})

	t.Run("Withdraw more than allowed", func(t *testing.T) {
		startingBalance := Bitcoin(40)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		if err == nil {
			t.Errorf("Should error when trying to remove more balance.")
		}
	})
}
