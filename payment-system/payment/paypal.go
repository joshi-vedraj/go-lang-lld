package payment

import "fmt"

type Paypal struct {
	Email string
}

func (p *Paypal) Pay(amount float64) error {
	fmt.Printf("%0.2f Rs is debited using Paypal (Email : %v)\n", amount, p.Email)
	return nil
}
