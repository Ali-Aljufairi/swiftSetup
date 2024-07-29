package main
import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbles/textinput"
)

type initialModel struct {
    textInput textinput.Model
    choice    string
}

func (m initialModel) Init() tea.Cmd {
    return textinput.Blink
}

func (m initialModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            return m, tea.Quit
        case "enter":
            m.choice = m.textInput.Value()
            return m, tea.Quit
        }
    }

    var cmd tea.Cmd
    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m initialModel) View() string {
    return lipgloss.NewStyle().Render("Enter the profile to install: " + m.textInput.View())
}

func newInitialModel() initialModel {
    ti := textinput.New()
    ti.Placeholder = "Profile name"
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 20
    return initialModel{
        textInput: ti,
    }
}