package main

import (
	"fmt"
	"math/rand"
)

func main(){
	var num int64 = rand.Int63()
	i := 20
	fmt.Printf("number is:%d base2:%b\n",num,num)
	fmt.Println("Alakazam!")
	mask := 1 << (i-1)
	fmt.Printf("number is:%d base2:%b\n",num^int64(mask),num^int64(mask))
}