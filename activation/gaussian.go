package activation

import "math"

func Gaussian(x float64) float64 {
	return math.Exp(math.Pow(-x, 2))
}