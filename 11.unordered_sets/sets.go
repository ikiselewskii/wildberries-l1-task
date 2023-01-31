package main

import "fmt"

type void struct{}
type set map[string]void

func main(){
	pets := set{"cat":void{}, "dog":void{}, "snake":void{}, "turtle":void{}}
	mammals := set{"cat":void{}, "dog":void{}, "lion":void{}, "tiger":void{}}
	interception := intercept(pets,mammals)
	for k := range interception{
		fmt.Println(k)
	}
}

func intercept(set1,set2 set)set{
	ans := make(set)
	for k := range set1{
		_,ok := set2[k]
		if ok{
			ans[k] = void{}
		}
	}
	return ans
}