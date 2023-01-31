package main
import "fmt"

type Linux struct{}

func (w *Linux) runProgramOnLinux() {
    fmt.Println("Program is running on Linux")
}