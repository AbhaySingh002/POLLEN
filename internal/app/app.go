package app

import (
    "fmt"

	"github.com/AbhaySingh002/Pollen/internal/app/component"
    "github.com/charmbracelet/bubbles/list"
    // "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)



type model struct {
	// Chat		            textinput.Model
	homelist     			list.Model
	choice   				string
	mainMenuListVisible 	bool
	isHome					bool

}

// InitialModel creates the starting state for the application.
func InitialModel() model {
    // ti := textinput.New()
    // ti.Placeholder = "Type here..."
    // ti.Focus()
    // ti.CharLimit = 200
    // ti.Width = 60 


	

	return model{
		// Chat:           ti,
		homelist: 		component.MainMenu(),
		choice: 		"",
		isHome: 		true,
		mainMenuListVisible: false,
	}
}


func (m model) Init() tea.Cmd {
	return nil
}

// Update handles incoming messages and updates the model's state.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmds []tea.Cmd 

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case tea.KeyShiftTab.String():
			m.isHome = false
			m.mainMenuListVisible = true
			return m, nil

		case tea.KeyEsc.String():
			m.isHome = true
			m.mainMenuListVisible = false
			return m, nil
		case "enter":
			if m.mainMenuListVisible {
				i, ok := m.homelist.SelectedItem().(component.Item)
				if ok {
					m.choice = string(i)
				}
			return m, nil
			}
		}
	}
    // Update input
    // var cmdInput tea.Cmd
    // m.Chat, cmdInput = m.Chat.Update(msg)
    // cmds = append(cmds, cmdInput)

    // Update list
    var cmdList tea.Cmd
    m.homelist, cmdList = m.homelist.Update(msg)
    cmds = append(cmds, cmdList)

    return m, tea.Batch(cmds...)
}




func (m model) View() string {
    header := TitleStyle.Render("Welcome to the POLLEN 🧠")
    var hero string
	var listView string
	var help string
	if m.isHome{
		hero = AsciiStyle.Render(HomeArt)
		listView = ""
		help = HelpStyle.Render(HomeHelp)
	}
	if m.mainMenuListVisible {
		hero = AsciiStyle.Render(Art)
		listView = m.homelist.View()
		listView = MenuFrame.Render(listView)
		help = HelpStyle.Render(MainMenuHelp)
	}

    // List

    

    // Input directly below the list
    // inputBox := lipgloss.NewStyle().
    //     Border(lipgloss.RoundedBorder()).
    //     Padding(0, 1).
    //     Margin(1, 0).
    //     Render(m.Chat.View())

    // Help
	

   	

    return fmt.Sprintf("\n%s\n\n%s\n\n%s\n\n%s",
        header,
        hero,
        listView,
        // inputBox,
        help,
    )
}
