package main

import (
	"fmt"
	"sync"
	"time"
)

const count = 100

type Person struct {
	ID   uint64
	Name string
	Age  uint8
}

type Account struct {
	id      uint64
	owner   Person
	balance int
	lock    sync.RWMutex
}

func (a *Account) WithdrawMoney(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance -= amount
}

func (a *Account) DepositMoney(amount int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.balance += amount
}

func (a *Account) GetBalance() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.balance
}

func NewAccount(id uint64, owner Person, balance int) *Account {
	return &Account{
		id:      123213123123,
		owner:   owner,
		balance: balance,
	}
}

func WithdrawMoney(waitGroup *sync.WaitGroup, account *Account) {
	for i := 0; i < count; i++ {
		account.WithdrawMoney(i)
		time.Sleep(time.Microsecond)
	}
	waitGroup.Done()
}

func DepositMoney(waitGroup *sync.WaitGroup, account *Account) {
	for i := 0; i < count; i++ {
		account.DepositMoney(i)
		time.Sleep(time.Microsecond)
	}
	waitGroup.Done()
}

func PrintBalance(number int, account *Account) {
	for {
		fmt.Printf("%d)Current account balance = %d\n",
			number, account.GetBalance())
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	var waitGroup sync.WaitGroup
	account := NewAccount(
		324234234,
		Person{
			ID:   233424,
			Name: "Alex",
			Age:  21,
		},
		30000,
	)
	waitGroup.Add(2)
	go WithdrawMoney(&waitGroup, account)
	go DepositMoney(&waitGroup, account)
	for i := 0; i <= 3; i++ {
		go PrintBalance(i, account)
		time.Sleep(50 * time.Millisecond)
	}
	waitGroup.Wait()
	fmt.Printf("Final   account balance = %d", account.GetBalance())
}
