package main

import (
	"fmt"
	"math/rand"
)

const ascii_letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func createHugeString(n int) string{
	b := make([]byte, n)
	for i := range b {
		b[i] = ascii_letters[rand.Intn(len(ascii_letters))]
	}
	return string(b)
}

func main() {
	fmt.Println(createHugeString(1 << 10)[:100])
}