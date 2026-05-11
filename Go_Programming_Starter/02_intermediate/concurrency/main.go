package main

import "fmt"

type Wallet struct {
	owner   string
	balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.balance += amount
}
func (w *Wallet) SpendMoney(amount float64) {
	if w.balance < amount {
		fmt.Printf("Insufficient balance you poor guy")
	} else {
		w.balance -= amount
	}
}
func main() {
	var wallet Wallet

	wallet.AddMoney(200.0)
	wallet.SpendMoney(100)

	fmt.Println(wallet)
}
