package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type safemap struct{
	mx sync.Mutex
	storage map[int]int
}

func main(){
	m := NewMap()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			m.Insert(rand.Intn(10),rand.Intn(100))
		}()
	}
	wg.Wait()
	for i := 0; i < 10;i++{
		fmt.Printf("a record with index %v is %v \n", i, m.Get(i))
	}
}

func NewMap()safemap{
	m := make(map[int]int, 10)
	return safemap{storage: m}
}

func(m *safemap)Insert(where, what int){
	m.mx.Lock()
	m.storage[where] = what
	m.mx.Unlock()
}

func(m *safemap)Get(where int)(what int){
	m.mx.Lock()
	what = m.storage[where]
	m.mx.Unlock()
	return
}