package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    words, _ := reader.ReadString('\n')
	fmt.Print("\n")
	s := strings.Split(words," ")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
	fmt.Println(strings.Join(s," "))
}