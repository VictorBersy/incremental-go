package main

func DisplayBuyersBlock(m model) string {
	s := ""
	s += block_buyers_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(upgrade_title)
	s += block_buyers_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(prestige_title)
	return (s)
}
