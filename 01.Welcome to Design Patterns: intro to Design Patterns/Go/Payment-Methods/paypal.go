package main

import "fmt"
type Paypal struct {
	email string
	password string
	signedIn bool
	db map[string]string
}

func initPaypal() *Paypal{
	db  := make(map[string]string)
	return &Paypal{
		email: nil,
		password: nil,
		signedIn: false,
		db: db,
	}
}

func (p *Paypal) amount(a int) {
	fmt.Println("Payting amout %d", a)
}

func (p* Paypal) collectPaymentDetails()  {
	fmt.Println("Collecting Payment Details")
}