package main

import (
	"fmt"
	"math/big"
)

func main(){
	x,y := big.NewFloat(1048577),big.NewFloat(1050501)
	add := new(big.Float).Add(x,y)
	divide := new(big.Float).Quo(x,y)
	multiply := new(big.Float).Mul(x,y)
	substract := new(big.Float).Sub(x,y)
	fmt.Printf("results of %1.f and %1.f`s \n addition: %1.f\n division: %1.f\n multiplication: %1.f\n substraction: %1.f\n",
	x,y,add,divide,multiply,substract)
}