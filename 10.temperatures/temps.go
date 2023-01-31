package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	clustered_temps := make(map[int][]float64)

	for _, v := range temperatures {
		i := int(v) - (int(v) % 10)
		clustered_temps[i] = append(clustered_temps[i], v)
	}
	for k, v := range clustered_temps{
	fmt.Println(k,": ",v)
	}
}