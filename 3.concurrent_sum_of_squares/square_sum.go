	package main

	import (
		"fmt"
		"sync"
	)

	var wg sync.WaitGroup

	func main(){
		var res int
		ch := make(chan int)
		done := make(chan bool)
		slice := []int{2,4,6,8,10}
		for _,v := range slice{
			wg.Add(1)
			go print_square(v, ch)
		}
		go func(done chan bool){
			wg.Wait()
			done <- true
		}(done)

		for{
			select{
			case sqare := <- ch:
				res += sqare
			case <- done:
				fmt.Printf("the squared sum of given slice is %d", res)
				return
			}
		}
	}

	func print_square(v int, ch chan int){
		defer wg.Done()
		ch <- v*v
		
	}