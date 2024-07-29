package shell

import (
    "fmt"
    "os"
    "os/user"
    "path/filepath"
)

func ConfigureShell(profile string, aliases map[string]string) error {
    usr, err := user.Current()
    if err != nil {
        return err
    }
    homeDir := usr.HomeDir

    zshrcPath := filepath.Join(homeDir, ".zshrc")
    bashrcPath := filepath.Join(homeDir, ".bashrc")

    var rcFile string
    if _, err := os.Stat(zshrcPath); err == nil {
        rcFile = zshrcPath
    } else if _, err := os.Stat(bashrcPath); err == nil {
        rcFile = bashrcPath
    } else {
        return fmt.Errorf("No .zshrc or .bashrc found")
    }

    return appendToFile(rcFile, profile, aliases)
}

func appendToFile(path, profile string, aliases map[string]string) error {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    if _, err := f.WriteString(fmt.Sprintf("\n# Added by CLI tool\nsource %s\n", profile)); err != nil {
        return err
    }

    for alias, command := range aliases {
        if _, err := f.WriteString(fmt.Sprintf("alias %s='%s'\n", alias, command)); err != nil {
            return err
        }
    }

    return nil
}