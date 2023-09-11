package twocrystalballs

import "math"

func TwoCrystalBalls(b []bool) int {
	stepSize := int(math.Sqrt(float64(len(b))))
	i := 0
	for !b[i] {
		i += stepSize
	}

	for j := i - stepSize; j < i; j++ {
		if b[j] {
			return j
		}
	}
	return -1
}
