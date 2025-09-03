package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

type CreditCard struct{}
type PayPal struct{}

func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using Credit Card\n", amount)
}

func (p PayPal) Pay(amount float64) {
	fmt.Printf("Paid ₹%.2f using PayPal\n", amount)
}

func checkout(p PaymentMethod, amount float64) {
	p.Pay(amount)
}

func main() {
	var cc CreditCard
	var pp PayPal

	checkout(cc, 500)
	checkout(pp, 800)
}
