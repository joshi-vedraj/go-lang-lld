package payment

import "fmt"

type UPI struct {
	UpiId string
}

func (u *UPI) Pay(amount float64) error {
	fmt.Printf("%0.2f Rs is debited using UPI (UPI ID : %v)\n", amount, u.UpiId)
	return nil
}
