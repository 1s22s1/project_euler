package answer

func q0001(n int) int {
	var sum int = 0

	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += n
		}
	}

	return sum
}
