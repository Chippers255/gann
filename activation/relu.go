package activation

func ReLU(r int, c int, x float64) float64 {
	if x > 0.0 {
		return x
	} else {
		return 0.0
	}
}
