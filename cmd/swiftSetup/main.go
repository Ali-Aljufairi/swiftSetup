package main

import (
    "fmt"
    "os"

    "github.com/Ali-Aljufairi/swiftSetup/internal/model"
    "github.com/Ali-Aljufairi/swiftSetup/internal/profile"
    "github.com/Ali-Aljufairi/swiftSetup/internal/shell"
    "github.com/Ali-Aljufairi/swiftSetup/pkg/tool"

    "github.com/charmbracelet/bubbles/list"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func main() {
    items := []list.Item{
        model.NewItem("dev", "Install development tools"),
        model.NewItem("data", "Install data science tools"),
        model.NewItem("add_tool", "Add a new tool"),
    }

    m := model.Model{List: list.New(items, list.NewDefaultDelegate(), 0, 0)}
    m.List.Title = "Server Setup Profiles"

    p := tea.NewProgram(m, tea.WithAltScreen())

    finalModel, err := p.Run()
    if err != nil {
        fmt.Printf("Error: %v", err)
        os.Exit(1)
    }

    if m, ok := finalModel.(model.Model); ok && m.Choice != "" {
        if m.Choice == "add_tool" {
            addNewToolInteractive()
        } else {
            fmt.Printf("Installing profile: %s\n", m.Choice)
            err := profile.InstallProfile(m.Choice)
            if err != nil {
                fmt.Printf("Error installing profile: %v\n", err)
                os.Exit(1)
            }
            fmt.Println("Profile installed successfully!")

            aliases := map[string]string{
                "update": "sudo apt-get update && sudo apt-get upgrade -y",
                "cls":    "clear",
                // Add more aliases as needed
            }
            err = shell.ConfigureShell(m.Choice, aliases)
            if err != nil {
                fmt.Printf("Error configuring shell: %v\n", err)
                os.Exit(1)
            }
            fmt.Println("Shell configured successfully!")
        }
    }
}

func addNewToolInteractive() {
    var name, installCmd, description string

    fmt.Print("Enter tool name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter install command: ")
    fmt.Scanln(&installCmd)
    fmt.Print("Enter description: ")
    fmt.Scanln(&description)

    tool.AddNewTool(name, installCmd, description)
    fmt.Println("New tool added successfully!")
}