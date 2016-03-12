package twod

import "math"

type Vector []float64

func NewVector(x, y float64) Vector {
	return []float64{x, y}
}

func (v Vector) Len() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}

func (v Vector) Sub(v2 Vector) Vector {
	return NewVector(v[0]-v2[0], v[1]-v2[1])
}

func (v Vector) Dot(v2 Vector) float64 {
	return v[0]*v2[0] + v[1]*v2[1]
}

func (v Vector) Atan2() float64 {
	return math.Atan2(v[1], v[0])
}
