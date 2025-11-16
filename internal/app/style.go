package app

import (
	"github.com/charmbracelet/lipgloss"
)


var (
	AsciiStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fbfe3fff")).
		Bold(true).
		Align(lipgloss.Center).
		Width(80)

	TitleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#050003ff")). 
		Background(lipgloss.Color("#f1af89ff")). 
		Margin(0, 2).
		Padding(0, 1)

	HelpStyle = lipgloss.NewStyle().PaddingLeft(4).Align(lipgloss.Center).Foreground(lipgloss.Color("#8a8a8aff"))

	MenuFrame = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 1)

)