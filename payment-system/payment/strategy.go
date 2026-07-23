package payment

type Strategy interface {
	Pay(amount float64) error
}
