package main

func addPoint(m *model, amount float64) model {
	m.pointsGenerated = m.pointsGenerated + amount
	m.points = m.points + amount
	return *m
}

func generatePoint(m *model) model {
	addPoint(m, m.config.generators["points_per_tick"]*float64(m.generator_points))
	return *m
}

func buyPointGenerator(m *model) model {
	if m.points >= float64(m.config.costs["generators"]["points"]) {
		m.points = m.points - float64(m.config.costs["generators"]["points"])
		m.generator_points++
	}
	return *m
}
