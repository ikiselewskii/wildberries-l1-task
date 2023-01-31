package main

import "fmt"

type Wine struct {
	linuxOS *Linux
}

func (w *Wine) RunProgramOnWindows() {
	fmt.Println("Wine runs Windows program on Linux.")
	w.linuxOS.runProgramOnLinux()
}
