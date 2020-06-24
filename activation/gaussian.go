package activation

import "math"

func Gaussian(r int, c int, x float64) float64 {
	return math.Exp(math.Pow(-x, 2))
}
