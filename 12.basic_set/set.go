package main
type void struct{}
type set map[string]void

func main(){
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	s := make(set)
	for _,v := range slice{
		s[v] = void{}
	}
}