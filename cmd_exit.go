package main

import "os"

func exitCommand(c *config, p []string) error {
    defer os.Exit(0)
    return nil
}
