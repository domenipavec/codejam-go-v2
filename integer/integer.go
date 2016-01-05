package integer

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

func Ceil(a, b int) int {
	return ((a - 1) / b) + 1
}
