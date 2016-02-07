package integer

import "log"

func Min(as ...int) int {
	min := as[0]
	for _, a := range as[1:] {
		if a < min {
			min = a
		}
	}
	return min
}

func Max(as ...int) int {
	max := as[0]
	for _, a := range as[1:] {
		if a > max {
			max = a
		}
	}
	return max
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func Ceil(a, b int) int {
	return ((a - 1) / b) + 1
}

func Range(params ...int) []int {
	var start, stop, step int

	switch len(params) {
	case 1:
		start = 0
		stop = params[0]
		step = 1
	case 2:
		start = params[0]
		stop = params[1]
		step = 1
	case 3:
		start = params[0]
		stop = params[1]
		step = params[2]
	default:
		log.Fatalln("Invalid params for range:", params)
	}

	slice := []int{}
	for i := start; i < stop; i += step {
		slice = append(slice, i)
	}
	return slice
}
