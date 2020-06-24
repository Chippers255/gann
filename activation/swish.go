package activation

import "math"

func Swish(x float64) float64 {
	return x / (1.0 + math.Exp(-x))
}
