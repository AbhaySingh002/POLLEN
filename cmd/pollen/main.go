package main

import (
	"fmt"
	"os"


	"github.com/AbhaySingh002/Pollen/internal/app" // Import your app package
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(app.InitialModel()) 
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Alas, there's been an error: %v\n", err)
		os.Exit(1)
	}
}
