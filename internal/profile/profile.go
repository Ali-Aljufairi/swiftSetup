package profile

import (
    "fmt"
    "github.com/Ali-Aljufairi/swiftSetup/pkg/tool"
)

func InstallProfile(profile string) error {
    switch profile {
    case "dev":
        return installTools([]string{"zsh", "curl"})
    case "data":
        return installTools([]string{"zsh", "python3", "jupyter"})
    default:
        return fmt.Errorf("unknown profile: %s", profile)
    }
}

func installTools(toolNames []string) error {
    for _, name := range toolNames {
        for _, t := range tool.AvailableTools {
            if t.Name == name {
                if err := tool.RunCommand(t.InstallCmd); err != nil {
                    return err
                }
                break
            }
        }
    }
    return nil
}