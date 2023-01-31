package main

import(
	"fmt"
	"strings"
)

type void struct{}

func main(){
	var word string
	set := make(map[byte]void)
	fmt.Print("Your word is: ")
	_, err := fmt.Scanln(&word)
	if err != nil{
		panic(err)
	}
	w := strings.ToLower(word)
	for _, l := range []byte(w){
		_,ok :=set[l]
		if ok {
			fmt.Printf("There`s more than one letter %c", l)
			return
		}else{
			set[l] = void{}
		}
	}
	fmt.Printf("All letters are unique in string %s",word)
	



}