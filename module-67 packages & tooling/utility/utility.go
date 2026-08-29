package utility

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) float64 {
	if a >= b {
		return float64(a) / float64(b)
	} else {
		return float64(b) / float64(a)
	}
}
