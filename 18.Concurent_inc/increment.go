package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SafeCounter struct{ //создаем структуру с мютексом
	mu sync.Mutex
	counter int
}
var wg sync.WaitGroup
func main(){
	rand.Seed(time.Now().UnixNano())	//обновляем сид 
	c := NewCounter()
	for i := 0;i<rand.Intn(100);i++{
		wg.Add(1) //+1 в счетчик ожидание горутин
		go func(){
			defer wg.Done() //отнимаем в счетчике 
			c.Inc()
		}() //вызываем горутину
	}
	wg.Wait() //ждём завершения всех горутин
	fmt.Println(c.counter)
}

func NewCounter()*SafeCounter{ //конструируем новый каунтер
	c := new(SafeCounter)
	c.counter = 0
	return c
}


func(c *SafeCounter)Inc(){
	c.mu.Lock()		//замыкаем мютекс или ждём пока мютекс откроется
	c.counter++
	c.mu.Unlock()
}