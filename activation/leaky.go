package activation

func LeakyReLU(x float64) float64 {
	if x >= 0.0 {
		return x
	} else {
		return 0.01 * x
	}
}