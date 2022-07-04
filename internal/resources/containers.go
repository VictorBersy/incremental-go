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

func (c Containers) Buy(r Resources) Resources {
	if r.Points.Count > c.cost() {
		r.Points.Count = r.Points.Count - c.cost()
		r.Containers = Containers.add(c, c.cost())
	}
	return r
}

func (c Containers) BuyGenerator() Containers {
	cost := c.generatorCost()
	if c.Count >= cost {
		c.Count = c.Count - cost
		c.Generators++
	}
	return c
}

func (c Containers) cost() float64 {
	return float64(10)
}

func (c Containers) generatorCost() float64 {
	return float64(25)
}
