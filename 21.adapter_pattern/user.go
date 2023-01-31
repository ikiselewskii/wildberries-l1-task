package main

import "fmt"

type User struct {
}

func (u *User) RunProgram(p Program) {
    fmt.Println("Client inserts Lightning connector into computer.")
    p.RunProgramOnWindows()
}
