package payment

import "errors"

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) Pay(amount float64) error {
	if c.strategy == nil {
		return errors.New("Payment Strategy not set")
	}
	return c.strategy.Pay(amount)
}
