package main
import (
	"math/rand"
	"fmt"
)

func main(){
	slice := []int{1,2,3,4,5,6,7}
	i := rand.Intn(len(slice))
	slice = remove(slice,i)
	fmt.Println(slice)
	
}

func remove(slice []int, idx int)[]int{
	return append(slice[:idx],slice[idx+1:]...)
}