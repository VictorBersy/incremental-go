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

func (p Points) ManuallyGenerate() Points {
	return Points.Add(p, addPerTick())
}

func (p Points) Generate() Points {
	return Points.Add(p, addPerTick()*float64(p.Generators))
}

func (p Points) BuyGenerator() Points {
	cost := generatorCost()
	if p.Count >= cost {
		p.Count = p.Count - cost
		p.Generators++
	}
	return p
}

func addPerTick() float64 {
	return float64(1)
}

func generatorCost() float64 {
	return float64(10)
}
