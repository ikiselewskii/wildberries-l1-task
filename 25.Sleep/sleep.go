package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Print("Wait for (seconds): ")
	var n uint
	fmt.Scanln(&n)
	sleep(n)
}

func sleep(seconds uint){
	t := time.After(time.Second*time.Duration(seconds))
	<-t
}