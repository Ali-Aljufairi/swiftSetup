package tool

import (
    "fmt"
    "os/exec"
    "strings"
)

type Tool struct {
    Name        string
    InstallCmd  string
    Description string
}

var AvailableTools = []Tool{
    {Name: "zsh", InstallCmd: "sudo apt-get install -y zsh", Description: "Z shell"},
    {Name: "curl", InstallCmd: "sudo apt-get install -y curl", Description: "Command-line tool for transferring data"},
    {Name: "python3", InstallCmd: "sudo apt-get install -y python3", Description: "Python programming language"},
    {Name: "jupyter", InstallCmd: "sudo apt-get install -y jupyter", Description: "Web-based interactive computational environment"},
    // Add more tools as needed
}

func RunCommand(cmd string) error {
    parts := strings.Fields(cmd)
    head := parts[0]
    parts = parts[1:]

    out, err := exec.Command(head, parts...).Output()
    if err != nil {
        return fmt.Errorf("error executing command '%s': %v", cmd, err)
    }
    fmt.Println(string(out))
    return nil
}

func AddNewTool(name, installCmd, description string) {
    AvailableTools = append(AvailableTools, Tool{
        Name:        name,
        InstallCmd:  installCmd,
        Description: description,
    })
}