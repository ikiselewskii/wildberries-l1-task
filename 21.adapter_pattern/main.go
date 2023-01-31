package main

func main() {

    user := &User{}
    win := &Windows{}

    user.RunProgram(win)

    archLinux := &Linux{}
    windowsMachineAdapter := &Wine{
        linuxOS: archLinux,
    }

    user.RunProgram(windowsMachineAdapter)
}