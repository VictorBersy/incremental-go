package resources

type Pods struct {
	Count     float64
	Generated float64
}

func (p Pods) add(amount float64) Pods {
	p.Generated = p.Generated + amount
	p.Count = p.Count + amount
	return p
}

func (p Pods) Buy() Pods {
	return Pods.add(p, float64(1))
}
