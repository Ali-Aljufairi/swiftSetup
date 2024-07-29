package main

import (
    "fmt"
    "os"

    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

type item struct {
    title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
    list     list.Model
    choice   string
    quitting bool
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch keypress := msg.String(); keypress {
        case "ctrl+c":
            m.quitting = true
            return m, tea.Quit
        case "enter":
            i, ok := m.list.SelectedItem().(item)
            if ok {
                m.choice = i.title
            }
            return m, tea.Quit
        }
    case tea.WindowSizeMsg:
        h, v := docStyle.GetFrameSize()
        m.list.SetSize(msg.Width-h, msg.Height-v)
    }

    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
    return m, cmd
}

func (m model) View() string {
    if m.quitting {
        return "Bye!"
    }
    if m.choice != "" {
        return fmt.Sprintf("You chose %s\n", m.choice)
    }
    return docStyle.Render(m.list.View())
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func main() {
    items := []list.Item{
        item{title: "dev", desc: "Install development tools"},
        item{title: "data", desc: "Install data science tools"},
    }

    m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
    m.list.Title = "Server Setup Profiles"

    p := tea.NewProgram(m, tea.WithAltScreen())

    finalModel, err := p.Run()
    if err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }

    if m, ok := finalModel.(model); ok && m.choice != "" {
        fmt.Printf("Installing profile: %s\n", m.choice)
        err := installProfile(m.choice)
        if err != nil {
            fmt.Printf("Error installing profile: %v\n", err)
            os.Exit(1)
        }
        fmt.Println("Profile installed successfully!")
    }
}