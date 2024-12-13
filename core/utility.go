package core

func Lerp[T float64 | int](a, b, t T) T {
	return a + t*(b-a)
}

func Clamp[T float64 | int](v, min, max T) T {
	if v < min {
		return min
	} else if v > max {
		return max
	} else {
		return v
	}
}
