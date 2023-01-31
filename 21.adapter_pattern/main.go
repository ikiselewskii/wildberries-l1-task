package main

func main() {

    user := &User{}
    win := &Windows{}

    user.RunProgram(win)

    archLinux := &Linux{}
    emulator := &Wine{
        linuxOS: archLinux,
    }

    user.RunProgram(emulator)
}