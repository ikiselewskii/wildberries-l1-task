package main

import "fmt"

type Windows struct {
}

func (w *Windows) RunProgramOnWindows() {
    fmt.Println("Lightning connector is plugged into mac machine.")
}