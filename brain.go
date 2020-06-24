package main

import (
	"fmt"
	"github.com/gann/activation"
	"gonum.org/v1/gonum/mat"
)

type Brain struct {
	Input         int
	Hidden        int
	Output        int
	HiddenWeights *mat.Dense
	OutputWeights *mat.Dense
}

//func (b Brain) Run(inputs )

func main() {
	fmt.Println("a")
	//data := mat.NewDense(3, 4, RandomArray(3*4, float64(3)))
	IHWeights := mat.NewDense(3, 2, []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0})
	fmt.Println(IHWeights)

	Inputs := mat.NewDense(2, 1, []float64{1.0, 2.0})
	fmt.Println(Inputs)

	Hiddens := mat.NewDense(3, 1, nil)
	Hiddens.Product(IHWeights, Inputs)
	fmt.Println(Hiddens)

	results := Apply(activation.ReLU, Hiddens)
	fmt.Println(results)
}
