package main

func DisplayBuyersBlock(m model) string {
	s := ""
	s += buying_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(upgrade_title)
	s += buying_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(prestige_title)
	return (s)
}
