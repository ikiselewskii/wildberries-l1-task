package main
import (
	"fmt"
	"reflect"
)

func main(){
	i := 12
	generic(i)

	s := "hello"
	generic(s)

	ok := "true"
	generic(ok)

	ch := make(chan int)
	generic(ch)

	var x struct{}
	generic(x)

}


func generic(obj interface{}){
	t := reflect.ValueOf(obj)
	switch t.Kind(){
		case reflect.Int:
			fmt.Printf("%v is an integer \n",t)
		case reflect.String:
			fmt.Printf("%v is a string \n",t)
		case reflect.Bool:
			fmt.Printf("%v is a boolean value \n",t)
		case reflect.Chan:
			fmt.Printf("%v is a channel \n",t)
		default:
			fmt.Printf("I don`t know what type is %v \n",t)
	}
	
}