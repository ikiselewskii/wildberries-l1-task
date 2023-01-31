package main

import "fmt"

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Printf("index of a value we`re looking for is %v\n",binsearch(63, items))
}

func binsearch(what int, where []int) int {
	var low int = 0
	var high int = len(where) - 1
  
	for low <= high {
	  var mid int = low + (high-low)/2
	  var midVal int = where[mid]
	  fmt.Println("middle value is:", midVal)
  
	  if midVal == what {
		return mid
	  } else if midVal > what {
		high = mid - 1
	  } else {
		low = mid + 1
	  }
	}
	return -1
  }