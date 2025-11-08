package app

import (
	"fmt"
	// "os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	orch "github.com/AbhaySingh002/Pollen/internal/orchestrator"
	// Import your AI package if you need to trigger AI functions from here
	// For now, IntentChecking is called in main.go as a goroutine.
	// If you need to display results or trigger AI based on user input,
	// you'll need to import and use it here.
	// "github.com/AbhaySingh002/Pollen/internal/ai"
)

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	// aiResponse             string
	// aiError                error
}

// InitialModel creates the starting state for the application.
func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Type the Prompt here...."
	ti.Focus()
	ti.CharLimit = 30
	ti.Width = 50 // This width might need adjustment based on terminal size or lipgloss styling

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		// aiResponse:             "",
		// aiError:                nil,
	}
}

// Init is called once when the program starts.
func (m model) Init() tea.Cmd {
	// If you wanted to trigger an AI process *after* the UI loads,
	// you would return it here as a command.
	// For example:
	// return ai.IntentChecking() // This would require IntentChecking to return a tea.Cmd
	return nil // No I/O operations to perform at startup
}

// Update handles incoming messages and updates the model's state.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd // To collect multiple commands

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			// Quitting the application
			return m, tea.Quit
		case "ctrl+p":
			// Toggle the visibility of the input field
			m.createFileInputVisible = !m.createFileInputVisible
			// If we are opening the input, focus it
			if m.createFileInputVisible {
				m.newFileInput.Focus()
			} else {
				m.newFileInput.Blur() // Blur if hiding it
			}
			return m, nil
		case "enter":
			// Handle what happens when Enter is pressed on the text input
			if m.createFileInputVisible {
				// You could trigger the AI call here based on the input
				// For example:
				prompt := m.newFileInput.Value()
				go orch.Orchestrator(prompt) // Implement this function
				// m.createFileInputVisible = false // Hide input after submitting
				// m.newFileInput.SetValue("") // Clear input
				// return m, nil // Return with updated model state

				// For now, just hide it and clear the value
				// m.createFileInputVisible = false
				m.newFileInput.SetValue("")
				// m.newFileInput.Blur()
				return m, nil
			}
		}
		// Add cases here to handle messages from your AI, e.g.:
		// case ai.AIParsedResultMsg: // Assuming you define such a message type
		// 	m.aiResponse = msg.Result
		// 	return m, nil
		// case ai.AIErrorMsg:
		// 	m.aiError = msg.Err
		// 	return m, nil
	}

	// If the input field is visible, process its messages
	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...) // Return the model and any collected commands
}

// View renders the current state of the model to the terminal.
func (m model) View() string {
	// Styles
	var titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#050003ff")). // Dark foreground
		Background(lipgloss.Color("#ec60d7ff")). // Pinkish background
		Margin(0, 2).
		Padding(0, 1)

	var asciiStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#139379ff")). // Teal-like color
		Bold(true).
		Align(lipgloss.Center).
		Width(80) // Fixed width for the ASCII art

	// Content
	welcome := titleStyle.Render("Welcome to the POLLEN 🧠")
	asciiArt := `██████╗  ██████╗ ██╗     ██╗     ███████╗███╗   ██╗
██╔══██╗██╔═══██╗██║     ██║     ██╔════╝████╗  ██║
██████╔╝██║   ██║██║     ██║     █████╗  ██╔██╗ ██║
██╔═══╝ ██║   ██║██║     ██║     ██╔══╝  ██║╚██╗██║
██║     ╚██████╔╝███████╗███████╗███████╗██║ ╚████║
╚═╝      ╚═════╝ ╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝
												`

	help := "Ctrl+P: Send Prompt · Ctrl+C: Quit " // Simplified help string

	// Render input field if visible
	inputView := ""
	if m.createFileInputVisible {
		inputView = m.newFileInput.View()
	}

	// Combine all parts
	return fmt.Sprintf("\n%s\n\n%s\n\n%s\n\n%s",
		welcome,
		asciiStyle.Render(asciiArt),
		inputView, // Render the input view or empty string
		lipgloss.NewStyle().Align(lipgloss.Center).Render(help), // Center the help text
	)
}

// Note: The original `ai.IntentChecking()` call is now in `main.go`.
// If you need to trigger AI functionality based on user input (e.g., after pressing Enter
// in the text input), you would add a command here in the `Update` function that calls
// into the `ai` package and potentially returns a custom message type to update the UI.
