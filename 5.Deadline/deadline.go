package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	fmt.Print("Seconds until deadline: ")
	var n int
	wg := sync.WaitGroup{}
	fmt.Scanln(&n)
	ch := make(chan int)
	wg.Add(1)
	go func(ch chan int){
		defer wg.Done()
		for v:=range ch{
			fmt.Printf("Gopher: \"Got the value %v\"\n",v)
		}
		fmt.Println("Gopher: \"I`ve finished\"")
	}(ch)
	t :=  time.After(time.Second*time.Duration(n))
	for{
		select{
			case <-t:
				fmt.Println("Curator: \"Deadline is over\"")
				close(ch)
				wg.Wait()
				return
			default:
				ch<-rand.Intn(1000)
		}
	}


}