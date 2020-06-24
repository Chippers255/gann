package activation

import "math"

func Sigmoid(r int, c int, x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}
