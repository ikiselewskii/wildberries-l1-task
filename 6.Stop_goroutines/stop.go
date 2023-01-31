package main

import (
	"sync"
	"fmt"
	"time"
)

var wg sync.WaitGroup
type void struct{}

func main(){
	quit := make(chan void)
	wg.Add(1)
	go func(){
		defer wg.Done()
		for{
		select{
			case <-quit:
				fmt.Println("Gopher: \"I`m done\"")
				return
			default:
				fmt.Println("Gopher: \"I`m working\"")
			}
		}
	}()
	time.Sleep(time.Second)
	quit <- void{}
	wg.Wait()
	ch := make(chan void)
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for range ch{
			fmt.Println("Gopher: \"I`m working\"")
		}
		fmt.Println("Gopher: \"I`m done\"")
	}()
	
	for i:=0;i<10;i++{
		ch<-void{}
	}

}