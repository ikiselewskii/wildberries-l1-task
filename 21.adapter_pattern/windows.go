package main

import "fmt"

type Windows struct {
}

func (w *Windows) RunProgramOnWindows() {
    fmt.Println("Program is running on Windows")
}