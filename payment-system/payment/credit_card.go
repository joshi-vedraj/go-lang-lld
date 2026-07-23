package payment

import "fmt"

type CreditCard struct {
	Number string
	Name   string
}

func (cc *CreditCard) Pay(amount float64) error {
	fmt.Printf("%0.2f Rs is debited from Credit Card (Name : %v, Card Number : %v)\n", amount, cc.Name, cc.Number)
	return nil
}
