package answer

import "strconv"

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

func q0003(n int) int {
	var x []int

	for {
		if n == 1 {
			break
		}

		for i := 2; i <= n; i++ {
			if n%i == 0 {
				x = append(x, i)
				n = n / i
				break
			}
		}
	}

	max := -1
	for _, x := range x {
		if max < x {
			max = x
		}
	}

	return max
}

func isPalindrome(s string) bool {
	var runes []rune = []rune(s)
	var reverseRunes []rune = []rune(s)

	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		reverseRunes[j] = runes[i]
	}

	return s == string(reverseRunes)
}

func q0004(first int, last int) int {
	var x []int

	for i := first; i <= last; i++ {
		for j := first; j <= last; j++ {
			if isPalindrome(strconv.Itoa(i * j)) {
				x = append(x, i*j)
			}
		}
	}

	max := -1
	for _, x := range x {
		if max < x {
			max = x
		}
	}

	return max
}
