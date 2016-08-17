// Package account provides a thread-safe Account API.
package account

import (
	"sync"
)

// Account has unexported state which is accessed via exported methods
type Account struct {
	mu      sync.Mutex
	open    bool
	balance int64
}

// Open creates and initialises a new account, and returns a pointer to it.
// If given a negative initial deposit, it returns nil.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	} else {
		return &Account{open: true, balance: initialDeposit}
	}
}

// Close closes a bank account.
// If any Account method is called on an closed account, it must not modify
// the account and must return ok = false.
func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
  defer a.mu.Unlock()
	if a.open {
		a.open = false
		return a.balance, true
	} else {
		return 0, false
	}
}

// Balance returns the current balance on the account,
// or ok == false if the account has been closed
func (a *Account) Balance() (balance int64, ok bool) {
	a.mu.Lock()
  defer a.mu.Unlock()
	if a.open {
		return a.balance, true
	} else {
		return a.balance, false
	}
}

// Deposit handles a negative amount as a withdrawal.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mu.Lock()
  defer a.mu.Unlock()
	if a.open && a.balance + amount >= 0 {
		a.balance += amount
		return a.balance, true
	} else {
		return a.balance, false
	}
}
