package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func installProfile(profile string) error {
	switch profile {
	case "dev":
		return installDevTools()
	case "data":
		return installDataTools()
	default:
		return fmt.Errorf("unknown profile: %s", profile)
	}
}

func installDevTools() error {
	cmds := []string{
		"sudo apt-get install jq",
		// Add more commands as needed
	}
	return runCommands(cmds)
}

func installDataTools() error {
	cmds := []string{
		"sudo apt-get install -y python3",
		"sudo apt-get install -y jupyter",
		// Add more commands as needed
	}
	return runCommands(cmds)
}

func runCommands(cmds []string) error {
	for _, cmd := range cmds {
		parts := strings.Fields(cmd)
		head := parts[0]
		parts = parts[1:]

		out, err := exec.Command(head, parts...).Output()
		if err != nil {
			return fmt.Errorf("error executing command '%s': %v", cmd, err)
		}
		fmt.Println(string(out))
	}
	return nil
}
