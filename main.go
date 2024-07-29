package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "log"
)

func main() {
    p := tea.NewProgram(initialModel{})
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
}