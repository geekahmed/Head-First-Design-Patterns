package main

type IPay interface {
	amount(a int)
	collectPaymentDetails()
}