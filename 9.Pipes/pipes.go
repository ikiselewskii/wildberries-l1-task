package main

import "fmt"

func main(){
	slice := []int{1,2,3,4,5,6,7,8,9,10}
	ch1 := make(chan int)
	ch2 := make(chan int)
	finish := make(chan bool)

	go func(){
		for _,v := range slice{
			ch1 <- v
		}
		close(ch1)
	}()

	go func(){
		for v := range ch1{
			ch2 <- v*v
		}
		close(ch2)

	}()

	go func(){
		for v := range ch2{
			fmt.Println(v)
		}
		finish <- true
	}()
	<-finish
	


}