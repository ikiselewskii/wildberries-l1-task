package main

import "fmt"

type Human struct {
	Age  uint
	Name string
	Job  string
}

func (h *Human) SayHello() {
	fmt.Printf("Hello my name is %v,\n I`m %d years old,\n I`ve got a job as a %v \n", h.Name, h.Age, h.Job)
}

type Action struct {
	Human
	Duration int
	Task     string
}

func main() {
	h := Human{22, "Ilya", "Python Developer"}
	h.SayHello()
	a := Action{h, 40, "Writing code"}
	a.SayHello()
}
