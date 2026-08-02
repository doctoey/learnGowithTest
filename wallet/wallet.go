package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

// ก่อนใช้ pointer
// address of balance in Deposit is 0x4d79dcc8e0b0
// address of balance in test is 0x4d79dcc8e0a8
// คนละตัวกัน

// หลังใช้ pointer
// address of balance in Deposit is 0x4d79dcc8e0b0
// address of balance in test is 0x4d79dcc8e0b0

// หลักการของ pointer
// 1. * คือการ dereference คือการเข้าถึงค่าของ pointer
// 2. & คือการ reference คือการเข้าถึง address ของตัวแปร
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}

	fmt.Printf("address of balance in Withdraw is %p \n", &w.balance)
	w.balance -= amount
	return nil
}
