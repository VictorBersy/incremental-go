package main

func (m model) View() string {
	s := generateLayout(m)
	if m.quitting {
		s += "\n"
	}
	return s
}
