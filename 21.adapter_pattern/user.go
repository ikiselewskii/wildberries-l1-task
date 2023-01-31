package main

import "fmt"

type User struct {
}

func (u *User) RunProgram(p Program) {
    fmt.Println("User tries to run a program on his OS.")
    p.RunProgramOnWindows()
}
