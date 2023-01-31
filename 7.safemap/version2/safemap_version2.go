package main

import (
	"fmt"
	"math/rand"
	"sync"
)


type safemap map[int]int
type insertion struct{
	where int
	what int
}

var wg sync.WaitGroup

func main(){
	storage := make(safemap,10)
	ch := make(chan insertion,10)
	for i:=0;i<10;i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			ch <- insertion{where: rand.Intn(10),what: rand.Intn(1000)}
		}()
	}
	go func()  {
		wg.Wait()
		close(ch)
	}()
	for v := range ch{
		storage[v.where] = v.what
	}
	for i:=0;i<10;i++{
		fmt.Printf("idx №%v is set to %v \n",i,storage[i])
	}
}