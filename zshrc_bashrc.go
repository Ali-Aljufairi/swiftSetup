package main

import (
    "fmt"
    "os"
    "os/user"
    "path/filepath"
)

func configureShell(profile string) error {
    usr, err := user.Current()
    if err != nil {
        return err
    }
    homeDir := usr.HomeDir

    zshrcPath := filepath.Join(homeDir, ".zshrc")
    bashrcPath := filepath.Join(homeDir, ".bashrc")

    if _, err := os.Stat(zshrcPath); err == nil {
        return appendToFile(zshrcPath, profile)
    } else if _, err := os.Stat(bashrcPath); err == nil {
        return appendToFile(bashrcPath, profile)
    }

    return fmt.Errorf("No .zshrc or .bashrc found")
}

func appendToFile(path, profile string) error {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    if _, err := f.WriteString(fmt.Sprintf("\n# Added by CLI tool\nsource %s\n", profile)); err != nil {
        return err
    }

    return nil
}