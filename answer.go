package answer

func q0001(n int) int {
	var sum int = 0

	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func q0002(max int) int {
	var (
		m      int = 1
		n      int = 1
		answer int = 0
	)

	for {
		if n > max {
			return answer
		}

		if n%2 == 0 {
			answer += n
		}

		n, m = n+m, n
	}
}
