package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin the new currency
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet defines a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit to deposit money!
func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
}

// Withdraw to withdraw money!
func (w *Wallet) Withdraw(quantity Bitcoin) error {
	tmpBalance := w.balance - quantity
	if tmpBalance < 0 {
		return errors.New("Not enough balance ")
	}
	w.balance = tmpBalance
	return nil
}

// Balance to retrieve the balance!
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
