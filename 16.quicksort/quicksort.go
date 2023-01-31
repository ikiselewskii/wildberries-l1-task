package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := generateSlice(20)
	fmt.Println("Unsorted: ", slice)
	quicksort(slice)
	fmt.Println("Sorted: ", slice)
}

func generateSlice(size int) []int {

	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
 
func quicksort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
     
    left, right := 0, len(arr)-1
     
    pivot := rand.Int() % len(arr)
     
    arr[pivot], arr[right] = arr[right], arr[pivot]
     
    for i := range arr {
        if arr[i] < arr[right] {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }
     
    arr[left], arr[right] = arr[right], arr[left]
     
    quicksort(arr[:left])
    quicksort(arr[left+1:])
     
    return arr
}