package twod

import "math"

func AngleDiff(a1, a2 float64) float64 {
	ad := a1 - a2
	if ad <= 0 {
		ad += 2 * math.Pi
	}
	return ad
}
