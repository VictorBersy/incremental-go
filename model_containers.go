package main

func addContainer(m *model, amount float64) model {
	m.containersGenerated = m.containersGenerated + amount
	m.containers = m.containers + amount
	return *m
}

func buyContainer(m *model) model {
	addContainer(m, m.config.generators["containers_per_tick"]*float64(m.containerGenerator))
	return *m
}

func buyContainerGenerator(m *model) model {
	if m.containers >= float64(m.config.costs["generators"]["containers"]) {
		m.containers = m.containers - float64(m.config.costs["generators"]["containers"])
		m.containerGenerator++
	}
	return *m
}
