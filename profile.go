package main

import (
    "os"
    "os/exec"
)

func installProfile(profile string) error {
    switch profile {
    case "dev":
        return installDevTools()
    case "data":
        return installDataTools()
    default:
        return nil
    }
}

func installDevTools() error {
    cmds := []string{
        "sudo apt-get install git",
    }
    return runCommands(cmds)
}

func installDataTools() error {
    cmds := []string{
        "sudo apt-get install python3",
        "sudo apt-get install jupyter",
        // Add more commands as needed
    }
    return runCommands(cmds)
}

func runCommands(cmds []string) error {
    for _, cmd := range cmds {
        parts := strings.Fields(cmd)
        head := parts[0]
        parts = parts[1:len(parts)]

        out, err := exec.Command(head, parts...).Output()
        if err != nil {
            return err
        }
        fmt.Println(string(out))
    }
    return nil
}