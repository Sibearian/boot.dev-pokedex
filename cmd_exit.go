package main

import "os"

func exitCommand(c *config) error {
    defer os.Exit(0)
    return nil
}
