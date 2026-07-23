package main

import (
	"fmt"

	"github.com/joshi-vedraj/payment-system/payment"
)

func main() {
	context := payment.Context{}

	card := payment.CreditCard{
		Number: "1234 5678 9012",
		Name:   "Virat Kohli",
	}

	upi := payment.UPI{
		UpiId: "1234567890@abc",
	}

	paypal := payment.Paypal{
		Email: "viratkohli@abc.com",
	}

	err := context.Pay(6000)
	if err != nil {
		fmt.Println(err)
	}

	context.SetStrategy(&card)
	err = context.Pay(12000)
	if err != nil {
		fmt.Println(err)
	}

	context.SetStrategy(&upi)
	err = context.Pay(5000)
	if err != nil {
		fmt.Println(err)
	}

	context.SetStrategy(&paypal)
	err = context.Pay(6000)
	if err != nil {
		fmt.Println(err)
	}

}
