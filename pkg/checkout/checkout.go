package checkout

type Checkout struct {
	count int
	price int
}

func (c *Checkout) add(count int, price int) {
	c.count = count
	c.price = price
}
