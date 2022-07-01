package resources

type Points struct {
	Count      float64
	Generated  float64
	Boost      float64
	Generators int
}

func (p Points) Add(amount float64) Points {
	p.Generated = p.Generated + amount
	p.Count = p.Count + amount
	return p
}

func (p Points) Generate() Points {
	Points.Add(p, float64(1))
	return p
}

func (p Points) BuyGenerator() Points {
	if p.Count >= float64(1) {
		p.Count = p.Count - float64(1)
		p.Generators++
	}
	return p
}
