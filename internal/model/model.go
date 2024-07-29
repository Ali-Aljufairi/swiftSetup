package model

import (
    "fmt"

    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type Item struct {
    title       string
    description string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.description }
func (i Item) FilterValue() string { return i.title }

type Model struct {
    List     list.Model
    Choice   string
    Quitting bool
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch keypress := msg.String(); keypress {
        case "ctrl+c":
            m.Quitting = true
            return m, tea.Quit
        case "enter":
            i, ok := m.List.SelectedItem().(Item)
            if ok {
                m.Choice = i.Title()
            }
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.List.SetSize(msg.Width-h, msg.Height-v)
    }

    var cmd tea.Cmd
    m.List, cmd = m.List.Update(msg)
    return m, cmd
}

func (m Model) View() string {
    if m.Quitting {
        return "Bye!"
    }
    if m.Choice != "" {
        return fmt.Sprintf("You chose %s\n", m.Choice)
    }
    return docStyle.Render(m.List.View())
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

// NewItem creates a new Item
func NewItem(title, description string) Item {
    return Item{
        title:       title,
        description: description,
    }
}