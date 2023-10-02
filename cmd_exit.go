package main

import "os"

func exitCommand() error {
    defer os.Exit(0)
    return nil
}
