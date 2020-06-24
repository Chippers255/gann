package main

import (
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

// Generate an array of random floating point values with a user define size.
func RandomArray(size int, v float64) []float64 {
	dist := distuv.Uniform{
		Min: -1 / math.Sqrt(v),
		Max: 1 / math.Sqrt(v),
	}

	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = dist.Rand()
	}

	return data
}

// Saves a few lines of code by making dot product of two matrices easier to execute.
func Dot(m mat.Matrix, n mat.Matrix) mat.Matrix {
	// We need to collect the size of the input matrices for the resulting matrix
	r, _ := m.Dims()
	_, c := n.Dims()

	// Calculate the dot product and return the result as a new matrix
	result := mat.NewDense(r, c, nil)
	result.Product(m, n)

	return result
}

// Saves lines of code again by allowing a function to be applied to all values in a matrix.
func Apply(fn func(i int, j int, v float64) float64, m mat.Matrix) mat.Matrix {
	r, c := m.Dims()

	result := mat.NewDense(r, c, nil)
	result.Apply(fn, m)

	return result
}
