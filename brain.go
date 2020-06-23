package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Brain struct {
	Input         int
	Hidden        int
	Output        int
	HiddenWeights *mat.Dense
}

func main() {
	fmt.Println("a")
	data := mat.NewDense(3, 4, RandomArray(3*4, float64(3)))
	fmt.Println(data)
}
