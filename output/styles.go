package output

import "charm.land/lipgloss/v2"

const (
	contentWidth = 80
	sidePadding  = 4
)

var GlobalGameStyle = lipgloss.NewStyle().
	Width(contentWidth).
	Padding(0, sidePadding)
