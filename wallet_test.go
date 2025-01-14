package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)

		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("test deposit negative amount", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(-20.0)

		assertError(t, err)
	})

	t.Run("test withdrawal", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(50.0)
		wallet.Withdraw(40.0)

		assertBalance(t, wallet, Bitcoin(10.0))
	})

	t.Run("test withdrawal over availability", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(50.0)
		err := wallet.Withdraw(60.0)

		assertError(t, err)
	})
}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("expected error but got nil")
		t.Fatal()
	}
}

func assertBalance(t testing.TB, w Wallet, balance Bitcoin) {
	t.Helper()

	if got := w.Balance(); got != balance {
		t.Errorf("got %s, want %s", got, balance)
	}
}
