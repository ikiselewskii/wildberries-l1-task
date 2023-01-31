package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main(){
	fmt.Print("Insert a number of workers: ")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil{
		panic(err)
	}
	ch := make(chan int)
	exit := make(chan os.Signal,1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)	//уведомить о нажатии Ctrl+C или о сигнале TERMINATE
	for i := 0; i < n; i++{
		wg.Add(1)
		go worker(ch,i)
	}
	for{
		select{
		case <-exit:	//при получении сигнала из канала exit закрываем канал ch и ждём завершения всех горутин
			close(ch)
			fmt.Println("The channel is closed")
			wg.Wait()
			return
		default:
			ch <- rand.Intn(1000)
		}
	}
}

func worker(ch chan int, name int){
	defer wg.Done()
	for v:= range ch{	//range по каналу берёт значения из канала пока канал не будет закрыт
		fmt.Printf("worker %v: got the value %v\n",name,v)
		time.Sleep(time.Second)
	}
	fmt.Printf("worker %v has finished execution\n",name)
}
