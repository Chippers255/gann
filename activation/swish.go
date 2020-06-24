package activation

import "math"

func Swish(r int, c int, x float64) float64 {
	return x / (1.0 + math.Exp(-x))
}
