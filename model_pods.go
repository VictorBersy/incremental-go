package main

func buyPods(m *model) model {
	if m.containers >= 100 {
		m.containers = m.containers - 100
		m.pods++
	}
	return *m
}
