package resources

type Containers struct {
	Count      float64
	Generated  float64
	Boost      float64
	Generators int
}

func (c Containers) add(amount float64) Containers {
	c.Generated = c.Generated + amount
	c.Count = c.Count + amount
	return c
}

func (c Containers) Buy() Containers {
	Containers.add(c, float64(1))
	return c
}

func (c Containers) BuyGenerator() Containers {
	if c.Count >= float64(1) {
		c.Count = c.Count - float64(1)
		c.Generators++
	}
	return c
}
