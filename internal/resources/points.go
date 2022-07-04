package resources

import "math"

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
	return Points.Add(p, CalculatePerTick(float64(1)))
}

func (p Points) Generate() Points {
	return Points.Add(p, p.AddPerTick()*float64(p.Generators))
}

func (p Points) BuyGenerator() Points {
	cost := p.GeneratorCost()
	if p.Count >= cost {
		p.Count = p.Count - cost
		p.Generators++
	}
	return p
}

func (p Points) AddPerTick() float64 {
	return CalculatePerTick(float64(0.5) * float64(p.Generators))
}

func (p Points) GeneratorCost() float64 {
	return math.Max(10.0,
		math.Pow(
			math.Max(1, float64(p.Generators)*10),
			float64(p.Generators),
		),
	)
}
