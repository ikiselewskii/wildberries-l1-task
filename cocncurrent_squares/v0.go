package main

import "fmt"
import "sync"

var wg sync.WaitGroup

func main(){

	slice := []int{1,2,3,2,2,2,3,5,4}
	for i,v := range slice{
		wg.Add(1)
		go print_square(i,v)
	}
		wg.Wait()
}

func print_square(i,v int){
	defer wg.Done()
	fmt.Printf("the square of element №%d is %d \n",i+1,v*v)
}